package datadogtarget

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/targets"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"github.com/zeiss/typhoon/pkg/metrics"
	targetce "github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

const (
	apiBaseDomain     = "https://api"
	logsAPIBaseDomain = "https://http-intake.logs"
	apiKeyHeader      = "DD-API-KEY"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

// NewTarget adapter implementation
func NewTarget(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &pkgadapter.MetricTag{
		ResourceGroup: targets.DatadogTargetResource.String(),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	metrics.MustRegisterEventProcessingStatsView()

	env := envAcc.(*envAccessor)

	replier, err := targetce.New(env.Component, logger.Named("replier"),
		targetce.ReplierWithStatefulHeaders(env.BridgeIdentifier),
		targetce.ReplierWithStaticResponseType(v1alpha1.EventTypeDatadogResponse),
		targetce.ReplierWithPayloadPolicy(targetce.PayloadPolicy(env.CloudEventPayloadPolicy)))
	if err != nil {
		logger.Panicf("Error creating CloudEvents replier: %v", err)
	}

	return &datadogAdapter{
		apiKey:     env.APIKey,
		apiURL:     fmt.Sprintf("%s.%s", apiBaseDomain, env.Site),
		apiLogsURL: fmt.Sprintf("%s.%s", logsAPIBaseDomain, env.Site),

		replier:    replier,
		httpClient: http.DefaultClient,
		ceClient:   ceClient,
		logger:     logger,

		sr: metrics.MustNewEventProcessingStatsReporter(mt),
	}
}

var _ pkgadapter.Adapter = (*datadogAdapter)(nil)

type datadogAdapter struct {
	apiKey     string
	apiURL     string
	apiLogsURL string

	replier    *targetce.Replier
	httpClient *http.Client
	ceClient   cloudevents.Client
	logger     *zap.SugaredLogger

	sr *metrics.EventProcessingStatsReporter
}

// Returns if stopCh is closed or Send() returns an error.
func (a *datadogAdapter) Start(ctx context.Context) error {
	a.logger.Info("Starting Datadog adapter")
	return a.ceClient.StartReceiver(ctx, a.dispatch)
}

func (a *datadogAdapter) dispatch(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	switch typ := event.Type(); typ {
	case v1alpha1.EventTypeDatadogMetric:
		return a.postMetric(ctx, event)
	case v1alpha1.EventTypeDatadogEvent:
		return a.postEvent(ctx, event)
	case v1alpha1.EventTypeDatadogLog:
		return a.postLog(ctx, event)
	default:
		return a.replier.Error(&event, targetce.ErrorCodeEventContext, fmt.Errorf("event type %q is not supported", typ), nil)
	}
}

func (a *datadogAdapter) postLog(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if err := event.DataAs(&LogData{}); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	request, err := newLogsAPIRequest(ctx, a.apiLogsURL, "/v1/input", a.apiKey, event.Data())
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	res, err := a.httpClient.Do(request)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeParseResponse, err, nil)
	}

	if res.StatusCode != http.StatusOK {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess,
			fmt.Errorf("received HTTP code %d", res.StatusCode),
			map[string]string{"body": string(resBody)})
	}

	return a.replier.Ok(&event, resBody)
}

func (a *datadogAdapter) postEvent(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if err := event.DataAs(&EventData{}); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	request, err := newAPIRequest(ctx, a.apiURL, "/api/v1/events", a.apiKey, event.Data())
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	res, err := a.httpClient.Do(request)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeParseResponse, err, nil)
	}

	if res.StatusCode != http.StatusAccepted {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess,
			fmt.Errorf("received HTTP code %d", res.StatusCode),
			map[string]string{"body": string(resBody)})
	}

	return a.replier.Ok(&event, resBody)
}

func (a *datadogAdapter) postMetric(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if err := event.DataAs(&MetricData{}); err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeRequestParsing, err, nil)
	}

	request, err := newAPIRequest(ctx, a.apiURL, "/api/v1/series", a.apiKey, event.Data())
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	res, err := a.httpClient.Do(request)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, nil)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return a.replier.Error(&event, targetce.ErrorCodeParseResponse, err, nil)
	}

	if res.StatusCode != http.StatusAccepted {
		return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess,
			fmt.Errorf("received HTTP code %d", res.StatusCode),
			map[string]string{"body": string(resBody)})
	}

	return a.replier.Ok(&event, resBody)
}

// newAPIRequest returns a POST http.Request that is ready to send to the Datadog general-purpose API.
func newAPIRequest(ctx context.Context, host, path, apiKey string, body []byte) (*http.Request, error) {
	return newAPIRequestWithHost(ctx, host, path, apiKey, body)
}

// newLogsAPIRequest returns a POST http.Request that is ready to send to the Datadog logs API.
func newLogsAPIRequest(ctx context.Context, host, path, apiKey string, body []byte) (*http.Request, error) {
	return newAPIRequestWithHost(ctx, host, path, apiKey, body)
}

// newAPIRequestWithHost returns a POST http.Request that is ready to send to the Datadog API.
func newAPIRequestWithHost(ctx context.Context, host, path, apiKey string, body []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, host+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set(contentTypeHeader, contentTypeJSON)
	req.Header.Set(apiKeyHeader, apiKey)

	return req, nil
}
