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
	EventTypeAzureEventHubsGenericResponse = "com.zeiss.azure.eventhubs.put.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*AzureEventHubsTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AzureEventHubsTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*AzureEventHubsTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *AzureEventHubsTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *AzureEventHubsTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*AzureEventHubsTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeWildcard,
	}
}

// GetEventTypes implements EventSource.
func (*AzureEventHubsTarget) GetEventTypes() []string {
	return []string{
		EventTypeAzureEventHubsGenericResponse,
	}
}

// AsEventSource implements EventSource.
func (t *AzureEventHubsTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *AzureEventHubsTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *AzureEventHubsTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *AzureEventHubsTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
