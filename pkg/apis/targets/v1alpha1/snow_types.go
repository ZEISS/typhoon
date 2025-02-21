package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceNowTarget ...
type ServiceNowTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceNowTargetSpec `json:"spec"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*SalesforceTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*SalesforceTarget)(nil)
	_ v1alpha1.EventReceiver       = (*SalesforceTarget)(nil)
	_ v1alpha1.EventSource         = (*SalesforceTarget)(nil)
)

// ServiceNowTargetSpec defines the desired state of the event target.
type ServiceNowTargetSpec struct {
	// Authentication information to interact with the Salesforce API.
	Auth SalesforceAuth `json:"auth"`

	// Instance is the ServiceNow instance to connect to.
	Instance string `json:"instance"`

	// Source is the source of the event.
	Source string `json:"source"`

	// EventOptions for targets
	// +optional
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// ServiceNowAuth contains the authentication information for the ServiceNow API.
type ServiceNowAuth struct {
	// User is the username used to authenticate with the ServiceNow API.
	User string `json:"user"`
	// Password is the password used to authenticate with the ServiceNow API.
	Password string `json:"password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceNowTargetList is a list of event target instances.
type ServiceNowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ServiceNowTarget `json:"items"`
}
