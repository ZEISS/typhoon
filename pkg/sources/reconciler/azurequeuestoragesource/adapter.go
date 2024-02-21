

package azurequeuestoragesource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// adapterConfig contains properties used to configure the adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/azurequeuestoragesource-adapter"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

//github.com/zeiss/typhoon common.AdapterBuilder.
var _ comgithub.com/zeiss/typhoonent] = (*Reconciler)(nil)
github.com/zeiss/typhoon
// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AzureQueueStorageSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AzureQueueStorageSource) []corev1.EnvVar {
	storageQueueEnvs := common.MaybeAppendValueFromEnvVar([]corev1.EnvVar{}, "AZURE_ACCOUNT_KEY", o.Spec.AccountKey)

	envs := append(storageQueueEnvs, []corev1.EnvVar{
		{
			Name:  "AZURE_ACCOUNT_NAME",
			Value: o.Spec.AccountName,
		}, {
			Name:  "AZURE_QUEUE_NAME",
			Value: o.Spec.QueueName,
		},
	}...)

	if o.Spec.VisibilityTimeout != nil {
		envs = append(envs, corev1.EnvVar{
			Name:  "AZURE_VISIBILITY_TIMEOUT",
			Value: *o.Spec.VisibilityTimeout,
		})
	}

	return envs
}
