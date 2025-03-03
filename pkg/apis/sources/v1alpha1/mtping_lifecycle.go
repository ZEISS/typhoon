package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*PingSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("PingSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *PingSource) GetConditionSet() apis.ConditionSet {
	return v1alpha1.EventSenderConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *PingSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements EventSender.
func (s *PingSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *PingSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// AsEventSource implements EventSource.
func (s *PingSource) AsEventSource() string {
	sourceName := s.Name
	if s.Namespace != "" {
		sourceName = s.Namespace + "." + sourceName
	}

	return sourceName
}

// GetEventTypes implements EventSource.
func (s *PingSource) GetEventTypes() []string {
	return []string{
		s.Spec.EventType,
	}
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *PingSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (s *PingSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *PingSource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
