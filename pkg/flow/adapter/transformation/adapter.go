package transformation

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/flow"
	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation/common/storage"
	"github.com/zeiss/typhoon/pkg/metrics"
)

type envConfig struct {
	pkgadapter.EnvConfig

	// Sink URL where to send cloudevents
	Sink string `envconfig:"K_SINK"`

	// Transformation specifications
	TransformationContext string `envconfig:"TRANSFORMATION_CONTEXT"`
	TransformationData    string `envconfig:"TRANSFORMATION_DATA"`
}

// adapter contains Pipelines for CE transformations and CloudEvents client.
type adapter struct {
	ContextPipeline *Pipeline
	DataPipeline    *Pipeline

	mt *pkgadapter.MetricTag
	sr *metrics.EventProcessingStatsReporter

	sink string

	client cloudevents.Client
	logger *zap.SugaredLogger
}

// ceContext represents CloudEvents context structure but with exported Extensions.
type ceContext struct {
	*cloudevents.EventContextV1 `json:",inline"`
	Extensions                  map[string]interface{} `json:"Extensions,omitempty"`
}

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() pkgadapter.EnvConfigAccessor {
	return &envConfig{}
}

func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: flow.TransformationResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envConfig)

	trnContext, trnData := []v1alpha1.Transform{}, []v1alpha1.Transform{}
	err := json.Unmarshal([]byte(env.TransformationContext), &trnContext)
	if err != nil {
		logger.Fatalf("Cannot unmarshal context transformation env variable: %v", err)
	}
	err = json.Unmarshal([]byte(env.TransformationData), &trnData)
	if err != nil {
		logger.Fatalf("Cannot unmarshal data transformation env variable: %v", err)
	}

	sharedStorage := storage.New()

	contextPl, err := newPipeline(trnContext, sharedStorage)
	if err != nil {
		logger.Fatalf("Cannot create context transformation pipeline: %v", err)
	}

	dataPl, err := newPipeline(trnData, sharedStorage)
	if err != nil {
		logger.Fatalf("Cannot create data transformation pipeline: %v", err)
	}

	return &adapter{
		ContextPipeline: contextPl,
		DataPipeline:    dataPl,

		mt: mt,
		sr: metrics.MustNewEventProcessingStatsReporter(mt),

		sink:   env.Sink,
		client: ceClient,
		logger: logger,
	}
}

// Start runs CloudEvent receiver and applies transformation Pipeline
// on incoming events.
func (t *adapter) Start(ctx context.Context) error {
	t.logger.Info("Starting Transformation adapter")

	var receiver interface{}
	receiver = t.receiveAndReply
	if t.sink != "" {
		ctx = cloudevents.ContextWithTarget(ctx, t.sink)
		receiver = t.receiveAndSend
	}

	ctx = pkgadapter.ContextWithMetricTag(ctx, t.mt)

	return t.client.StartReceiver(ctx, receiver)
}

func (t *adapter) receiveAndReply(event cloudevents.Event) (*cloudevents.Event, error) {
	ceTypeTag := metrics.TagEventType(event.Type())
	ceSrcTag := metrics.TagEventSource(event.Source())

	start := time.Now()
	defer func() {
		t.sr.ReportProcessingLatency(time.Since(start), ceTypeTag, ceSrcTag)
	}()

	result, err := t.applyTransformations(event)
	if err != nil {
		t.sr.ReportProcessingError(false, ceTypeTag, ceSrcTag)
	} else {
		t.sr.ReportProcessingSuccess(ceTypeTag, ceSrcTag)
	}

	return result, err
}

func (t *adapter) receiveAndSend(ctx context.Context, event cloudevents.Event) error {
	ceTypeTag := metrics.TagEventType(event.Type())
	ceSrcTag := metrics.TagEventSource(event.Source())

	start := time.Now()
	// nolint:contextcheck
	defer func() {
		t.sr.ReportProcessingLatency(time.Since(start), ceTypeTag, ceSrcTag)
	}()

	result, err := t.applyTransformations(event)
	if err != nil {
		// nolint:contextcheck
		t.sr.ReportProcessingError(false, ceTypeTag, ceSrcTag)
		return err
	}

	if result := t.client.Send(ctx, *result); !cloudevents.IsACK(result) {
		// nolint:contextcheck
		t.sr.ReportProcessingError(false, ceTypeTag, ceSrcTag)
		return result
	}

	// nolint:contextcheck
	t.sr.ReportProcessingSuccess(ceTypeTag, ceSrcTag)
	return nil
}

// nolint:gocyclo
func (t *adapter) applyTransformations(event cloudevents.Event) (*cloudevents.Event, error) {
	// HTTPTargets sets content type from HTTP headers, i.e.:
	// "datacontenttype: application/json; charset=utf-8"
	// so we must use "contains" instead of strict equality
	if !strings.Contains(event.DataContentType(), cloudevents.ApplicationJSON) {
		err := fmt.Errorf("CE Content-Type %q is not supported", event.DataContentType())
		t.logger.Error("Bad Content-Type", zap.Error(err))
		return nil, err
	}

	localContext := ceContext{
		EventContextV1: event.Context.AsV1(),
		Extensions:     event.Context.AsV1().GetExtensions(),
	}

	localContextBytes, err := json.Marshal(localContext)
	if err != nil {
		t.logger.Error("Cannot encode CE context", zap.Error(err))
		return nil, fmt.Errorf("cannot encode CE context: %w", err)
	}

	// init indicates if we need to run initial step transformation
	init := true
	var errs []error

	eventUniqueID := fmt.Sprintf("%s-%s", event.ID(), event.Source())

	// remove event-related variables after the transformation is done.
	// since the storage is shared, flush can be done for one pipeline.
	defer t.ContextPipeline.Storage.Flush(eventUniqueID)

	// Run init step such as load Pipeline variables first
	eventContext, err := t.ContextPipeline.apply(eventUniqueID, localContextBytes, init)
	if err != nil {
		errs = append(errs, err)
	}
	eventPayload, err := t.DataPipeline.apply(eventUniqueID, event.Data(), init)
	if err != nil {
		errs = append(errs, err)
	}

	// CE Context transformation
	if eventContext, err = t.ContextPipeline.apply(eventUniqueID, eventContext, !init); err != nil {
		errs = append(errs, err)
	}

	newContext := ceContext{}
	if err := json.Unmarshal(eventContext, &newContext); err != nil {
		t.logger.Error("Cannot decode CE new context", zap.Error(err))
		return nil, fmt.Errorf("cannot decode CE new context: %w", err)
	}
	event.Context = newContext
	for k, v := range newContext.Extensions {
		if err := event.Context.SetExtension(k, v); err != nil {
			t.logger.Error("Cannot set CE extension", zap.Error(err))
			return nil, fmt.Errorf("cannot set CE extension: %w", err)
		}
	}

	// CE Data transformation
	if eventPayload, err = t.DataPipeline.apply(eventUniqueID, eventPayload, !init); err != nil {
		errs = append(errs, err)
	}
	if err = event.SetData(cloudevents.ApplicationJSON, eventPayload); err != nil {
		t.logger.Error("Cannot set CE data", zap.Error(err))
		return nil, fmt.Errorf("cannot set CE data: %w", err)
	}
	// Failed transformation operations should not stop event flow
	// therefore, just log the errors
	if len(errs) != 0 {
		t.logger.Error("Event transformation errors", zap.Errors("errors", errs))
	}

	return &event, nil
}
