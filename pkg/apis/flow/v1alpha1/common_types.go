package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/zeiss/typhoon/pkg/targets/adapter/cloudevents"
)

/* Provide common structs that are used by the targets such as secret definitions */

// ValueFromField is a struct field that can have its value either defined
// explicitly or sourced from another entity.
type ValueFromField struct {
	// Optional: no more than one of the following may be specified.

	// Field value.
	// +optional
	Value string `json:"value,omitempty"`
	// Field value from a Kubernetes Secret.
	// +optional
	ValueFromSecret *corev1.SecretKeySelector `json:"valueFromSecret,omitempty"`
	// Field value from a Kubernetes ConfigMap.
	// +optional
	ValueFromConfigMap *corev1.ConfigMapKeySelector `json:"valueFromConfigMap,omitempty"`
}

// EventOptions modifies CloudEvents management at Targets.
type EventOptions struct {
	// PayloadPolicy indicates if replies from the target should include
	// a payload if available. Possible values are:
	//
	// - always: will return a with the reply payload if avaliable.
	// - errors: will only reply with payload in case of an error.
	// - never: will not reply with payload.
	//
	// +optional
	PayloadPolicy *cloudevents.PayloadPolicy `json:"payloadPolicy,omitempty"`
}
