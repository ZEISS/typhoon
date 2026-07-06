package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	pkgapis "knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPPollerSource is the schema for the event source.
type HTTPPollerSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPPollerSourceSpec `json:"spec,omitempty"`
	Status v1alpha1.Status      `json:"status,omitempty"`
}

// Check the interfaces the event source should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*HTTPPollerSource)(nil)
	_ v1alpha1.AdapterConfigurable = (*HTTPPollerSource)(nil)
	_ v1alpha1.EventSource         = (*HTTPPollerSource)(nil)
	_ v1alpha1.EventSender         = (*HTTPPollerSource)(nil)
)

// HTTPPollerSourceSpec defines the desired state of the event source.
type HTTPPollerSourceSpec struct {
	duckv1.SourceSpec `json:",inline"`
	EventSource       *string                    `json:"eventSource,omitempty"`
	SkipVerify        *bool                      `json:"skipVerify,omitempty"`
	CACertificate     *string                    `json:"caCertificate,omitempty"`
	BasicAuthUsername *string                    `json:"basicAuthUsername,omitempty"`
	BasicAuthPassword *v1alpha1.ValueFromField   `json:"basicAuthPassword,omitempty"`
	Headers           map[string]string          `json:"headers,omitempty"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	EventType         string                     `json:"eventType"`
	Method            string                     `json:"method"`
	Endpoint          pkgapis.URL                `json:"endpoint"`
	Interval          apis.Duration              `json:"interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPPollerSourceList contains a list of event sources.
type HTTPPollerSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPPollerSource `json:"items"`
}
