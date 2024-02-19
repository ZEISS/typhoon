package poller

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/zeiss/typhoon/api/sources"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"
	kadapter "knative.dev/eventing/pkg/adapter/v2"
	logging "knative.dev/pkg/logging"
)

var _ kadapter.Adapter = (*pollerAdapter)(nil)

type pollerAdapter struct {
	logger    *zap.SugaredLogger
	metricTag *kadapter.MetricTag

	client *http.Client
	ce     cloudevents.Client
	cfg    *config
}

// NewAdapter ...
func NewAdapter(ctx context.Context, env kadapter.EnvConfigAccessor, client cloudevents.Client) kadapter.Adapter {
	logger := logging.FromContext(ctx)

	mt := &kadapter.MetricTag{
		ResourceGroup: sources.HTTPResource.String(),
		Namespace:     env.GetNamespace(),
		Name:          env.GetName(),
	}

	cfg := env.(*config)

	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Transport: t}

	return &pollerAdapter{
		logger:    logger,
		metricTag: mt,
		client:    c,
		ce:        client,
		cfg:       cfg,
	}
}

// Start ...
func (h *pollerAdapter) Start(ctx context.Context) error {
	h.logger.Infow("Starting poller", "event-source", h.cfg.EventSource, "event-type", h.cfg.EventType, "url", h.cfg.URL, "interval", h.cfg.Interval)

	ctx = kadapter.ContextWithMetricTag(ctx, h.metricTag)

	req, err := http.NewRequest(h.cfg.Method, h.cfg.URL, nil)
	if err != nil {
		return err
	}

	if h.cfg.BasicAuthUsername != "" && h.cfg.BasicAuthPassword != "" {
		req.SetBasicAuth(h.cfg.BasicAuthUsername, h.cfg.BasicAuthPassword)
	}

	for k, v := range h.cfg.Headers {
		req.Header.Set(k, v)
	}

	t := time.NewTicker(h.cfg.Interval)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			err := h.dispatch(ctx, req)
			if err != nil {
				return err
			}
		}
	}
}

func (h *pollerAdapter) dispatch(ctx context.Context, req *http.Request) error {
	res, err := h.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("non-200 status code")
	}

	event := cloudevents.NewEvent(cloudevents.VersionV1)
	event.SetType(h.cfg.EventType)
	event.SetSource(h.cfg.EventSource)

	err = event.SetData(cloudevents.ApplicationJSON, body)
	if err != nil {
		return err
	}

	r := h.ce.Send(ctx, event)
	if r != nil && !cloudevents.IsACK(r) {
		return errors.New("failed to send event")
	}

	return nil
}
