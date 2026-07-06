package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureServiceBusTopicSource is the Schema for the event source.
type AzureServiceBusTopicSource struct {
	Spec              AzureServiceBusTopicSourceSpec   `json:"spec,omitempty"`
	Status            AzureServiceBusTopicSourceStatus `json:"status,omitempty"`
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*AzureServiceBusTopicSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*AzureServiceBusTopicSource)(nil)
	_ v1alpha1.EventSource         = (*AzureServiceBusTopicSource)(nil)
	_ v1alpha1.EventSender         = (*AzureServiceBusTopicSource)(nil)
)

// AzureServiceBusTopicSourceSpec defines the desired state of the event source.
type AzureServiceBusTopicSourceSpec struct {
	duckv1.SourceSpec `json:",inline"`
	Auth              AzureAuth                  `json:"auth"`
	WebSocketsEnable  *bool                      `json:"webSocketsEnable,omitempty"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	TopicID           AzureResourceID            `json:"topicID"`
}

// AzureServiceBusTopicSourceStatus defines the observed state of the event source.
type AzureServiceBusTopicSourceStatus struct {
	SubscriptionID  *AzureResourceID `json:"subscriptionID,omitempty"`
	v1alpha1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureServiceBusTopicSourceList contains a list of event sources.
type AzureServiceBusTopicSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureServiceBusTopicSource `json:"items"`
}
