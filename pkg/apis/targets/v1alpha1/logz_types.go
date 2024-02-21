package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogzTarget is the Schema for the Logz Target.
type LogzTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogzTargetSpec  `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*LogzTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*LogzTarget)(nil)
	_ v1alpha1.EventSource         = (*LogzTarget)(nil)
)

// LogzTargetSpec defines the desired state of the event target.
type LogzTargetSpec struct {
	// ShippingToken defines the API token.
	ShippingToken SecretValueFromSource `json:"shippingToken"`

	// LogsListenerURL Defines the Log listener URL
	LogsListenerURL string `json:"logsListenerURL"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogzTargetList is a list of event target instances.
type LogzTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []LogzTarget `json:"items"`
}
