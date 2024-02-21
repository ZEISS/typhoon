

package twiliosource

import (
	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// adapterConfig contains properties used to configure the adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/twiliosource-adapter"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ comgithub.com/zeiss/typhoonice] = (*Reconciler)(nil)
github.com/zeiss/typhoon
// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*servingv1.Service, error) {
	return common.NewAdapterKnService(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.VisibilityPublic,

		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}
