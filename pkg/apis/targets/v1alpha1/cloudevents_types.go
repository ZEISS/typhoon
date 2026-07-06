package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudEventsTarget is a gateway that produces received CloudEvents to a destination.
type CloudEventsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudEventsTargetSpec `json:"spec,omitempty"`
	Status v1alpha1.Status       `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*CloudEventsTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*CloudEventsTarget)(nil)
)

// CloudEventsTargetSpec defines the desired state of the event target.
type CloudEventsTargetSpec struct {
	Credentials      *CloudEventsCredentials    `json:"credentials,omitempty"`
	Path             *string                    `json:"path,omitempty"`
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Endpoint         apis.URL                   `json:"endpoint"`
}

// CloudEventsCredentials to be used when sending requests.
type CloudEventsCredentials struct {
	BasicAuth HTTPBasicAuth `json:"basicAuth,omitempty"`
}

// HTTPBasicAuth credentials.
type HTTPBasicAuth struct {
	Password v1alpha1.ValueFromField `json:"password"`
	Username string                  `json:"username"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudEventsTargetList is a list of event target instances.
type CloudEventsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []CloudEventsTarget `json:"items"`
}
