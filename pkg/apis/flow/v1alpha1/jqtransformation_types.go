package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type JQTransformation struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JQTransformationSpec `json:"spec"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

var (
	_ v1alpha1.Reconcilable        = (*JQTransformation)(nil)
	_ v1alpha1.AdapterConfigurable = (*JQTransformation)(nil)
	_ v1alpha1.EventSender         = (*JQTransformation)(nil)
)

// JQTransformationSpec defines the desired state of the component.
type JQTransformationSpec struct {
	// The query that gets passed to the JQ library
	Query string `json:"query"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Support sending to an event sink instead of replying.
	duckv1.SourceSpec `json:",inline"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JQTransformationList is a list of component instances.
type JQTransformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []JQTransformation `json:"items"`
}
