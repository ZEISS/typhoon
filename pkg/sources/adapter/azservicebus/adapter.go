package azservicebus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/devigned/tab"
	"go.uber.org/zap"

	// nolint:staticcheck
	"nhooyr.io/websocket"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/zeiss/typhoon/pkg/apis/sources"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"github.com/zeiss/typhoon/pkg/sources/adapter/azservicebus/trace"
)

const (
	resourceProviderServiceBus = "Microsoft.ServiceBus"

	resourceTypeQueues        = "queues"
	resourceTypeTopics        = "topics"
	resourceTypeSubscriptions = "subscriptions"
)

const (
	envKeyName  = "SERVICEBUS_KEY_NAME"
	envKeyValue = "SERVICEBUS_KEY_VALUE"
	envConnStr  = "SERVICEBUS_CONNECTION_STRING"
)

// envConfig is a set parameters sourced from the environment for the source's
// adapter.
type envConfig struct {
	pkgadapter.EnvConfig

	// Resource ID of the Service Bus entity (Queue or Topic subscription).
	EntityResourceID string `envconfig:"SERVICEBUS_ENTITY_RESOURCE_ID" required:"true"`

	// Name of a message processor which takes care of converting Service
	// Bus messages to CloudEvents.
	//
	// Supported values: [ default ]
	MessageProcessor string `envconfig:"SERVICEBUS_MESSAGE_PROCESSOR" default:"default"`

	// WebSocketsEnable.
	WebSocketsEnable bool `envconfig:"SERVICEBUS_WEBSOCKETS_ENABLE" default:"false"`

	// MaxConcurrent is the maximum number of goroutines that
	// will be used to process messages.
	MaxConcurrent int `envconfig:"SERVICEBUS_MAX_CONCURRENT" default:"10"`

	// The environment variables below aren't read from the envConfig struct
	// by the Service Bus SDK, but rather directly using os.Getenv().
	// They are nevertheless listed here for documentation purposes.
	_ string `envconfig:"AZURE_TENANT_ID"`
	_ string `envconfig:"AZURE_CLIENT_ID"`
	_ string `envconfig:"AZURE_CLIENT_SECRET"`
	_ string `envconfig:"SERVICEBUS_KEY_NAME"`
	_ string `envconfig:"SERVICEBUS_KEY_VALUE"`
	_ string `envconfig:"SERVICEBUS_CONNECTION_STRING"`
}

// adapter implements the source's adapter.
type adapter struct {
	logger *zap.SugaredLogger
	mt     *pkgadapter.MetricTag

	msgRcvr  *azservicebus.Receiver
	ceClient cloudevents.Client

	msgPrcsr      MessageProcessor
	maxConcurrent int
}

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() pkgadapter.EnvConfigAccessor {
	return &envConfig{}
}

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		Namespace: envAcc.GetNamespace(),
		Name:      envAcc.GetName(),
	}

	env := envAcc.(*envConfig)

	entityID, err := parseServiceBusResourceID(env.EntityResourceID)
	if err != nil {
		logger.Panic("Unable to parse entity ID "+strconv.Quote(env.EntityResourceID), zap.Error(err))
	}

	client, err := clientFromEnvironment(entityID, newAzureServiceBusClientOptions(
		webSocketsClientOption(env.WebSocketsEnable)))
	if err != nil {
		logger.Panic("Unable to obtain interface for Service Bus Namespace", zap.Error(err))
	}

	var rcvr *azservicebus.Receiver
	switch entityID.ResourceType {
	case resourceTypeQueues:
		rcvr, err = client.NewReceiverForQueue(entityID.ResourceName, nil)
		mt.ResourceGroup = sources.AzureServiceBusQueueSourceResource.String()
	case resourceTypeSubscriptions, resourceTypeTopics:
		rcvr, err = client.NewReceiverForSubscription(entityID.ResourceName, entityID.SubResourceName, nil)
		mt.ResourceGroup = sources.AzureServiceBusTopicSourceResource.String()
	}
	if err != nil {
		logger.Panic("Unable to obtain message receiver for Service Bus entity "+strconv.Quote(strconv.Quote(entityPath(entityID))), zap.Error(err))
	}

	ceSource := env.EntityResourceID

	var msgPrcsr MessageProcessor
	switch env.MessageProcessor {
	case "default":
		msgPrcsr = &defaultMessageProcessor{ceSource: ceSource}
	default:
		logger.Panic("unsupported message processor " + strconv.Quote(env.MessageProcessor))
	}

	// The Service Bus client uses the default "NoOpTracer" tab.Tracer
	// implementation, which does not produce any log message. We register
	// a custom implementation so that event handling errors are logged via
	// Knative's logging facilities.
	tab.Register(trace.NewNoOpTracerWithLogger(logger))

	return &adapter{
		logger: logger,
		mt:     mt,

		ceClient: ceClient,

		msgRcvr:       rcvr,
		msgPrcsr:      msgPrcsr,
		maxConcurrent: env.MaxConcurrent,
	}
}

