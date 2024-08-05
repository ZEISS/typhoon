package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bridge is the Schema for the Bridge target.
type Bridge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BridgeSpec      `json:"spec"`
	Status v1alpha1.Status `json:"status,omitempty"`
}

// Check the interfaces Bridge should be implementing.
var (
	_ v1alpha1.Reconcilable = (*Bridge)(nil)
)

// BridgeSpec defines the desired state of the component.
type BridgeSpec struct {
	Components []Component `json:"components"`
}

// Component holds a component of a bridge.
type Component struct {
	// Object is the component object.
	// +optional
	// +nullable
	// +kubebuilder:pruning:PreserveUnknownFields
	// x-kubernetes-embedded-resource: false
	Object runtime.RawExtension `json:"object,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BridgeList is a list of component instances.
type BridgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Bridge `json:"items"`
}

// nolint:unused,deadcode
type isBridgeObject_BridgeObject interface {
	isBridgeObject_BridgeObject()
}

// BridgeObject_Transformation is a component of a bridge.
type BridgeObject_Transformation struct {
	Transformation TransformationSpec
}

// Check the interfaces BridgeObject_Transformation should be implementing.
func (*BridgeObject_Transformation) isBridgeObject_BridgeObject() {}
