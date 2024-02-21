package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*Synchronizer) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Synchronizer")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *Synchronizer) GetConditionSet() apis.ConditionSet {
	if s.Spec.Sink.Ref != nil || s.Spec.Sink.URI != nil {
		return v1alpha1.EventSenderConditionSet
	}
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *Synchronizer) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetStatusManager implements Reconcilable.
func (s *Synchronizer) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// GetSink implements EventSender.
func (s *Synchronizer) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *Synchronizer) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (s *Synchronizer) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *Synchronizer) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
