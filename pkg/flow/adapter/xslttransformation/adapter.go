//go:build !noclibs

package xslttransformation

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	xslt "github.com/wamuir/go-xslt"
	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/flow"
	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"github.com/zeiss/typhoon/pkg/metrics"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

var _ pkgadapter.Adapter = (*xsltTransformAdapter)(nil)

type xsltTransformAdapter struct {
	defaultXSLT  *xslt.Stylesheet
	xsltOverride bool

	replier  *targetce.Replier
	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
	sink     string

	mt *pkgadapter.MetricTag
	sr *metrics.EventProcessingStatsReporter
}

// NewTarget adapter implementation
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: flow.XSLTTransformationResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envAccessor)

	if err := env.validate(); err != nil {
		logger.Panicf("Configuration error: %v", err)
	}

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticDataContentType(cloudevents.ApplicationXML),
		targetce.ReplierWithStaticErrorDataContentType(*cloudevents.StringOfApplicationJSON()),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicyAlways))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	adapter := &xsltTransformAdapter{
		xsltOverride: env.AllowXSLTOverride,

		replier:  replier,
		ceClient: ceClient,
		logger:   logger,
		sink:     env.Sink,

		mt: mt,
		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}

	if env.XSLT != "" {
		adapter.defaultXSLT, err = xslt.NewStylesheet([]byte(env.XSLT))
		if err != nil {
			logger.Panicf("XSLT validation error: %v", err)
		}

		runtime.SetFinalizer(adapter.defaultXSLT, (*xslt.Stylesheet).Close)
	}

	return adapter
}

// Start is a blocking function and will return if an error occurs
// or the context is cancelled.
func (a *xsltTransformAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting XSLT transformer")
	ctx = pkgadapter.ContextWithMetricTag(ctx, a.mt)
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

// nolint:gocyclo
func (a *xsltTransformAdapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	isStructuredTransform := event.Type() == v1alpha1.EventTypeXSLTTransformation
	if isStructuredTransform && !a.xsltOverride {
		return a.replier.Error(&event, targetce.ErrorCodeRequestValidation,
			errors.New("it is not allowed to override XSLT per CloudEvent"), nil)
	}

	isXML := event.DataMediaType() == cloudevents.ApplicationXML

	var style *xslt.Stylesheet
	var xmlin []byte
	var err error

	switch {
	case isStructuredTransform:
		req := &XSLTTransformationStructuredRequest{}
		if err := event.DataAs(req); err != nil {
			return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
		}

		xmlin = []byte(req.XML)
		style, err = xslt.NewStylesheet([]byte(req.XSLT))
		if err != nil {
			return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
		}
		defer style.Close()

	case isXML:
		xmlin = event.DataEncoded
		style = a.defaultXSLT

	default:
		return a.replier.Error(&event, targetce.ErrorCodeRequestValidation,
			errors.New("unexpected type or media-type for the incoming event"), nil)
	}

	res, err := style.Transform(xmlin)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestValidation,
			fmt.Errorf("error processing XML with XSLT: %w", err), nil)
	}

	if a.sink != "" {
		event.SetType(event.Type() + ".response")
		if err := event.SetData(cloudevents.ApplicationXML, res); err != nil {
			return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
		}

		if result := a.ceClient.Send(ctx, event); !cloudevents.IsACK(result) {
			return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, "sending the cloudevent to the sink")
		}
		return nil, cloudevents.ResultACK
	}

	return a.replier.Ok(&event, res, targetce.ResponseWithDataContentType(cloudevents.ApplicationXML))
}
