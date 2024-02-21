package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Managed event types
const (
	EventTypeDatadogMetric = "com.zeiss.datadog.metric.submit"
	EventTypeDatadogEvent  = "com.zeiss.datadog.event.post"
	EventTypeDatadogLog    = "com.zeiss.datadog.log.send"

	EventTypeDatadogResponse = "com.zeiss.datadog.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*DatadogTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("DatadogTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*DatadogTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *DatadogTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *DatadogTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*DatadogTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeDatadogMetric,
		EventTypeDatadogEvent,
		EventTypeDatadogLog,
	}
}

// GetEventTypes implements EventSource.
func (*DatadogTarget) GetEventTypes() []string {
	return []string{
		EventTypeDatadogResponse,
	}
}

// AsEventSource implements EventSource.
func (t *DatadogTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *DatadogTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *DatadogTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *DatadogTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
