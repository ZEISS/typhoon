package eventfilter

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Filtering result conditions.
const (
	PassFilter FilterResult = "pass"
	FailFilter FilterResult = "fail"
	NoFilter   FilterResult = "no_filter"
)

// FilterResult has the result of the filtering operation.
type FilterResult string

// And implements filter's logical conjunction.
func (x FilterResult) And(y FilterResult) FilterResult {
	if x == NoFilter {
		return y
	}
	if y == NoFilter {
		return x
	}
	if x == PassFilter && y == PassFilter {
		return PassFilter
	}
	return FailFilter
}

// Filter is an interface representing an event filter of the trigger filter.
type Filter interface {
	// Filter compute the predicate on the provided event and returns the result of the matching
	Filter(ctx context.Context, event cloudevents.Event) FilterResult
}

// Filters is a wrapper that runs each filter.
type Filters []Filter

// Filter applies filtering conditions on events.
func (filters Filters) Filter(ctx context.Context, event cloudevents.Event) FilterResult {
	res := NoFilter
	for _, f := range filters {
		res = res.And(f.Filter(ctx, event))
		// Short circuit to optimize it
		if res == FailFilter {
			return FailFilter
		}
	}
	return res
}

var _ Filter = Filters{}
