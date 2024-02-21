package logzmetricstarget

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/apis"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	commonv1alpha1 "github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	common "github.com/zeiss/typhoon/pkg/reconciler"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

const (
	envCortexEndpoint           = "OPENTELEMETRY_CORTEX_ENDPOINT"
	envCortexBearerToken        = "OPENTELEMETRY_CORTEX_BEARER_TOKEN"
	envOpenTelemetryInstruments = "OPENTELEMETRY_INSTRUMENTS"

	envEventsPayloadPolicy = "EVENTS_PAYLOAD_POLICY"
)

// adapterConfig contains properties used to configure the target's adapter.
// Public fields are automatically populated by envconfig.
type adapterConfig struct {
	// Configuration accessor for logging/metrics/tracing
	obsConfig source.ConfigAccessor
	// Container image
	// Uses the adapter for OpenTelemetry instead of a source-specific image.
	Image string `envconfig:"OPENTELEMETRYTARGET_IMAGE" default:"ghcr.io/zeiss/typhoon/opentelemetrytarget-adapter"`
}

// Verify that Reconciler implements common.AdapterBuilder.
var _ common.AdapterBuilder[*servingv1.Service] = (*Reconciler)(nil)

// BuildAdapter implements common.AdapterBuilder.
func (r *Reconciler) BuildAdapter(trg commonv1alpha1.Reconcilable, _ *apis.URL) (*servingv1.Service, error) {
	typedTrg := trg.(*v1alpha1.LogzMetricsTarget)

	return common.NewAdapterKnService(trg, nil,
		resource.Image(r.adapterCfg.Image),
		resource.EnvVars(MakeAppEnv(typedTrg)...),
		resource.EnvVars(r.adapterCfg.obsConfig.ToEnvVars()...),
	), nil
}

// MakeAppEnv extracts environment variables from the object.
// Exported to be used in external tools for local test environments.
func MakeAppEnv(o *v1alpha1.LogzMetricsTarget) []corev1.EnvVar {
	env := []corev1.EnvVar{
		{
			Name: envCortexBearerToken,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: o.Spec.Connection.Token.SecretKeyRef,
			},
		}, {
			Name:  envCortexEndpoint,
			Value: o.Spec.Connection.ListenerURL,
		}, {
			Name:  common.EnvBridgeID,
			Value: common.GetStatefulBridgeID(o),
		},
	}

	if instruments, err := json.Marshal(o.Spec.Instruments); err == nil {
		env = append(env, corev1.EnvVar{
			Name:  envOpenTelemetryInstruments,
			Value: string(instruments),
		})
	}

	if o.Spec.EventOptions != nil && o.Spec.EventOptions.PayloadPolicy != nil {
		env = append(env, corev1.EnvVar{
			Name:  envEventsPayloadPolicy,
			Value: string(*o.Spec.EventOptions.PayloadPolicy),
		})
	}

	return env
}
