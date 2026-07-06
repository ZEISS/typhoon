package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebhookSource is the schema for the event source.
type WebhookSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status   `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*WebhookSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*WebhookSource)(nil)
	_ v1alpha1.EventSource         = (*WebhookSource)(nil)
	_ v1alpha1.EventSender         = (*WebhookSource)(nil)
)

// WebhookSourceSpec defines the desired state of the event source.
type WebhookSourceSpec struct {
	duckv1.SourceSpec        `json:",inline"`
	EventSource              *string                          `json:"eventSource,omitempty"`
	EventExtensionAttributes *WebhookEventExtensionAttributes `json:"eventExtensionAttributes,omitempty"`
	BasicAuthUsername        *string                          `json:"basicAuthUsername,omitempty"`
	BasicAuthPassword        *v1alpha1.ValueFromField         `json:"basicAuthPassword,omitempty"`
	CORSAllowOrigin          *string                          `json:"corsAllowOrigin,omitempty"`
	AdapterOverrides         *v1alpha1.AdapterOverrides       `json:"adapterOverrides,omitempty"`
	EventType                string                           `json:"eventType"`
}

// WebhookEventExtensionAttributes sets the policy for converting HTTP data into.
type WebhookEventExtensionAttributes struct {
	// From informs HTTP elements that will be converted into CloudEvents attributes
	// +optional
	From []string `json:"from,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebhookSourceList contains a list of event sources.
type WebhookSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebhookSource `json:"items"`
}
