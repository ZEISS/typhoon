package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (s *Splitter) Validate(ctx context.Context) *apis.FieldError {
	return s.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (ss *SplitterSpec) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
