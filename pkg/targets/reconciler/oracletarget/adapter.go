

package oracletarget

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

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/oracletarget-adapter"`
}

//github.com/zeiss/typhoon common.AdapterBuilder.
var _ comgithub.com/zeiss/typhoonice] = (*Reconciler)(nil)
github.com/zeiss/typhoon
// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.OracleTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.OracleTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name: "ORACLE_API_PRIVATE_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.OracleAPIPrivateKey.SecretKeyRef,
			},
		}, {
			Name: "ORACLE_API_PRIVATE_KEY_PASSPHRASE",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.OracleAPIPrivateKeyPassphrase.SecretKeyRef,
			},
		}, {
			Name: "ORACLE_API_PRIVATE_KEY_FINGERPRINT",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.OracleAPIPrivateKeyFingerprint.SecretKeyRef,
			},
		}, {
			Name:  "TENANT_OCID",
			Value: o.Spec.Tenancy,
		}, {
			Name:  "ORACLE_REGION",
			Value: o.Spec.Region,
		}, {
			Name:  "USER_OCID",
			Value: o.Spec.User,
		},
	}

	if o.Spec.OracleFunctionSpec != nil {
		env = append(env, corev1.EnvVar{
			Name:  "ORACLE_FUNCTION",
			Value: o.Spec.OracleFunctionSpec.Function,
		})
	}

	return env
}
