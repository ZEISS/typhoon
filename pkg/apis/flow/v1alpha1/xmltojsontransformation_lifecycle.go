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
	EventTypeXMLToJSONGenericResponse = "com.zeiss.xmltojsontransformation.error"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*XMLToJSONTransformation) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("XMLToJSONTransformation")
}

// GetConditionSet implements duckv1.KRShaped.
func (t *XMLToJSONTransformation) GetConditionSet() apis.ConditionSet {
	if t.Spec.Sink.Ref != nil || t.Spec.Sink.URI != nil {
		return v1alpha1.EventSenderConditionSet
	}
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *XMLToJSONTransformation) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *XMLToJSONTransformation) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetSink implements EventSender.
func (t *XMLToJSONTransformation) GetSink() *duckv1.Destination {
	return &t.Spec.Sink
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *XMLToJSONTransformation) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// SetDefaults implements apis.Defaultable
func (t *XMLToJSONTransformation) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (t *XMLToJSONTransformation) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
