

package twiliotarget

import (
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

const (
	envTwilioSID           = "TWILIO_SID"
	envTwilioToken         = "TWILIO_TOKEN"
	envTwilioDefaultFrom   = "TWILIO_DEFAULT_FROM"
	envTwilioDefaultTo     = "TWILIO_DEFAULT_TO"
	envEventsPayloadPolicy = "EVENTS_PAYLOAD_POLICY"
)

// adapterConfig contains properties used to configure the target's adapter.
//github.com/zeiss/typhoonopulated by envconfig.
type adapgithub.com/zeiss/typhoon
	/github.com/zeiss/typhoonng/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/twiliotarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.TwilioTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.TwilioTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name: envTwilioSID,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.AccountSID.SecretKeyRef,
			},
		}, {
			Name: envTwilioToken,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.Token.SecretKeyRef,
			},
		}, {
			Name:  common.EnvBridgeID,
			Value: common.GetStatefulBridgeID(o),
		},
	}

	if o.Spec.DefaultPhoneFrom != nil {
		env = append(env, corev1.EnvVar{
			Name:  envTwilioDefaultFrom,
			Value: *o.Spec.DefaultPhoneFrom,
		})
	}

	if o.Spec.DefaultPhoneTo != nil {
		env = append(env, corev1.EnvVar{
			Name:  envTwilioDefaultTo,
			Value: *o.Spec.DefaultPhoneTo,
		})
	}

	if o.Spec.EventOptions != nil && o.Spec.EventOptions.PayloadPolicy != nil {
		env = append(env, corev1.EnvVar{
			Name:  envEventsPayloadPolicy,
			Value: string(*o.Spec.EventOptions.PayloadPolicy),
		})
	}

	return env
}
