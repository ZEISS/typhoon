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
	KafkaSourceEventType = "com.zeiss.kafka.event"
)

// GetEventTypes implements EventSource.
func (*KafkaSource) GetEventTypes() []string {
	return []string{
		KafkaSourceEventType,
	}
}

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*KafkaSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("KafkaSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (*KafkaSource) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *KafkaSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements EventSender.
func (s *KafkaSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *KafkaSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// AsEventSource implements EventSource.
func (s *KafkaSource) AsEventSource() string {
	return s.Spec.Topic
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *KafkaSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (s *KafkaSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *KafkaSource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
