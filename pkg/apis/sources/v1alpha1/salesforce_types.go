package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SalesforceSource is the Schema for the event source.
type SalesforceSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SalesforceSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*SalesforceSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*SalesforceSource)(nil)
	_ v1alpha1.EventSource         = (*SalesforceSource)(nil)
	_ v1alpha1.EventSender         = (*SalesforceSource)(nil)
)

// SalesforceSourceSpec defines the desired state of the event source.
type SalesforceSourceSpec struct {
	duckv1.SourceSpec `json:",inline"`

	// Authentication method to interact with the Salesforce API.
	Auth SalesforceAuth `json:"auth"`

	// APIVersion at Salesforce.
	// +optional
	APIVersion *string `json:"apiVersion"`

	// Subscription to a Salesforce channel
	Subscription SalesforceSubscription `json:"subscription"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// SalesforceSubscription to connect to.
type SalesforceSubscription struct {
	Channel  string `json:"channel"`
	ReplayID *int   `json:"replayID,omitempty"`
}

// SalesforceAuth contains Salesforce credentials.
type SalesforceAuth struct {
	ClientID string                  `json:"clientID"`
	Server   string                  `json:"server"`
	User     string                  `json:"user"`
	CertKey  v1alpha1.ValueFromField `json:"certKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SalesforceSourceList contains a list of event sources.
type SalesforceSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SalesforceSource `json:"items"`
}
