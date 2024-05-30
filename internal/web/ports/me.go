package ports

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
)

// Me ...
type Me interface {
	// GetProfile is a method that returns the profile of the current user
	GetProfile(ctx context.Context, user *adapters.GothUser) error
}
