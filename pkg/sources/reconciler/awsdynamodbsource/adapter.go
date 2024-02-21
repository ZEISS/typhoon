

package awsdynamodbsource

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
	"github.com/zeiss/typhoon/pkg/sources/reconciler"
)

const healthPortName = "health"

// adapterConfig contains properties used to configure the source's adapter.
// These are automatically populated by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/awsdynamodbsource"`
	// Configuration accessor for logging/metrics/tracing
	cgithub.com/zeiss/typhoon
}github.com/zeiss/typhoon
github.com/zeiss/typhoon
//github.com/zeiss/typhoon common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AWSDynamoDBSource)

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
func MakeAppEnv(o *v1alpha1.AWSDynamoDBSource) []corev1.EnvVar {
	return append(reconciler.MakeAWSAuthEnvVars(o.Spec.Auth),
		[]corev1.EnvVar{
			{
				Name:  common.EnvARN,
				Value: o.Spec.ARN.String(),
			},
		}...,
	)
}
