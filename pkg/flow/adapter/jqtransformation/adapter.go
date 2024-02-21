package jqtransformation

import (
	"context"
	"encoding/json"

	"github.com/itchyny/gojq"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/flow"
	"github.com/zeiss/typhoon/pkg/metrics"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

// NewAdapter adapter implementation
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: flow.JQTransformationResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	mgithub.com / zeiss / typhoonStatsView()
	github.com / zeiss / typhoon
	env := envAcc.(*envAccessor)

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType("com.zeiss.jqtransformation.error"),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	query, err := gojq.Parse(env.Query)
	if err != nil {
		logger.Panicf("Error creating query: %v", err)
	}

	return &jqadapter{
		query: query,

		sink:     env.Sink,
		replier:  replier,
		ceClient: ceClient,
		logger:   logger,

		mt: mt,
		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}
}

var _ pkgadapter.Adapter = (*jqadapter)(nil)

type jqadapter struct {
	query *gojq.Query

	sink     string
	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger

	mt *pkgadapter.MetricTag
	sr *metrics.EventProcessingStatsReporter
}

// Start is a blocking function and will return if an error occurs
// or the context is cancelled.
func (a *jqadapter) Start(ctx context.Context) error {
	a.logger.Info("Starting JQTransformation Adapter")
	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *jqadapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	var data interface{}
	var qd interface{}
	if err := event.DataAs(&data); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	iter := a.query.Run(data)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
		}
		qd = v
	}

	// Reserialize the query results for the response
	bs, err := json.Marshal(&qd)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	if err := event.SetData(cloudevents.ApplicationJSON, bs); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	if a.sink != "" {
		if result := a.ceClient.Send(ctx, event); !cloudevents.IsACK(result) {
			return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, result, "sending the cloudevent to the sink")
		}
		return nil, cloudevents.ResultACK
	}

	return &event, cloudevents.ResultACK
}
