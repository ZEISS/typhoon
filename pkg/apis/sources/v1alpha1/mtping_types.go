package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PingSource is the schema for the event source.
type PingSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PingSourceSpec  `json:"spec,omitempty"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*WebhookSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*WebhookSource)(nil)
	_ v1alpha1.EventSource         = (*WebhookSource)(nil)
	_ v1alpha1.EventSender         = (*WebhookSource)(nil)
)

// PingSourceSpec defines the desired state of the event source.
type PingSourceSpec struct {
	duckv1.SourceSpec `json:",inline"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Schedule          string                     `json:"schedule,omitempty"`
	Data              string                     `json:"data,omitempty"`
	EventType         string                     `json:"eventType"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PingSourceList contains a list of event sources.
type PingSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PingSource `json:"items"`
}
