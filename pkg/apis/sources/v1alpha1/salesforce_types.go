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
	Subscription      SalesforceSubscription     `json:"subscription"`
	APIVersion        *string                    `json:"apiVersion"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Auth              SalesforceAuth             `json:"auth"`
	InstanceURL       string                     `json:"instanceURL"`
}

// SalesforceSubscription to connect to.
type SalesforceSubscription struct {
	ReplayID *int   `json:"replayID,omitempty"`
	Channel  string `json:"channel"`
}

// SalesforceAuth contains Salesforce credentials.
type SalesforceAuth struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	TokenURL     string `json:"tokenURL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SalesforceSourceList contains a list of event sources.
type SalesforceSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SalesforceSource `json:"items"`
}
