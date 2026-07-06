package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureServiceBusQueueSource is the Schema for the event source.
type AzureServiceBusQueueSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureServiceBusQueueSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status                `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*AzureServiceBusQueueSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*AzureServiceBusQueueSource)(nil)
	_ v1alpha1.EventSource         = (*AzureServiceBusQueueSource)(nil)
	_ v1alpha1.EventSender         = (*AzureServiceBusQueueSource)(nil)
)

// AzureServiceBusQueueSourceSpec defines the desired state of the event source.
type AzureServiceBusQueueSourceSpec struct {
	duckv1.SourceSpec `json:",inline"`
	Auth              AzureAuth                  `json:"auth"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	QueueID           AzureResourceID            `json:"queueID"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureServiceBusQueueSourceList contains a list of event sources.
type AzureServiceBusQueueSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureServiceBusQueueSource `json:"items"`
}
