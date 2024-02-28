package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NatsTarget is the Schema for an NatsTarget.
type NatsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NatsTargetSpec  `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*NatsTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*NatsTarget)(nil)
)

// NatsTargetSpec defines the desired state of the event target.
type NatsTargetSpec struct {
	// Subject where messages are produced.
	Subject string `json:"subject"`

	// URL of the Nats server.
	URL string `json:"url"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NatsTargetList is a list of event target instances.
type NatsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NatsTarget `json:"items"`
}
