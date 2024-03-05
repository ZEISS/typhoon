package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
// +k8s:deepcopy-gen=false
type Component struct {
	Object isBridgeObject_BridgeObject `json:"object"`
}

// GetObject return the object of the component.
func (c *Component) GetObject() isBridgeObject_BridgeObject {
	if c != nil {
		return c.Object
	}

	return nil
}

// GetTransformation returns the transformation of the component.
func (c *Component) GetTransformation() *BridgeObject_Transformation {
	if x, ok := c.GetObject().(*BridgeObject_Transformation); ok {
		return x
	}

	return nil
}

// DeepCopy is a helper function for deepcopy-gen.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)

	return out
}

// DeepCopyInto is a helper function for deepcopy-gen.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.GetTransformation() != nil {
		in.GetTransformation().DeepCopyInto(out.GetTransformation())
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BridgeList is a list of component instances.
type BridgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Bridge `json:"items"`
}

type isBridgeObject_BridgeObject interface {
	isBridgeObject_BridgeObject()
}

// BridgeObject_Transformation is a component of a bridge.
type BridgeObject_Transformation struct {
	Transformation TransformationSpec
}

// Check the interfaces BridgeObject_Transformation should be implementing.
func (*BridgeObject_Transformation) isBridgeObject_BridgeObject() {}
