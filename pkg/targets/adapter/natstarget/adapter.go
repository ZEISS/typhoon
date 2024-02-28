package natstarget

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/zeiss/typhoon/pkg/apis/targets"
	"go.uber.org/zap"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	url     string `envconfig:"NATS_URL"`
	subject string `envconfig:"NATS_SUBJECT"`
}

var _ pkgadapter.Adapter = (*natsAdapter)(nil)

type natsAdapter struct {
	client cloudevents.Client
	js     jetstream.JetStream
	conn   *nats.Conn

	subject string

	logger *zap.SugaredLogger
	mt     *pkgadapter.MetricTag
}

// NewTarget adapter implementation
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, client cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: targets.KafkaTargetResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	nc, err := nats.Connect(env.url)
	if err != nil {
		logger.Panicw("failed to connect to NATS", zap.Error(err))
	}

	js, err := jetstream.New(nc)
	if err != nil {
		logger.Panicw("failed to connect to JetStream", zap.Error(err))
	}

	return &natsAdapter{
		conn:    nc,
		mt:      mt,
		js:      js,
		client:  client,
		subject: env.subject,
	}
}

// Start is the main entrypoint for the adapter
func (a *natsAdapter) Start(ctx context.Context) error {
	a.logger.Info("starting NATS.io adapter")

	defer func() {
		a.conn.Close()
	}()

	return a.client.StartReceiver(ctx, a.dispatch)
}

func (a *natsAdapter) dispatch(event cloudevents.Event) cloudevents.Result {
	msg := event.Data()

	f, err := a.js.PublishAsync(a.subject, msg)
	if err != nil {
		return err
	}

	select {
	case <-f.Err():
		return err
	case <-f.Ok():
		return cloudevents.ResultACK

	}
}
