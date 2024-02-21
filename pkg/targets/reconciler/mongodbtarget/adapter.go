package mongodbtarget

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
	envServerURL           = "MONGODB_SERVER_URL"
	envDefaultDatabase     = "MONGODB_DEFAULT_DATABASE"
	envDefaultCollection   = "MONGODB_DEFAULT_COLLECTION"
	envEventsPayloadPolicy = "EVENTS_PAYLOAD_POLICY"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/mongodbtarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.MongoDBTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv constructs the environment variables for the adapter.
func MakeAppEnv(trg *v1alpha1.MongoDBTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name:  envDefaultDatabase,
			Value: trg.Spec.Database,
		},
		{
			Name:  envDefaultCollection,
			Value: trg.Spec.Collection,
		},
		{
			Name: envServerURL,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: trg.Spec.ConnectionString.ValueFromSecret,
			},
		},
	}

	if trg.Spec.EventOptions != nil && trg.Spec.EventOptions.PayloadPolicy != nil {
		env = append(env, corev1.EnvVar{
			Name:  envEventsPayloadPolicy,
			Value: string(*trg.Spec.EventOptions.PayloadPolicy),
		})
	}

	return env
}
