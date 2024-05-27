package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Users ...
type Users interface {
	// GetUser is a method that returns a user by ID
	GetUser(ctx context.Context, user *models.User) error
	// ListUsers is a method that returns a list of users
	ListUsers(ctx context.Context, pagination *models.Pagination[models.User]) error
	// CreateUser is a method that creates a user
	CreateUser(ctx context.Context, user *models.User) error
	// UpdateUser is a method that updates a user
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser is a method that deletes a user
	DeleteUser(ctx context.Context, user *models.User) error
}
