package v1alpha1

import (
	"context"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Accepted event types
const (
	// EventTypeSendGridEmailSend represents a task to send an email.
	EventTypeSendGridEmailSend = "com.zeiss.sendgrid.email.send"
	// EventTypeSendGridEmailSendResponse represents a response from the API after sending an email
	EventTypeSendGridEmailSendResponse = "com.zeiss.sendgrid.email.send.response"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*SendGridTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("SendGridTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*SendGridTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *SendGridTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *SendGridTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// AcceptedEventTypes implements IntegrationTarget.
func (*SendGridTarget) AcceptedEventTypes() []string {
	return []string{
		EventTypeSendGridEmailSend,
		EventTypeWildcard,
	}
}

// GetEventTypes implements EventSource.
func (*SendGridTarget) GetEventTypes() []string {
	return []string{
		EventTypeResponse,
	}
}

// AsEventSource implements EventSource.
func (t *SendGridTarget) AsEventSource() string {
	kind := strings.ToLower(t.GetGroupVersionKind().Kind)
	return "com.zeiss." + kind + "." + t.Namespace + "." + t.Name
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *SendGridTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *SendGridTarget) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *SendGridTarget) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
