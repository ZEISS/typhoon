package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/zeiss/typhoon/pkg/apis/common/v1alpha1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Function is an addressable object that executes function code.
type Function struct {
	Status            FunctionStatus `json:"status,omitempty"`
	Spec              FunctionSpec   `json:"spec,omitempty"`
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

var (
	_ v1alpha1.Reconcilable        = (*Function)(nil)
	_ v1alpha1.AdapterConfigurable = (*Function)(nil)
	_ v1alpha1.EventSource         = (*Function)(nil)
	_ v1alpha1.EventSender         = (*Function)(nil)
)

// FunctionSpec holds the desired state of the Function Specification
type FunctionSpec struct {
	duckv1.SourceSpec `json:",inline"`
	AdapterOverrides  *v1alpha1.AdapterOverrides `json:"adapterOverrides,omitempty"`
	Runtime           string                     `json:"runtime"`
	Entrypoint        string                     `json:"entrypoint"`
	Code              string                     `json:"code"`
	EventStore        EventStoreConnection       `json:"eventStore,omitempty"`
	ResponseIsEvent   bool                       `json:"responseIsEvent,omitempty"`
}

// EventStoreConnection contains the data to connect to
// an EventStore instance
type EventStoreConnection struct {
	// URI is the gRPC location to the EventStore
	URI string `json:"uri"`
}

// FunctionStatus defines the observed state of the Function.
type FunctionStatus struct {
	ConfigMap       *FunctionConfigMapIdentity `json:"configMap,omitempty"`
	v1alpha1.Status `json:",inline"`
}

// FunctionConfigMapIdentity represents the identity of the ConfigMap
// containing the code of a Function.
type FunctionConfigMapIdentity struct {
	Name            string `json:"name"`
	ResourceVersion string `json:"resourceVersion"`
}

// FunctionList is a list of Function resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Function `json:"items"`
}
