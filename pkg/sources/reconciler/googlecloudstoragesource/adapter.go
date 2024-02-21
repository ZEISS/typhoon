

package googlecloudstoragesource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/sources/cloudevents"
)

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	// Uses the adapter for Google Cloud Pub/Sub instead of a source-specific image.
	Image string `envconfig:"GOOGLECLOUDPUBSUBSOURCE_IMAGE" default:"ghcr.io/zeiss/typhoon/googlecloudpubsubsource-adapter"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}github.com/zeiss/typhoon
github.com/zeiss/typhoon
//github.com/zeiss/typhoon common.AdapterBuilder.
vagithub.com/zeiss/typhoonDeployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.GoogleCloudStorageSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.GoogleCloudStorageSource) []corev1.EnvVar {
	// we rely on the source's status to persist the ID of the Pub/Sub subscription
	var subsName string
	if sn := o.Status.Subscription; sn != nil {
		subsName = sn.String()
	}

	envVar := []corev1.EnvVar{
		{
			Name:  "GCLOUD_PUBSUB_MESSAGE_PROCESSOR",
			Value: "gcs",
		}, {
			Name:  common.EnvGCloudPubSubSubscription,
			Value: subsName,
		}, {
			Name:  common.EnvCESource,
			Value: o.AsEventSource(),
		}, {
			Name:  adapter.EnvConfigCEOverrides,
			Value: cloudevents.OverridesJSON(o.Spec.CloudEventOverrides),
		},
	}

	if o.Spec.Auth.ServiceAccountKey != nil {
		envVar = append(envVar, common.MaybeAppendValueFromEnvVar([]corev1.EnvVar{}, common.EnvGCloudSAKey, *o.Spec.Auth.ServiceAccountKey)...)
	}

	return envVar
}
