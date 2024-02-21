

package azureactivitylogssource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/sources/cloudevents"
)

const healthPortName = "health"

const defaultActivityLogsEventHubName = "insights-activity-logs"

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	/github.com/zeiss/typhoon
	/github.com/zeiss/typhoon Hubs instead of a source-specific image.
	Image stgithub.com/zeiss/typhoonRCE_IMAGE" default:"ghcr.io/zeiss/typhoon/azureeventhubssource-adapter"`
	/github.com/zeiss/typhoonng/metrics/tracing
	cgithub.com/zeiss/typhoon
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AzureActivityLogsSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),

		resource.Port(healthPortName, 8080),
		resource.StartupProbe("/health", healthPortName),
	), nil
}

// makeEventHubID returns the Resource ID of an Event Hubs instance based on
// the given Event Hubs namespace and Hub name.
func makeEventHubID(namespaceID *v1alpha1.AzureResourceID, hubName string) string {
	hubID := *namespaceID
	hubID.Namespace = namespaceID.ResourceName
	hubID.ResourceType = resourceTypeEventHubs
	hubID.ResourceName = hubName
	return hubID.String()
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AzureActivityLogsSource) []corev1.EnvVar {
	hubNamespaceID := o.Spec.Destination.EventHubs.NamespaceID
	eventHubName := defaultActivityLogsEventHubName
	if hubName := o.Spec.Destination.EventHubs.HubName; hubName != nil && *hubName != "" {
		eventHubName = *hubName
	}

	var hubEnvs []corev1.EnvVar
	if spAuth := o.Spec.Auth.ServicePrincipal; spAuth != nil {
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADTenantID, spAuth.TenantID)
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADClientID, spAuth.ClientID)
		hubEnvs = common.MaybeAppendValueFromEnvVar(hubEnvs, common.EnvAADClientSecret, spAuth.ClientSecret)
	}

	if o.Spec.Destination.EventHubs.ConsumerGroup != nil {
		hubEnvs = append(hubEnvs, corev1.EnvVar{
			Name:  common.EnvHubConsumerGroup,
			Value: *o.Spec.Destination.EventHubs.ConsumerGroup,
		})
	}

	ceType := v1alpha1.AzureEventType(sources.AzureServiceMonitor, v1alpha1.AzureActivityLogsActivityLogEventType)
	ceOverridesStr := cloudevents.OverridesJSON(o.Spec.CloudEventOverrides)

	return append(hubEnvs, []corev1.EnvVar{
		{
			Name:  common.EnvHubResourceID,
			Value: makeEventHubID(&hubNamespaceID, eventHubName),
		}, {
			Name:  common.EnvHubNamespace,
			Value: hubNamespaceID.ResourceName,
		}, {
			Name:  common.EnvHubName,
			Value: eventHubName,
		}, {
			Name:  common.EnvCESource,
			Value: o.AsEventSource(),
		}, {
			Name:  common.EnvCEType,
			Value: ceType,
		}, {
			Name:  adapter.EnvConfigCEOverrides,
			Value: ceOverridesStr,
		},
	}...,
	)
}
