package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSheetTarget is the Schema for an GoogleSheet Target.
type GoogleSheetTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleSheetTargetSpec `json:"spec"`
	Status v1alpha1.Status       `json:"status,omitempty"`
}

// Check the interfaces the event target should be implementing.
var (
	_ v1alpha1.Reconcilable           = (*GoogleSheetTarget)(nil)
	_ v1alpha1.AdapterConfigurable    = (*GoogleSheetTarget)(nil)
	_ v1alpha1.EventReceiver          = (*GoogleSheetTarget)(nil)
	_ v1alpha1.EventSource            = (*GoogleSheetTarget)(nil)
	_ v1alpha1.ServiceAccountProvider = (*GoogleSheetTarget)(nil)
)

// GoogleSheetTargetSpec defines the desired state of the event target.
type GoogleSheetTargetSpec struct {
	// GoogleSheet credential JSON for auth
	// Deprecated, please use "auth" object.
	GoogleServiceAccount *SecretValueFromSource `json:"googleServiceAccount"`

	// Authentication methods common for all GCP targets.
	Auth *v1alpha1.GoogleCloudAuth `json:"auth,omitempty"`

	// ID of Google a spreadsheet
	ID string `json:"id"`

	// DefaultPrefix is a pre-defined prefix for the individual sheets.
	DefaultPrefix string `json:"defaultPrefix"`

	// Adapter spec overrides parameters.
	// +optional
	AdapterOverrides *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSheetTargetList is a list of event target instances.
type GoogleSheetTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []GoogleSheetTarget `json:"items"`
}
