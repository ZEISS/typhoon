package transformation

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

const (
	envTransformationCtx  = "TRANSFORMATION_CONTEXT"
	envTransformationData = "TRANSFORMATION_DATA"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/transformation-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.Transformation)

	return common.NewAdapterKnService(trg, sinkURI,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.Transformation) []corev1.EnvVar {
	var trnContext string
	if b, err := json.Marshal(o.Spec.Context); err == nil {
		trnContext = string(b)
	}

	var trnData string
	if b, err := json.Marshal(o.Spec.Data); err == nil {
		trnData = string(b)
	}

	return []corev1.EnvVar{
		{
			Name:  envTransformationCtx,
			Value: trnContext,
		},
		{
			Name:  envTransformationData,
			Value: trnData,
		},
	}
}
