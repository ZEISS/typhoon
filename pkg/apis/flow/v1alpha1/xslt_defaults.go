package v1alpha1

import (
	"context"
)

// SetDefaults implements apis.Defaultable
func (t *XSLTTransformation) SetDefaults(ctx context.Context) {
	if t != nil {
		t.Spec.SetDefaults(ctx)
	}
}

// SetDefaults implements apis.Defaultable
func (s *XSLTTransformationSpec) SetDefaults(ctx context.Context) {
	if s != nil && s.AllowPerEventXSLT == nil {
		f := false
		s.AllowPerEventXSLT = &f
	}
}