// parseServiceBusResourceID parses the given resource ID string to a
// structured resource ID, and validates that this resource ID refers to a
// Service Bus entity.
func parseServiceBusResourceID(resIDStr string) (*v1alpha1.AzureResourceID, error) {
	resID := &v1alpha1.AzureResourceID{}

	err := json.Unmarshal([]byte(strconv.Quote(resIDStr)), resID)
	if err != nil {
		return nil, fmt.Errorf("deserializing resource ID string: %w", err)
	}

	// Must match one of the following patterns:
	//  - /.../providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}
	//  - /.../providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subsName}
	if resID.ResourceProvider != resourceProviderServiceBus ||
		resID.Namespace == "" ||
		resID.ResourceType != resourceTypeQueues && resID.ResourceType != resourceTypeTopics ||
		resID.ResourceType == resourceTypeQueues && resID.SubResourceType != "" ||
		resID.ResourceType == resourceTypeTopics && resID.SubResourceType != resourceTypeSubscriptions {

		return nil, errors.New("resource ID does not refer to a Service Bus entity")
	}

	return resID, nil
}

// entityPath returns the entity path of the given Service Bus entity.
func entityPath(entityID *v1alpha1.AzureResourceID) string {
	switch entityID.ResourceType {
	case resourceTypeQueues:
		queueName := entityID.ResourceName
		return queueName
	case resourceTypeTopics:
		topicName := entityID.ResourceName
		subsName := entityID.SubResourceName
		return topicName + "/Subscriptions/" + subsName
	default:
		return ""
	}
}

// clientFromEnvironment mimics the behaviour of eventhub.NewHubFromEnvironment.
// It returns a azservicebus.Client that is suitable for the
// authentication method selected via environment variables.
func clientFromEnvironment(entityID *v1alpha1.AzureResourceID, clientOptions *azservicebus.ClientOptions) (*azservicebus.Client, error) {
	// SAS authentication (token, connection string)
	connStr := connectionStringFromEnvironment(entityID.Namespace, entityPath(entityID))
	if connStr != "" {
		client, err := azservicebus.NewClientFromConnectionString(connStr, clientOptions)
		if err != nil {
			return nil, fmt.Errorf("creating client from connection string: %w", err)
		}
		return client, nil
	}

	// AAD authentication (service principal)
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create Azure credentials: %w", err)
	}

	fqNamespace := entityID.Namespace + ".servicebus.windows.net"
	client, err := azservicebus.NewClient(fqNamespace, cred, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("creating client from service principal: %w", err)
	}
	return client, nil
}

// connectionStringFromEnvironment returns a Service Bus connection string
// based on values read from the environment.
func connectionStringFromEnvironment(namespace, entityPath string) string {
	connStr := os.Getenv(envConnStr)

	// if a key is set explicitly, it takes precedence and is used to
	// compose a new connection string
	if keyName, keyValue := os.Getenv(envKeyName), os.Getenv(envKeyValue); keyName != "" && keyValue != "" {
		azureEnv := &azure.PublicCloud
		connStr = fmt.Sprintf("Endpoint=sb://%s.%s;SharedAccessKeyName=%s;SharedAccessKey=%s;EntityPath=%s",
			namespace, azureEnv.ServiceBusEndpointSuffix, keyName, keyValue, entityPath)
	}

	return connStr
}

// Start implements adapter.Adapter.
//
// Required permissions:
//
//	Service Bus Queues:
//	  - Microsoft.ServiceBus/namespaces/queues/read
//	Service Bus Topics:
//	  - Microsoft.ServiceBus/namespaces/topics/read
//	  - Microsoft.ServiceBus/namespaces/topics/subscriptions/read
//	Both (DataAction):
//	- Microsoft.ServiceBus/namespaces/messages/receive/action
func (a *adapter) Start(ctx context.Context) error {
	logging.FromContext(ctx).Info("Listening for messages")
	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)

	// We might need to cancel the context to make routines
	// exit if an error occurs at any of them.
	cctx, cancel := context.WithCancel(ctx)

	// Waitgroup makes sure all routines have finished before
	// returning from start.
	wg := &sync.WaitGroup{}

	// We are communicating with routines via channels.
	// Create errChan with capacity to deal with the worst case,
	// which would be one error returned from every routine.
	errChan := make(chan error, a.maxConcurrent)
	msgChan := make(chan *fullMessage)

	// Launch maxConcurrent consumers
	for i := 0; i < a.maxConcurrent; i++ {
		wg.Add(1)
		go func() {
			a.consume(cctx, msgChan, errChan)
			wg.Done()
		}()
	}

	// Launch one producer.
	wg.Add(1)
	go func() {
		a.produce(cctx, msgChan, errChan)
		wg.Done()
	}()

	// This variable store all errors returned from routines.
	errs := []string{}

	// Wait for either context done or an error from any routine.
	select {
	case <-cctx.Done():
	case err := <-errChan:
		// If an error occurs, write it at the errors store, we
		// will
		// errs = append(errs, err)
		errs = append(errs, err.Error())

	}

	// cancel the context to bring all routines to an end.
	cancel()

	// Wait for all routines to exit. If routines fail while exiting
	// they will write to the errChan, which has capacity to store
	// an error per routine without blocking.
	wg.Wait()

	// Gather and sumarize errors from routines
	for err := range errChan {
		// errs = append(errs, err)
		errs = append(errs, err.Error())
	}

	// If there are errors, return them as a single error.
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ". "))
	}

	return nil
}

