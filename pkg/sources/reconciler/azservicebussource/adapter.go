package azservicebussource

import (
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// adapterConfig contains properties used to configure the source's adapter.
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
	typedSrc := src.(*v1alpha1.AzureServiceBusSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AzureServiceBusSource) []corev1.EnvVar {
	var webSocketsEnable bool
	if wss := o.Spec.WebSocketsEnable; wss != nil {
		webSocketsEnable = *wss
	}

	var appEnvs []corev1.EnvVar
	//TODO should this be here?
	if sasAuth := o.Spec.Auth.SASToken; sasAuth != nil {
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvServiceBusKeyName, sasAuth.KeyName)
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvServiceBusKeyValue, sasAuth.KeyValue)
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvServiceBusConnStr, sasAuth.ConnectionString)
	}
	if spAuth := o.Spec.Auth.ServicePrincipal; spAuth != nil {
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvAADTenantID, spAuth.TenantID)
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvAADClientID, spAuth.ClientID)
		appEnvs = common.MaybeAppendValueFromEnvVar(appEnvs, common.EnvAADClientSecret, spAuth.ClientSecret)
	}

	if o.Spec.MaxConcurrent != nil {
		appEnvs = append(appEnvs, corev1.EnvVar{
			Name:  common.EnvServiceBusMaxConcurrent,
			Value: strconv.Itoa(*o.Spec.MaxConcurrent),
		})
	}

	var resourceID string
	if o.Spec.TopicID != nil {
		if sID := o.Status.SubscriptionID; sID != nil {
			resourceID = sID.String()
		}
	} else if o.Spec.QueueID != nil {
		resourceID = o.Spec.QueueID.String()
	}
	return append(appEnvs, []corev1.EnvVar{{
		Name:  common.EnvServiceBusEntityResourceID,
		Value: resourceID,
	}, {
		Name:  common.EnvServiceBusWebSocketsEnable,
		Value: strconv.FormatBool(webSocketsEnable),
	}}...)
}
