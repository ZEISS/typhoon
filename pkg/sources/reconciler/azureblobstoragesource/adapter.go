

package azureblobstoragesource

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

const healthPortName = "health"

const envMessageProcessor = "EVENTHUB_MESSAGE_PROCESSOR"

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	// Uses the adapter for Azure Event Hubs instead of a source-specific image.
	Igithub.com/zeiss/typhoonHUBSSOURCE_IMAGE" default:"ghcr.io/zeiss/typhoon/azureeventhubssource-adapter"`
	// Configithub.com/zeiss/typhoonics/tracing
	cgithub.com/zeiss/typhoon
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AzureBlobStorageSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),

		resource.Port(healthPortName, 8080),
		resource.StartupProbe("/health", healthPortName),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AzureBlobStorageSource) []corev1.EnvVar {
	// the user may or may not provide an Event Hub name in the source's
	// spec, so the source's status is unfortunately our only source of
	// truth here
	var hubResID string
	var hubName string
	if ehID := o.Status.EventHubID; ehID != nil {
		hubResID = ehID.String()
		hubName = ehID.ResourceName
	}

	var hubEnvs []corev1.EnvVar
	if spAuth := o.Spec.Auth.ServicePrincipal; spAuth != nil {
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADTenantID, spAuth.TenantID)
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADClientID, spAuth.ClientID)
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADClientSecret, spAuth.ClientSecret)
	}

	if o.Spec.Endpoint.EventHubs.ConsumerGroup != nil {
		hubEnvs = append(hubEnvs, corev1.EnvVar{
			Name:  common.EnvHubConsumerGroup,
			Value: *o.Spec.Endpoint.EventHubs.ConsumerGroup,
		})
	}

	return append(hubEnvs, []corev1.EnvVar{
		{
			Name:  common.EnvHubResourceID,
			Value: hubResID,
		}, {
			Name:  common.EnvHubNamespace,
			Value: o.Spec.Endpoint.EventHubs.NamespaceID.ResourceName,
		}, {
			Name:  common.EnvHubName,
			Value: hubName,
		}, {
			Name:  envMessageProcessor,
			Value: "eventgrid",
		},
	}...,
	)
}