// convenience structure for message processing.
type fullMessage struct {
	received     *azservicebus.ReceivedMessage
	serializable *Message
}

func (a *adapter) produce(ctx context.Context, msgChan chan *fullMessage, errChan chan error) {
	const maxMessages = 100

	for {
		messages, err := a.msgRcvr.ReceiveMessages(ctx, maxMessages, nil)

		switch {
		case err == nil:
			for _, m := range messages {
				msg, err := toMessage(m)
				if err != nil {
					errChan <- fmt.Errorf("error transforming message: %w", err)
					return
				}

				msgChan <- &fullMessage{
					received:     m,
					serializable: msg,
				}
			}
		case errors.Is(err, context.Canceled):
			return
		default:
			errChan <- fmt.Errorf("error receiving messages: %w", err)
			return
		}
	}
}

func (a *adapter) consume(ctx context.Context, msgChan chan *fullMessage, errChan chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		case fm := <-msgChan:
			if err := a.handleMessage(ctx, fm.serializable); err != nil {
				errChan <- fmt.Errorf("error handling message: %w", err)
				return
			}
			if err := a.msgRcvr.CompleteMessage(ctx, fm.received, nil); err != nil {
				errChan <- fmt.Errorf("error completing message: %w", err)
				return
			}
		}
	}
}

// handleMessage handles a single Service Bus message.
func (a *adapter) handleMessage(ctx context.Context, msg *Message) error {
	if msg == nil {
		return nil
	}

	events, err := a.msgPrcsr.Process(msg)
	if err != nil {
		return fmt.Errorf("processing Service Bus message with ID %s: %w", msg.ReceivedMessage.MessageID, err)
	}

	var sendErrs errList
	var ve event.ValidationError

	for _, ev := range events {
		if err := ev.Validate(); err != nil && errors.As(err, &ve) {
			ev = sanitizeEvent(ve, ev)
		}

		if err := sendCloudEvent(ctx, a.ceClient, ev); err != nil {
			sendErrs.errs = append(sendErrs.errs,
				fmt.Errorf("failed to send event with ID %s: %w", ev.ID(), err),
			)
			continue
		}
	}

	if len(sendErrs.errs) != 0 {
		return fmt.Errorf("sending events to the sink: %w", sendErrs)
	}

	return nil
}

// sendCloudEvent sends a single CloudEvent to the event sink.
func sendCloudEvent(ctx context.Context, cli cloudevents.Client, event *cloudevents.Event) protocol.Result {
	if result := cli.Send(ctx, *event); !cloudevents.IsACK(result) {
		return result
	}
	return nil
}

// errList is an aggregate of errors.
type errList struct {
	errs []error
}

var _ error = (*errList)(nil)

// Error implements the error interface.
func (e errList) Error() string {
	if len(e.errs) == 0 {
		return ""
	}
	return fmt.Sprintf("%q", e.errs)
}

// sanitizeEvent tries to fix the validation issues listed in the given
// cloudevents.ValidationError, and returns a sanitized version of the event.
//
// For now, this helper exists solely to fix CloudEvents sent by Azure Event
// Grid, which often contain
//
//	"dataschema": "#"
func sanitizeEvent(validErrs event.ValidationError, origEvent *cloudevents.Event) *cloudevents.Event {
	for attr := range validErrs {
		if attr == "dataschema" {
			origEvent.SetDataSchema("")
		}
	}

	return origEvent
}

type clientOption func(*azservicebus.ClientOptions)

func newAzureServiceBusClientOptions(opts ...clientOption) *azservicebus.ClientOptions {
	co := &azservicebus.ClientOptions{}
	for _, opt := range opts {
		opt(co)
	}
	return co
}

func webSocketsClientOption(webSocketsEnable bool) clientOption {
	return func(opts *azservicebus.ClientOptions) {
		if webSocketsEnable {
			opts.NewWebSocketConn = func(ctx context.Context, args azservicebus.NewWebSocketConnArgs) (net.Conn, error) {
				// nolint:staticcheck
				opts := &websocket.DialOptions{Subprotocols: []string{"amqp"}}
				// nolint:staticcheck
				wssConn, _, err := websocket.Dial(ctx, args.Host, opts)
				if err != nil {
					return nil, fmt.Errorf("creating client: %w", err)
				}

				// nolint:contextcheck,staticcheck
				return websocket.NetConn(context.Background(), wssConn, websocket.MessageBinary), nil
			}
		}
	}
}
