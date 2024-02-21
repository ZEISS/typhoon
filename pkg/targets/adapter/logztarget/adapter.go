package logztarget

import (
	"context"

	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/logzio/logzio-go"

	"github.com/zeiss/typhoon/pkg/apis/targets"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"github.com/zeiss/typhoon/pkg/metrics"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

// NewTarget adapter implementation
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: targets.LogzTargetResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envAccessor)

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType(v1alpha1.EventTypeLogzShipResponse),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	l, err := logzio.New(
		env.ShippingToken,
		logzio.SetUrl("https://"+env.LogsListenerURL+":8071"),
	)
	if err != nil {
		panic(err)
	}

	return &logzAdapter{
		l: l,

		replier:  replier,
		ceClient: ceClient,
		logger:   logger,

		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}
}

var _ pkgadapter.Adapter = (*logzAdapter)(nil)

type logzAdapter struct {
	l *logzio.LogzioSender

	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger

	sr *metrics.EventProcessingStatsReporter
}

// Returns if stopCh is closed or Send() returns an error.
func (a *logzAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting Logz adapter")
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *logzAdapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	err := a.l.Send(event.Data())
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	return a.replier.Ok(&event, "ok")
}
