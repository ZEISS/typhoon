

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

// Managed event types
const (
	EventTypeGoogleCloudWorkflowsRun = "io.trigermesh.google.workflows.run"

	EventTypeGoogleCloudWorkflowsRunResponse = "io.trigermesh.google.workflows.run.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*GoogleCloudWorkflowsTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("GoogleCloudWorkflowsTarget")
}
github.com/zeiss/typhoon
// GetConditionSet implements duckv1.KRShaped.
func (*GoogleCloudWorkflowsTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *GoogleCloudWorkflowsTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *GoogleCloudWorkflowsTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*GoogleCloudWorkflowsTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeGoogleCloudWorkflowsRun,
	}
}

// GetEventTypes implements EventSource.
func (*GoogleCloudWorkflowsTarget) GetEventTypes() []string {
	return []string{
		EventTypeGoogleCloudWorkflowsRunResponse,
	}
}

// AsEventSource implements EventSource.
func (t *GoogleCloudWorkflowsTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *GoogleCloudWorkflowsTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *GoogleCloudWorkflowsTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth != nil && t.Spec.Auth.WantsOwnServiceAccount()
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *GoogleCloudWorkflowsTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	if t.Spec.Auth == nil {
		return []resource.ServiceAccountOption{}
	}
	return t.Spec.Auth.ServiceAccountOptions()
}

// SetDefaults implements apis.Defaultable
func (t *GoogleCloudWorkflowsTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *GoogleCloudWorkflowsTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
