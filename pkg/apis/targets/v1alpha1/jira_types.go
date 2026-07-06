package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JiraTarget is the Schema for the Jira Target.
type JiraTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JiraTargetSpec  `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable        = (*JiraTarget)(nil)
	_ v1alpha1.AdapterConfigurable = (*JiraTarget)(nil)
	_ v1alpha1.EventSource         = (*JiraTarget)(nil)
)

// JiraTargetSpec defines the desired state of the event target.
type JiraTargetSpec struct {
	Auth             JiraAuth                   `json:"auth"`
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	URL              string                     `json:"url"`
}

// JiraAuth contains Jira credentials.
type JiraAuth struct {
	Token SecretValueFromSource `json:"token"`
	User  string                `json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JiraTargetList is a list of event target instances.
type JiraTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []JiraTarget `json:"items"`
}
