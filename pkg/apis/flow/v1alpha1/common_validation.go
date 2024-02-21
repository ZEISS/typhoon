package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// Validate makes sure that only one of the choices is properly informed.
func (v *ValueFromField) Validate(_ context.Context) *apis.FieldError {
	if v == nil {
		return nil
	}

	val := v.Value != ""
	secret := v.ValueFromSecret != nil && (v.ValueFromSecret.Name != "" || v.ValueFromSecret.Key != "")
	cm := v.ValueFromConfigMap != nil && (v.ValueFromConfigMap.Name != "" || v.ValueFromConfigMap.Key != "")

	if val && secret || val && cm || secret && cm {
		return apis.ErrMultipleOneOf("value", "valueFromSecret", "valueFromConfigMap")
	}

	if secret && (v.ValueFromSecret.Name == "" || v.ValueFromSecret.Key == "") {
		return apis.ErrMissingField("name", "key").ViaField("ValueFromSecret")
	}

	if cm && (v.ValueFromConfigMap.Name == "" || v.ValueFromConfigMap.Key == "") {
		return apis.ErrMissingField("name", "key").ViaField("ValueFromConfigMap")
	}

	return nil
}
