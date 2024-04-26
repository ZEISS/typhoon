package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Users ...
type Users interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *models.User) error
	// UpdateUser updates an existing user.
	UpdateUser(ctx context.Context, user *models.User) error
	// GetUser returns the user with the given ID.
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	// ListUsers returns a list of users.
	ListUsers(ctx context.Context, pagination models.Pagination[*models.User]) (*models.Pagination[*models.User], error)
	// DeleteUser deletes the user with the given ID.
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
