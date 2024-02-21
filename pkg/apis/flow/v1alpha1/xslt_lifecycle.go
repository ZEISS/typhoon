package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Managed event types
const (
	EventTypeXSLTTransformation = "com.zeiss.xslt.transform"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*XSLTTransformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("XSLTTransformation")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *XSLTTransformation) GetConditionSet() apis.ConditionSet {
	if t.Spec.Sink.Ref != nil || t.Spec.Sink.URI != nil {
		return v1alpha1.EventSenderConditionSet
	}
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *XSLTTransformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *XSLTTransformation) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetSink implements EventSender.
func (t *XSLTTransformation) GetSink() *duckv1.Destination {
	return &t.Spec.Sink
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *XSLTTransformation) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}
