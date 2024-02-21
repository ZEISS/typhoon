package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"github.com/zeiss/typhoon/pkg/reconciler/resource"
)

// Accepted event types
const (
	// EventTypeGoogleSheetAppend represents a task to append a row to a sheet.
	EventTypeGoogleSheetAppend = "com.zeiss.googlesheet.append"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*GoogleSheetTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleSheetTarget")
}

// github.com/zeiss/typhoon.KRShaped.
func (*GoogleSheetTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *GoogleSheetTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *GoogleSheetTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*GoogleSheetTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeGoogleSheetAppend,
	}
}

// GetEventTypes implements EventSource.
func (*GoogleSheetTarget) GetEventTypes() []string {
	return []string{
		EventTypeResponse,
	}
}

// AsEventSource implements EventSource.
func (t *GoogleSheetTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *GoogleSheetTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *GoogleSheetTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth != nil && t.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *GoogleSheetTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	if t.Spec.Auth == nil {
		return []resource.ServiceAccountOption{}
	}
	return t.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (t *GoogleSheetTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *GoogleSheetTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
