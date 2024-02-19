package webhook

import (
	"context"

	"github.com/zeiss/typhoon/api/sources"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	kadapter "knative.dev/eventing/pkg/adapter/v2"
	logging "knative.dev/pkg/logging"
)

var _ kadapter.Adapter = (*webhookAdapter)(nil)

type webhookAdapter struct {
	logger    *zap.SugaredLogger
	client    cloudevents.Client
	metricTag *kadapter.MetricTag

	eventType   string
	eventSource string
}

// NewAdapter ...
func NewAdapter(ctx context.Context, env kadapter.EnvConfigAccessor, client cloudevents.Client) kadapter.Adapter {
	mt := &kadapter.MetricTag{
		ResourceGroup: sources.WebhookResource.String(),
		Namespace:     env.GetNamespace(),
		Name:          env.GetName(),
	}

	logger := logging.FromContext(ctx)

	return &webhookAdapter{
		client: client,

		logger:    logger,
		metricTag: mt,
	}
}

// Start ...
func (h *webhookAdapter) Start(ctx context.Context) error {
	ctx = kadapter.ContextWithMetricTag(ctx, h.metricTag)

	app := fiber.New()
	app.Post("/", h.HandleAll)
	app.Get("/healthz", h.HandleHealthz)

	err := app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}

// HandleHealthz ...
func (h *webhookAdapter) HandleHealthz(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// HandleAll ...
func (h *webhookAdapter) HandleAll(c *fiber.Ctx) error {
	logging.FromContext(c.Context()).Info("Received request")

	event := cloudevents.NewEvent(cloudevents.VersionV1)
	event.SetType(h.eventType)
	event.SetSource(h.eventSource)

	err := event.SetData(c.Get("Content-Type"), c.Body())
	if err != nil {
		return err
	}

	e, err := h.client.Request(c.Context(), event)
	if err != nil && !cloudevents.IsACK(err) {
		return err
	}

	if e == nil || e.Data() == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return nil
}
