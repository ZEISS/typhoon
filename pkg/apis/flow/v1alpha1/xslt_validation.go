package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (t *XSLTTransformation) Validate(ctx context.Context) *apis.FieldError {
	return t.Spec.Validate(ctx).ViaField("spec")
}

// Validate XSLT spec
func (s *XSLTTransformationSpec) Validate(ctx context.Context) *apis.FieldError {
	var errs *apis.FieldError

	if (s.AllowPerEventXSLT == nil || !*s.AllowPerEventXSLT) && !s.XSLT.IsInformed() {
		errs = errs.Also(apis.ErrGeneric("when XSLT is empty, per event XSLT must be allowed", "allowPerEventXSLT", "xslt"))
	}

	if err := s.XSLT.Validate(ctx); err != nil {
		errs = errs.Also(err.ViaField("XSLT"))
	}

	return errs
}
