

package awscomprehendtarget

import (
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

const (
	envRegion              = "COMPREHEND_REGION"
	envLanguage            = "COMPREHEND_LANGUAGE"
	envEventsPayloadPolicy = "EVENTS_PAYLOAD_POLICY"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
tygithub.com/zeiss/typhoon
	// Configithub.com/zeiss/typhoonics/tracing
	ogithub.com/zeiss/typhoon
	/github.com/zeiss/typhoon
	Image string `default:"ghcr.io/zeiss/typhoon/awscomprehendtarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.AWSComprehendTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AWSComprehendTarget) []corev1.EnvVar {
	env := append(reconciler.MakeAWSAuthEnvVars(o.Spec.Auth),
		[]corev1.EnvVar{
			{
				Name:  envRegion,
				Value: o.Spec.Region,
			}, {
				Name:  envLanguage,
				Value: o.Spec.Language,
			},
		}...)

	if o.Spec.EventOptions != nil && o.Spec.EventOptions.PayloadPolicy != nil {
		env = append(env, corev1.EnvVar{
			Name:  envEventsPayloadPolicy,
			Value: string(*o.Spec.EventOptions.PayloadPolicy),
		})
	}

	return env
}
