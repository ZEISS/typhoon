package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// Supported event types
const (
	FilterGenericEventType = "com.zeiss.routing.filter"
)

// GetGroupVersionKind implements kmeta.OwnerRefable
func (*Filter) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("Filter")
}

// GetStatus implements duckv1.KRShaped.
func (f *Filter) GetStatus() *duckv1.Status {
	return &f.Status.Status
}

// GetConditionSet implements duckv1.KRShaped.
func (*Filter) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatusManager implements Reconcilable.
func (f *Filter) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: f.GetConditionSet(),
		Status:       &f.Status,
	}
}

// GetEventTypes implements EventSource.
func (*Filter) GetEventTypes() []string {
	return []string{
		FilterGenericEventType,
	}
}

// AsEventSource implements EventSource.
func (f *Filter) AsEventSource() string {
	return "filter/" + f.Name
}

// GetSink implements EventSender.
func (f *Filter) GetSink() *duckv1.Destination {
	return f.Spec.Sink
}

// IsMultiTenant implements MultiTenant.
func (*Filter) IsMultiTenant() bool {
	return true
}

// GetAdapterOverrides implements AdapterConfigurable.
func (f *Filter) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return f.Spec.AdapterOverrides
}
