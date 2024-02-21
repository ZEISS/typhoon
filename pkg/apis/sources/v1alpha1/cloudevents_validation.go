package v1alpha1

import (
	"context"
	"encoding/json"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (s *CloudEventsSource) Validate(ctx context.Context) *apis.FieldError {
	return s.Spec.Validate(ctx).ViaField("spec")
}

// Validate CloudEventsSource spec
func (s *CloudEventsSourceSpec) Validate(ctx context.Context) *apis.FieldError {
	if s.Credentials == nil {
		return nil
	}

	return s.Credentials.Validate(ctx).ViaField("credentials")
}

func (c *HTTPCredentials) Validate(ctx context.Context) *apis.FieldError {
	var errs *apis.FieldError

	if len(c.BasicAuths) != 0 {
		if _, err := json.Marshal(c.BasicAuths); err != nil {
			errs = errs.Also(apis.ErrInvalidValue(
				"basic authentication parameter cannot be marshaled into JSON", "basicAuths", err.Error()))
		}
	}

	return errs
}
