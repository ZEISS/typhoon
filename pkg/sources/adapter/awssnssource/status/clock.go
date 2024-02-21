package status

import (
	"context"

	"knative.dev/pkg/apis"
)

type clockKey struct{}

// Clock can override the timestamp returned by time.Now().
type Clock interface {
	Now() apis.VolatileTime
}

// WithClock returns a copy of the parent context in which the value
// associated with the clock key is the given Clock.
func WithClock(ctx context.Context, c Clock) context.Context {
	return context.WithValue(ctx, clockKey{}, c)
}

// ClockFromContext returns the Clock stored in the context.
func ClockFromContext(ctx context.Context) Clock {
	if c, ok := ctx.Value(clockKey{}).(Clock); ok {
		return c
	}
	return nil
}
