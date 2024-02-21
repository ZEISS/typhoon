package v1alpha1

import (
	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*Transformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Transformation")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *Transformation) GetConditionSet() apis.ConditionSet {
	if t.Spec.Sink.Ref != nil || t.Spec.Sink.URI != nil {
		return v1alpha1.EventSenderConditionSet
	}
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *Transformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *Transformation) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetSink implements EventSender.
func (t *Transformation) GetSink() *duckv1.Destination {
	return &t.Spec.Sink
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *Transformation) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}
