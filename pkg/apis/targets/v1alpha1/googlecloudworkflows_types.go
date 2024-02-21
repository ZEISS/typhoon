package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudWorkflowsTarget is the Schema for an Google Cloud Workflows Target.
type GoogleCloudWorkflowsTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleCloudWorkflowsTargetSpec `json:"spec"`
	Status v1alpha1.Status                `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable           = (*GoogleCloudWorkflowsTarget)(nil)
	_ v1alpha1.AdapterConfigurable    = (*GoogleCloudWorkflowsTarget)(nil)
	_ v1alpha1.EventReceiver          = (*GoogleCloudWorkflowsTarget)(nil)
	_ v1alpha1.EventSource            = (*GoogleCloudWorkflowsTarget)(nil)
	_ v1alpha1.ServiceAccountProvider = (*GoogleCloudWorkflowsTarget)(nil)
)

// GoogleCloudWorkflowsTargetSpec defines the desired state of the event target.
type GoogleCloudWorkflowsTargetSpec struct {
	// GoogleCloudWorkflowsApiKey represents how GoogleCloudWorkflows credentials should be provided in the secret
	// Deprecated, please use "auth" object.
	Credentials *SecretValueFromSource `json:"credentialsJson,omitempty"`

	// Authentication methods common for all GCP targets.
	Auth *v1alpha1.GoogleCloudAuth `json:"auth,omitempty"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudWorkflowsTargetList is a list of event target instances.
type GoogleCloudWorkflowsTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GoogleCloudWorkflowsTarget `json:"items"`
}
