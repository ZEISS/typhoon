package ports

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
)

// Users ...
type Users interface {
	// GetUser is a method that returns a user by ID
	GetUser(ctx context.Context, user *adapters.GothUser) error
}
