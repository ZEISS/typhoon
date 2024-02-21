package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (t *Transformation) Validate(ctx context.Context) *apis.FieldError {
	return t.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (ts *TransformationSpec) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
