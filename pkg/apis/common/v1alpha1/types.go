package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ValueFromField is a struct field that can have its value either defined
// explicitly or sourced from another entity.
//
// +k8s:deepcopy-gen=true
type ValueFromField struct {
	ValueFromSecret *corev1.SecretKeySelector `json:"valueFromSecret,omitempty"`
	Value           string                    `json:"value,omitempty"`
}

// AdapterOverrides are applied on top of the default adapter parameters.
//
// +k8s:deepcopy-gen=true
type AdapterOverrides struct {
	Public       *bool                        `json:"public,omitempty"`
	Resources    *corev1.ResourceRequirements `json:"resources,omitempty"`
	NodeSelector map[string]string            `json:"nodeSelector,omitempty"`
	Affinity     *corev1.Affinity             `json:"affinity,omitempty"`
	Labels       map[string]string            `json:"labels,omitempty"`
	Annotations  map[string]string            `json:"annotations,omitempty"`
	Tolerations  []corev1.Toleration          `json:"tolerations,omitempty"`
	Env          []corev1.EnvVar              `json:"env,omitempty"`
}

// GroupObject holds the API group object types.
//
// +k8s:deepcopy-gen=false
type GroupObject struct {
	Single, List runtime.Object
}
