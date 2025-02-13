package azservicebusqueuesource

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
	// Uses a common adapter for both Azure Service Bus sources instead of a source-specific image.
	Image string `envconfig:"AZURESERVICEBUSSOURCE_IMAGE" default:"gcr.io/zeiss/typhoon/azservicebussource-adapter"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AzureServiceBusQueueSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AzureServiceBusQueueSource) []corev1.EnvVar {
	var authEnvs []corev1.EnvVar
	if sasAuth := o.Spec.Auth.SASToken; sasAuth != nil {
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvServiceBusKeyName, sasAuth.KeyName)
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvServiceBusKeyValue, sasAuth.KeyValue)
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvServiceBusConnStr, sasAuth.ConnectionString)
	}
	if spAuth := o.Spec.Auth.ServicePrincipal; spAuth != nil {
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvAADTenantID, spAuth.TenantID)
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvAADClientID, spAuth.ClientID)
		authEnvs = common.MaybeAppendValueFromEnvVar(authEnvs, common.EnvAADClientSecret, spAuth.ClientSecret)
	}
	return append(authEnvs, corev1.EnvVar{
		Name:  common.EnvServiceBusEntityResourceID,
		Value: o.Spec.QueueID.String(),
	},
	)
}
