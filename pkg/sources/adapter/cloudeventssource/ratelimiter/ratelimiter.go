package ratelimiter

import (
	"context"
	"net/http"
	"time"

	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

const (
	// token to be used globally for every request.
	globalToken = "global"
)

type rateLimiter struct {
	store limiter.Store
}

// New creates a new rate limiter.
func New(rps uint64) (cehttp.RateLimiter, error) {
	if store, err := memorystore.New(&memorystore.Config{
		Tokens:   rps,
		Interval: time.Second,
	}); err != nil {
		return nil, err
	} else {
		return &rateLimiter{
			store: store,
		}, nil
	}
}

// Allow checks if a request is allowed to pass the rate limiter filter.
func (rl *rateLimiter) Allow(ctx context.Context, _ *http.Request) (ok bool, reset uint64, err error) {
	_, _, reset, ok, err = rl.store.Take(ctx, globalToken)
	return ok, reset, err
}

// Close cleans up rate limiter resources.
func (rl *rateLimiter) Close(ctx context.Context) error {
	return rl.store.Close(ctx)
}
