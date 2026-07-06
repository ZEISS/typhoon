package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SplunkTarget is the Schema for the event target.
type SplunkTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SplunkTargetSpec `json:"spec,omitempty"`
	Status v1alpha1.Status  `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*SplunkTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*SplunkTarget)(nil)
)

// SplunkTargetSpec defines the desired state of the event target.
type SplunkTargetSpec struct {
	Token            v1alpha1.ValueFromField    `json:"token"`
	Index            *string                    `json:"index,omitempty"`
	SkipTLSVerify    *bool                      `json:"skipTLSVerify,omitempty"`
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Endpoint         apis.URL                   `json:"endpoint"`
	DiscardCEContext bool                       `json:"discardCloudEventContext"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SplunkTargetList is a list of event target instances.
type SplunkTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SplunkTarget `json:"items"`
}
