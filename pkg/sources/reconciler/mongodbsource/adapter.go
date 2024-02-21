package mongodbsource

import (
	appsv1 "k8s.io/api/apps/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"

	corev1 "k8s.io/api/core/v1"
)

const (
	envMongoDBURI      = "MONGODB_URI"
	envMongoDBDatabase = "MONGODB_DATABASE"
	envMongoCollection = "MONGODB_COLLECTION"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
	// Container image
	Image string `default:"ghcr.io/zeiss/typhoon/mongodbsource-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.MongoDBSource)

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(MakeAppEnv(typedSrc)...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.MongoDBSource) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{Name: envMongoDBURI, Value: o.Spec.ConnectionString},
		{Name: envMongoDBDatabase, Value: o.Spec.Database},
		{Name: envMongoCollection, Value: o.Spec.Collection},
	}

	return env
}
