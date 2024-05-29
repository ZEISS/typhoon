package salesforcesource

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/oauth2/clientcredentials"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/sources"
	sfclient "github.com/zeiss/typhoon/pkg/sources/adapter/salesforcesource/client"
)

const eventType = "com.salesforce.stream.message"

type salesforceAdapter struct {
	sfVersion         string
	sfChannel         string
	sfInstanceURL     string
	sfInitialReplayID int

	client *http.Client

	dispatcher *eventDispatcher
	logger     *zap.SugaredLogger
	mt         *pkgadapter.MetricTag
}

type eventDispatcher struct {
	eventSource string

	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
}

var (
	_ pkgadapter.Adapter       = (*salesforceAdapter)(nil)
	_ sfclient.EventDispatcher = (*eventDispatcher)(nil)
)

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: sources.SalesforceSourceResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	source := env.Name
	if env.SubscriptionChannel[0] != '/' {
		source += "/"
	}
	source += env.SubscriptionChannel

	dispatcher := &eventDispatcher{
		eventSource: source,
		ceClient:    ceClient,
		logger:      logger.Named("dispatcher"),
	}

	cfg := clientcredentials.Config{
		TokenURL:     env.TokenURL,
		ClientID:     env.ClientID,
		ClientSecret: env.ClientSecret,
		// TODO(@katallaxie) we need to figure out how to get the full scope
		// Scopes:       []string{"full"},
	}

	client := cfg.Client(ctx)

	adapter := &salesforceAdapter{
		sfVersion:         env.Version,
		sfChannel:         env.SubscriptionChannel,
		sfInitialReplayID: env.SubscriptionReplayID,
		sfInstanceURL:     env.InstanceURL,

		client: client,

		dispatcher: dispatcher,
		logger:     logger,
		mt:         mt,
	}

	return adapter
}

// Start runs the handler.
func (a *salesforceAdapter) Start(ctx context.Context) (err error) {
	replayID := a.sfInitialReplayID

	subs := []sfclient.Subscription{
		{
			Channel:  a.sfChannel,
			ReplayID: replayID,
		},
	}

	client := sfclient.NewBayeux(a.sfVersion, a.sfInstanceURL, subs, a.dispatcher, a.client, a.logger.Named("bayeux"))

	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)

	return client.Start(ctx)
}

func (e *eventDispatcher) DispatchEvent(ctx context.Context, msg *sfclient.ConnectResponse) {
	event := cloudevents.NewEvent(cloudevents.VersionV1)

	event.SetType(eventType)
	event.SetSource(e.eventSource)
	event.SetID(uuid.New().String())
	event.SetSubject(subjectNameFromConnectResponse(msg))
	if err := event.SetData(cloudevents.ApplicationJSON, msg.Data); err != nil {
		e.logger.Errorw("Failed to set event data", zap.Error(err))
		return
	}

	if result := e.ceClient.Send(ctx, event); !cloudevents.IsACK(result) {
		e.logger.Errorw("Could not send CloudEvent", zap.Error(result))
		return
	}
}

func (e *eventDispatcher) DispatchError(err error) {
	e.logger.Errorw("Error receiving events", zap.Error(err))
}

func subjectNameFromConnectResponse(msg *sfclient.ConnectResponse) string {
	// if ChangeDataCapture look for entity/operation
	cdc := &sfclient.ChangeDataCapturePayload{}
	if err := json.Unmarshal(msg.Data.Payload, cdc); err == nil {
		return cdc.ChangeEventHeader.EntityName + "/" + cdc.ChangeEventHeader.ChangeType
	}

	// if PushTopic look for object-name/event-operation
	ptso := &sfclient.PushTopicSObject{}
	if err := json.Unmarshal(msg.Data.Payload, ptso); err == nil {
		return ptso.Name + "/" + msg.Data.Event.Type
	}

	// default to channel
	return msg.Channel
}
