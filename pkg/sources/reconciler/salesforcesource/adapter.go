package salesforcesource

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

const (
	envSalesforceTokenURL     = "SALESFORCE_TOKEN_URL"
	envSalesforceClientID     = "SALESFORCE_CLIENT_ID"
	envSalesforceClientSecret = "SALESFORCE_CLIENT_SECRET"
	envSalesforceInstanceURL  = "SALESFORCE_URL"
	envSalesforceAPIVersion   = "SALESFORCE_API_VERSION"
	envSalesforceChannel      = "SALESFORCE_SUBCRIPTION_CHANNEL"
	envSalesforceReplayID     = "SALESFORCE_SUBCRIPTION_REPLAY_ID"
)

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/salesforcesource-adapter"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.SalesforceSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.SalesforceSource) []corev1.EnvVar {
	appEnv := []corev1.EnvVar{
		{
			Name:  envSalesforceClientID,
			Value: o.Spec.Auth.ClientID,
		},
		{
			Name:  envSalesforceClientSecret,
			Value: o.Spec.Auth.ClientSecret,
		},
		{
			Name:  envSalesforceTokenURL,
			Value: o.Spec.Auth.TokenURL,
		},
		{
			Name:  envSalesforceChannel,
			Value: o.Spec.Subscription.Channel,
		},
		{
			Name:  envSalesforceInstanceURL,
			Value: o.Spec.InstanceURL,
		},
	}

	if o.Spec.Subscription.ReplayID != nil {
		appEnv = append(appEnv, corev1.EnvVar{
			Name:  envSalesforceReplayID,
			Value: strconv.Itoa(*o.Spec.Subscription.ReplayID),
		})
	}

	if o.Spec.APIVersion != nil {
		appEnv = append(appEnv, corev1.EnvVar{
			Name:  envSalesforceAPIVersion,
			Value: *o.Spec.APIVersion,
		})
	}

	return appEnv
}
