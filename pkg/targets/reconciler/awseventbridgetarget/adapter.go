

package awseventbridgetarget

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/targets/reconciler"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/awseventbridgetarget-adapter"`
}
github.com/zeiss/typhoon
// Verifygithub.com/zeiss/typhoon.AdapterBuilder.
vagithub.com/zeiss/typhoonv1.Service] = (*Reconciler)(nil)
github.com/zeiss/typhoon
// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.AWSEventBridgeTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AWSEventBridgeTarget) []corev1.EnvVar {
	return append(reconciler.MakeAWSAuthEnvVars(o.Spec.Auth),
		[]corev1.EnvVar{
			{
				Name:  common.EnvARN,
				Value: o.Spec.ARN,
			}, {
				Name:  "AWS_DISCARD_CE_CONTEXT",
				Value: strconv.FormatBool(o.Spec.DiscardCEContext),
			},
		}...)
}
