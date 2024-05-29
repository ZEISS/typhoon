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
	EventTypeJiraIssueCreate = "com.zeiss.typhoon.jira.issue.create"
	EventTypeJiraIssueGet    = "com.zeiss.typhoon.jira.issue.get"
	EventTypeJiraCustom      = "com.zeiss.typhoon.jira.custom"

	EventTypeJiraIssue          = "com.zeiss.typhoon.jira.issue"
	EventTypeJiraCustomResponse = "com.zeiss.typhoon.jira.custom.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*JiraTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("JiraTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*JiraTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *JiraTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *JiraTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*JiraTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeJiraIssueCreate,
		EventTypeJiraIssueGet,
		EventTypeJiraCustom,
	}
}

// GetEventTypes implements EventSource.
func (*JiraTarget) GetEventTypes() []string {
	return []string{
		EventTypeJiraIssue,
		EventTypeJiraCustomResponse,
	}
}

// AsEventSource implements EventSource.
func (t *JiraTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss.typhoon." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *JiraTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *JiraTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *JiraTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
