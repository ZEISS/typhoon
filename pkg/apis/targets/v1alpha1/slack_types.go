package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SlackTarget defines the schema for the Slack target.
type SlackTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SlackTargetSpec `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// SlackTargetSpec defines the desired state of the event target.
type SlackTargetSpec struct {
	// Token for Slack App
	Token SecretValueFromSource `json:"token"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*SlackTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*SlackTarget)(nil)
	_ v1alpha1.EventReceiver       = (*SlackTarget)(nil)
	_ v1alpha1.EventSource         = (*SlackTarget)(nil)
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SlackTargetList is a list of event target instances.
type SlackTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SlackTarget `json:"items"`
}
