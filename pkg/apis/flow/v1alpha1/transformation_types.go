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

// Transformation allows to declaratively perform data transformations on CloudEvents.
type Transformation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TransformationSpec `json:"spec,omitempty"`
	Status v1alpha1.Status    `json:"status,omitempty"`
}

var (
	_ apis.Validatable = (*Transformation)(nil)
	_ apis.Defaultable = (*Transformation)(nil)

	_ v1alpha1.Reconcilable        = (*Transformation)(nil)
	_ v1alpha1.AdapterConfigurable = (*Transformation)(nil)
	_ v1alpha1.EventSender         = (*Transformation)(nil)
)

// TransformationSpec defines the desired state of the component.
type TransformationSpec struct {
	// Context contains Transformations that must be applied on CE Context
	Context []Transform `json:"context,omitempty"`
	// Data contains Transformations that must be applied on CE Data
	Data []Transform `json:"data,omitempty"`

	// Support sending to an event sink instead of replying.
	duckv1.SourceSpec `json:",inline"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// Transform describes transformation schemes for different CE types.
type Transform struct {
	Operation string `json:"operation"`
	Paths     []Path `json:"paths"`
}

// Path is a key-value pair that represents JSON object path
type Path struct {
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Separator string `json:"separator,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TransformationList is a list of component instances.
type TransformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Transformation `json:"items"`
}
