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
	EventTypeAzureSentinelTargetIncident = "com.zeiss.azure.sentinel.incident"

	EventTypeAzureSentinelTargetGenericResponse = "com.zeiss.azure.sentinel.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*AzureSentinelTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AzureSentinelTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*AzureSentinelTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// AcceptedEventTypes implements IntegrationTarget.
func (*AzureSentinelTarget) AcceptedEventTypes() []string {
	return []string{
		"*",
	}
}

// GetEventTypes implements EventSource.
func (*AzureSentinelTarget) GetEventTypes() []string {
	return []string{
		EventTypeAzureSentinelTargetGenericResponse,
	}
}

// GetStatus implements duckv1.KRShaped.
func (t *AzureSentinelTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *AzureSentinelTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AsEventSource implements EventSource.
func (t *AzureSentinelTarget) AsEventSource() string {
	return "sentinel." + t.Spec.ResourceGroup
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *AzureSentinelTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *AzureSentinelTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *AzureSentinelTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
