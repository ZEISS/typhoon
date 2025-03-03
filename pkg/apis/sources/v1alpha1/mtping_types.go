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
	// inherits duck/v1 SourceSpec, which currently provides:
	// * Sink - a reference to an object that will resolve to a domain name or
	//   a URI directly to use as the sink.
	// * CloudEventOverrides - defines overrides to control the output format
	//   and modifications of the event sent to the sink.
	duckv1.SourceSpec `json:",inline"`

	// Schedule is the cron schedule. Defaults to `* * * * *`.
	// +optional
	Schedule string `json:"schedule,omitempty"`

	// Data is data used as the body of the event posted to the sink. Default is empty.
	// Mutually exclusive with DataBase64.
	// +optional
	Data string `json:"data,omitempty"`

	// Value of the CloudEvents 'type' attribute to set on ingested events.
	// https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#type
	EventType string `json:"eventType"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PingSourceList contains a list of event sources.
type PingSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PingSource `json:"items"`
}
