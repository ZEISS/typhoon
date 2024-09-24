package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*WorkerTransformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("WorkerTransformation")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *WorkerTransformation) GetConditionSet() apis.ConditionSet {
	if t.Spec.Sink.Ref != nil || t.Spec.Sink.URI != nil {
		return v1alpha1.EventSenderConditionSet
	}
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *WorkerTransformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *WorkerTransformation) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetSink implements EventSender.
func (t *WorkerTransformation) GetSink() *duckv1.Destination {
	return &t.Spec.Sink
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *WorkerTransformation) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *WorkerTransformation) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *WorkerTransformation) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
