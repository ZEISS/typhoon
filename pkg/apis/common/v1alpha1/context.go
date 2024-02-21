package v1alpha1

import "context"

type reconcilableInstanceKey struct{}

// WithReconcilable returns a copy of the parent context in which the value
// associated with the reconcilableInstanceKey is the given component instance.
func WithReconcilable(ctx context.Context, r Reconcilable) context.Context {
	return context.WithValue(ctx, reconcilableInstanceKey{}, r)
}

// ReconcilableFromContext returns the component instance stored in the context.
func ReconcilableFromContext(ctx context.Context) Reconcilable {
	if r, ok := ctx.Value(reconcilableInstanceKey{}).(Reconcilable); ok {
		return r
	}
	return nil
}
