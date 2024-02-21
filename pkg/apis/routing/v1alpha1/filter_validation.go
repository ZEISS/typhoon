package v1alpha1

import (
	"context"
	"fmt"

	"knative.dev/pkg/apis"

	"github.com/zeiss/typhoon/pkg/routing/eventfilter/cel"
)

// Validate implements apis.Validatable
func (f *Filter) Validate(ctx context.Context) *apis.FieldError {
	return f.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (fs *FilterSpec) Validate(ctx context.Context) *apis.FieldError {
	if fs.Expression == "" {
		return apis.ErrMissingField("Expression")
	}
	if fs.Sink == nil {
		return apis.ErrMissingField("Sink")
	}
	if _, err := cel.CompileExpression(fs.Expression); err != nil {
		return apis.ErrInvalidValue(fmt.Sprintf("Cannot compile expression: %v", err), "Expression")
	}
	return nil
}
