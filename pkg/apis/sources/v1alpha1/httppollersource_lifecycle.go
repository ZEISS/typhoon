package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*HTTPPollerSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("HTTPPollerSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *HTTPPollerSource) GetConditionSet() apis.ConditionSet {
	return v1alpha1.EventSenderConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *HTTPPollerSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements EventSender.
func (s *HTTPPollerSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *HTTPPollerSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// AsEventSource implements EventSource.
func (s *HTTPPollerSource) AsEventSource() string {
	if s.Spec.EventSource != nil {
		return *s.Spec.EventSource
	}

	sourceName := s.Name
	if s.Namespace != "" {
		sourceName = s.Namespace + "." + sourceName
	}

	return sourceName
}

// GetEventTypes implements EventSource.
func (s *HTTPPollerSource) GetEventTypes() []string {
	return []string{
		s.Spec.EventType,
	}
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *HTTPPollerSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (s *HTTPPollerSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *HTTPPollerSource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
