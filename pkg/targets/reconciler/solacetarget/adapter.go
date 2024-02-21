package solacetarget

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
)

const (
	envURL        = "URL"
	envQueueName  = "QUEUE_NAME"
	envUsername   = "USERNAME"
	envPassword   = "PASSWORD"
	envCA         = "CA"
	envClientCert = "CLIENT_CERT"
	envClientKey  = "CLIENT_KEY"
	envSkipVerify = "SKIP_VERIFY"

	envSaslEnable = "SASL_ENABLE"
	envTLSEnable  = "TLS_ENABLE"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"gcr.io/triggermesh/solacetarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.SolaceTarget)

	var secretVolumes []corev1.Volume
	var secretVolMounts []corev1.VolumeMount

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
		resource.Volumes(secretVolumes...),
		resource.VolumeMounts(secretVolMounts...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.SolaceTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name:  envURL,
			Value: o.Spec.URL,
		},
		{
			Name:  envQueueName,
			Value: o.Spec.QueueName,
		},
		{
			Name:  "DISCARD_CE_CONTEXT",
			Value: strconv.FormatBool(o.Spec.DiscardCEContext),
		},
	}

	if o.Spec.Auth != nil {

		if o.Spec.Auth.SASLEnable != nil {
			env = append(env, corev1.EnvVar{
				Name:  envSaslEnable,
				Value: strconv.FormatBool(*o.Spec.Auth.SASLEnable),
			})
		}

		if o.Spec.Auth.TLSEnable != nil {
			env = append(env, corev1.EnvVar{
				Name:  envTLSEnable,
				Value: strconv.FormatBool(*o.Spec.Auth.TLSEnable),
			})
		}

		if o.Spec.Auth.Username != nil {
			env = append(env, corev1.EnvVar{
				Name:  envUsername,
				Value: *o.Spec.Auth.Username,
			})
		}

		if o.Spec.Auth.Password != nil {
			env = common.MaybeAppendValueFromEnvVar(
				env, envPassword, *o.Spec.Auth.Password,
			)
		}

		if o.Spec.Auth.TLS != nil {
			if o.Spec.Auth.TLS.CA != nil {
				env = common.MaybeAppendValueFromEnvVar(
					env, envCA, *o.Spec.Auth.TLS.CA,
				)
			}

			if o.Spec.Auth.TLS.ClientCert != nil {
				env = common.MaybeAppendValueFromEnvVar(
					env, envClientCert, *o.Spec.Auth.TLS.ClientCert,
				)
			}

			if o.Spec.Auth.TLS.ClientKey != nil {
				env = common.MaybeAppendValueFromEnvVar(
					env, envClientKey, *o.Spec.Auth.TLS.ClientKey,
				)
			}

			if o.Spec.Auth.TLS.SkipVerify != nil {
				env = append(env, corev1.EnvVar{
					Name:  envSkipVerify,
					Value: strconv.FormatBool(*o.Spec.Auth.TLS.SkipVerify),
				})
			}
		}
	}

	return env
}
