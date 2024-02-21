// Package skip allows a Context to carry the intention to skip parts of the
// code execution. Mainly used to avoid variances while testing certain
// functions.
package skip

import "context"

type skipKey struct{}

// EnableSkip returns a copy of a parent context with skipping enabled.
func EnableSkip(ctx context.Context) context.Context {
	return context.WithValue(ctx, skipKey{}, struct{}{})
}

// Skip returns whether the given context has skipping enabled.
func Skip(ctx context.Context) bool {
	return ctx.Value(skipKey{}) != nil
}
