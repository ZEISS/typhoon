

package awscloudwatchsource

import (
	"encoding/json"
	"fmt"
	"time"

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

const (
	envRegion          = "AWS_REGION"
	envQueries         = "QUERIES"
	envPollingInterval = "POLLING_INTERVAL"
)

const defaultPollingInterval = 5 * time.Minute

cogithub.com/zeiss/typhoon
github.com/zeiss/typhoon
//github.com/zeiss/typhoon used to configure the source's adapter.
//github.com/zeiss/typhoon by envconfig.
type adapterConfig struct {
	// Container image
	Image string `default:"gcr.io/triggermesh/awscloudwatchsource"`
	// Configuration accessor for logging/metrics/tracing
	configs source.ConfigAccessor
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*appsv1.Deployment] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(src commonv1alpha1.Reconcilable, sinkURI *apis.URL) (*appsv1.Deployment, error) {
	typedSrc := src.(*v1alpha1.AWSCloudWatchSource)

	env, err := MakeAppEnv(typedSrc)
	if err != nil {
		return nil, fmt.Errorf("building adapter environment: %w", err)
	}

	return common.NewAdapterDeployment(src, sinkURI,
		resource.Image(r.adapterCfg.Image),

		resource.EnvVars(env...),
		resource.EnvVars(r.adapterCfg.configs.ToEnvVars()...),

		resource.Port(healthPortName, 8080),
		resource.StartupProbe("/health", healthPortName),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.AWSCloudWatchSource) ([]corev1.EnvVar, error) {
	var queries string
	if qs := o.Spec.MetricQueries; len(qs) > 0 {
		q, err := json.Marshal(qs)
		if err != nil {
			return nil, fmt.Errorf("serializing spec.metricQueries to JSON: %w", err)
		}
		queries = string(q)
	}

	pollingInterval := defaultPollingInterval
	if f := o.Spec.PollingInterval; f != nil && time.Duration(*f).Nanoseconds() > 0 {
		pollingInterval = time.Duration(*f)
	}

	return append(reconciler.MakeAWSAuthEnvVars(o.Spec.Auth),
		[]corev1.EnvVar{
			{
				Name:  envRegion,
				Value: o.Spec.Region,
			}, {
				Name:  envQueries,
				Value: queries,
			}, {
				Name:  envPollingInterval,
				Value: pollingInterval.String(),
			},
		}...,
	), nil
}
