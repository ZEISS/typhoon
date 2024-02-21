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

// Splitter is an addressable object that splits incoming events according
// to provided specification.
type Splitter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SplitterSpec    `json:"spec,omitempty"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

var (
	_ apis.Validatable = (*Splitter)(nil)
	_ apis.Defaultable = (*Splitter)(nil)

	_ v1alpha1.Reconcilable        = (*Splitter)(nil)
	_ v1alpha1.AdapterConfigurable = (*Splitter)(nil)
	_ v1alpha1.EventSender         = (*Splitter)(nil)
	_ v1alpha1.EventSource         = (*Splitter)(nil)
	_ v1alpha1.MultiTenant         = (*Splitter)(nil)
)

// SplitterSpec defines the desired state of the component.
type SplitterSpec struct {
	Path      string              `json:"path"`
	CEContext CloudEventContext   `json:"ceContext"`
	Sink      *duckv1.Destination `json:"sink"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// CloudEventContext declares context attributes that will be propagated to resulting events.
type CloudEventContext struct {
	Type       string            `json:"type"`
	Source     string            `json:"source"`
	Extensions map[string]string `json:"extensions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SplitterList is a list of component instances.
type SplitterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Splitter `json:"items"`
}
