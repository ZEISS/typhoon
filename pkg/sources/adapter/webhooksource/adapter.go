package webhooksource

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/typhoon/pkg/apis/sources"
)

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc adapter.EnvConfigAccessor, ceClient cloudevents.Client) adapter.Adapter {
	mt := &adapter.MetricTag{
		ResourceGroup: conv.String(sources.WebhookSourceResource),
		Namespace:     envAcc.GetNamespace(),
		Name:          envAcc.GetName(),
	}

	env := envAcc.(*envAccessor)

	return &webhookHandler{
		eventType:               env.EventType,
		eventSource:             env.EventSource,
		extensionAttributesFrom: env.EventExtensionAttributesFrom,
		username:                env.BasicAuthUsername,
		password:                env.BasicAuthPassword,
		corsAllowOrigin:         env.CORSAllowOrigin,

		ceClient: ceClient,
		logger:   logging.FromContext(ctx),
		mt:       mt,
	}
}
