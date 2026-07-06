package workertransformation

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/flow"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

// NewAdapter adapter implementation
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: flow.WorkerTransformationResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType("com.zeiss.workertransformation.error"),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &workeradapter{
		sink:     env.Sink,
		replier:  replier,
		ceClient: ceClient,
		logger:   logger,

		mt: mt,
	}
}

var _ pkgadapter.Adapter = (*workeradapter)(nil)

type workeradapter struct {
	ceClient cloudevents.Client
	replier  *targetce.Replier
	logger   *zap.SugaredLogger
	mt       *pkgadapter.MetricTag
	sink     string
}

// Start is a blocking function and will return if an error occurs
// or the context is cancelled.
func (a *workeradapter) Start(ctx context.Context) error {
	a.logger.Info("Starting WorkerTransformation Adapter")
	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)

	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *workeradapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	var data interface{}
	if err := event.DataAs(&data); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	return &event, cloudevents.ResultACK
}
