package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SalesforceTarget ...
type SalesforceTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SalesforceTargetSpec `json:"spec"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*SalesforceTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*SalesforceTarget)(nil)
	_ v1alpha1.EventReceiver       = (*SalesforceTarget)(nil)
	_ v1alpha1.EventSource         = (*SalesforceTarget)(nil)
)

// SalesforceTargetSpec defines the desired state of the event target.
type SalesforceTargetSpec struct {
	// Authentication information to interact with the Salesforce API.
	Auth SalesforceAuth `json:"auth"`

	// APIVersion at Salesforce. If not set the latest version will be used.
	// +optional
	APIVersion *string `json:"apiVersion"`

	// EventOptions for targets
	// +optional
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// SalesforceAuth contains OAuth JWT information to interact with the
// Salesforce API. See:
// https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_jwt_flow.htm
type SalesforceAuth struct {
	// ClientID for the Salesforce connected app.
	ClientID string `json:"clientID"`
	// Server points to the authorization URL.
	Server string `json:"server"`
	// User configuring the connected app.
	User string `json:"user"`
	// CertKey is the private key used to sign requests from the target.
	CertKey SecretValueFromSource `json:"certKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SalesforceTargetList is a list of event target instances.
type SalesforceTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SalesforceTarget `json:"items"`
}
