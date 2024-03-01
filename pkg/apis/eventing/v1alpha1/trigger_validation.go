package v1alpha1

import (
	"context"

	"github.com/zeiss/typhoon/pkg/brokers/config/broker"

	"knative.dev/pkg/apis"
)

// Validate the Trigger.
func (t *Trigger) Validate(ctx context.Context) *apis.FieldError {
	errs := t.Spec.Validate(apis.WithinSpec(ctx)).ViaField("spec")
	return errs
}

// Validate the TriggerSpec.
func (ts *TriggerSpec) Validate(ctx context.Context) (errs *apis.FieldError) {
	errs = ts.Broker.Validate(ctx).ViaField("broker")

	return errs.Also(
		broker.ValidateSubscriptionAPIFiltersList(ctx, ts.Filters).ViaField("filters"),
	).Also(
		ts.Target.Validate(ctx).ViaField("target"),
	).Also(
		ts.Delivery.Validate(ctx).ViaField("delivery"),
	)
}
