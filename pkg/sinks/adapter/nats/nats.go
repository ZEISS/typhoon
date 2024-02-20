package natssink

import (
	"context"

	"github.com/zeiss/typhoon/pkg/apis/sinks"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	kadapter "knative.dev/eventing/pkg/adapter/v2"
	logging "knative.dev/pkg/logging"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

var _ kadapter.Adapter = (*natsAdapter)(nil)

type config struct {
	url     string
	subject string
}

type natsAdapter struct {
	cfg *config

	client cloudevents.Client
	js     jetstream.JetStream
	conn   *nats.Conn

	logger *zap.SugaredLogger
	mt     *kadapter.MetricTag
}

// NewSink returns a new NATS.io adapter
func NewSink(ctx context.Context, envAcc kadapter.EnvConfigAccessor, client cloudevents.Client) kadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &kadapter.MetricTag{
		ResourceGroup: sinks.NatsSinkResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	return &natsAdapter{
		client: client,
		logger: logger,
		mt:     mt,
	}
}

// Start is called to start the NATS.io adapter.
func (a *natsAdapter) Start(ctx context.Context) error {
	a.logger.Info("starting NATS.io adapter")

	nc, err := nats.Connect(a.cfg.url)
	if err != nil {
		return err
	}
	a.conn = nc

	js, err := jetstream.New(nc)
	if err != nil {
		return err
	}
	a.js = js

	err = a.client.StartReceiver(ctx, a.process)
	if err != nil {
		return err
	}

	return nil
}

func (a *natsAdapter) process(event cloudevents.Event) cloudevents.Result {
	msg := event.Data()

	f, err := a.js.PublishAsync(a.cfg.subject, msg)
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
