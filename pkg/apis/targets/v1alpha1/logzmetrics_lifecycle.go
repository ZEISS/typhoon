package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Managed event types
const (
	EventTypeOpenTelemetryMetricsPush = "com.zeiss.opentelemetry.metrics.push"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*LogzMetricsTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("LogzTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*LogzMetricsTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *LogzMetricsTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *LogzMetricsTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*LogzMetricsTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeOpenTelemetryMetricsPush,
	}
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *LogzMetricsTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *LogzMetricsTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *LogzMetricsTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
