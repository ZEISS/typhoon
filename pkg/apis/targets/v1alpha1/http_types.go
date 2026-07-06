package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPTarget is the Schema for an HTTP Target.
type HTTPTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPTargetSpec  `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.EventSource         = (*HTTPTarget)(nil)
	_ v1alpha1.EventReceiver       = (*HTTPTarget)(nil)
	_ v1alpha1.Reconcilable        = (*HTTPTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*HTTPTarget)(nil)
)

// HTTPTargetSpec defines the desired state of the event target.
type HTTPTargetSpec struct {
	BasicAuthUsername *string                    `json:"basicAuthUsername,omitempty"`
	Headers           map[string]string          `json:"headers,omitempty"`
	SkipVerify        *bool                      `json:"skipVerify"`
	CACertificate     *string                    `json:"caCertificate"`
	BasicAuthPassword SecretValueFromSource      `json:"basicAuthPassword,omitempty"`
	OAuthClientID     *string                    `json:"oauthClientID,omitempty"`
	OAuthClientSecret SecretValueFromSource      `json:"oauthClientSecret,omitempty"`
	OAuthTokenURL     *string                    `json:"oauthTokenURL,omitempty"`
	OAuthScopes       *[]string                  `json:"oauthScopes,omitempty"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Response          HTTPEventResponse          `json:"response"`
	Method            string                     `json:"method"`
	Endpoint          apis.URL                   `json:"endpoint"`
}

// HTTPEventResponse for reply events context.
type HTTPEventResponse struct {
	// EventType for the reply.
	EventType string `json:"eventType"`

	// EventSource for the reply.
	EventSource string `json:"eventSource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPTargetList is a list of event target instances.
type HTTPTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []HTTPTarget `json:"items"`
}
