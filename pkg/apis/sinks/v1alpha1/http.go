package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// Status defines the observed state of a component instance.
//
// +k8s:deepcopy-gen=true
type Status struct {
	duckv1.SourceStatus  `json:",inline"`
	duckv1.AddressStatus `json:",inline"`
}

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPSink is the Schema for an HTTP Sink.
type HTTPSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPSinkSpec `json:"spec"`
	Status Status       `json:"status,omitempty"`
}

// HTTPSinkSpec defines the desired state of an HTTP Sink.
type HTTPSinkSpec struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPSinkList contains a list of HTTPSink.
type HTTPSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HTTPSink `json:"items"`
}
