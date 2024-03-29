package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Filter is an addressable object that filters incoming events according
// to provided Common Language Expression
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FilterSpec      `json:"spec,omitempty"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

var (
	_ apis.Validatable = (*Filter)(nil)
	_ apis.Defaultable = (*Filter)(nil)

	_ v1alpha1.Reconcilable        = (*Filter)(nil)
	_ v1alpha1.AdapterConfigurable = (*Filter)(nil)
	_ v1alpha1.EventSender         = (*Filter)(nil)
	_ v1alpha1.EventSource         = (*Filter)(nil)
	_ v1alpha1.MultiTenant         = (*Filter)(nil)
)

// FilterSpec defines the desired state of the component.
type FilterSpec struct {
	Expression string `json:"expression"`

	// Sink is a reference to an object that will resolve to a domain name to use as the sink.
	Sink *duckv1.Destination `json:"sink"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FilterList is a list of component instances.
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Filter `json:"items"`
}
