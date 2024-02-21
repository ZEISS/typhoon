package webhooksource

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/zeiss/typhoon/pkg/apis/sources"
)

// NewAdapter satisfies pkgadapter.AdapterConstructor.
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	mt := &pkgadapter.MetricTag{
		ResourceGroup: sources.WebhookSourceResource.String(),
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

var _ pkgadapter.Adapter = (*webhookHandler)(nil)
