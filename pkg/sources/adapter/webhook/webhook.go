package webhook

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/zeiss/typhoon/api/sources"
	"go.uber.org/zap"
	kadapter "knative.dev/eventing/pkg/adapter/v2"
	logging "knative.dev/pkg/logging"
)

var _ kadapter.Adapter = (*webhookAdapter)(nil)

type webhookAdapter struct {
	logger    *zap.SugaredLogger
	client    cloudevents.Client
	metricTag *kadapter.MetricTag
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
func (a *webhookAdapter) Start(ctx context.Context) error {
	ctx = kadapter.ContextWithMetricTag(ctx, a.metricTag)

	// m := http.NewServeMux()
	// m.HandleFunc("/", h.handleAll(ctx))

	// s := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", serverPort),
	// 	Handler: m,
	// }

	return nil
}
