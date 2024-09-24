package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WorkerTransformation struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkerTransformationSpec `json:"spec"`
	Status v1alpha1.Status          `json:"status,omitempty"`
}

var (
	_ v1alpha1.Reconcilable        = (*WorkerTransformation)(nil)
	_ v1alpha1.AdapterConfigurable = (*WorkerTransformation)(nil)
	_ v1alpha1.EventSender         = (*WorkerTransformation)(nil)
)

// WorkerTransformationSpec defines the desired state of the component.
type WorkerTransformationSpec struct {
	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Support sending to an event sink instead of replying.
	duckv1.SourceSpec `json:",inline"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerTransformationList is a list of component instances.
type WorkerTransformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WorkerTransformation `json:"items"`
}
