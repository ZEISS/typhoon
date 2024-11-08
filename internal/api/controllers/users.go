package controllers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/models"
)

// CreateUserCommand ...
type CreateUserCommand struct {
	AccountID   uuid.UUID `json:"account_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// DeleteUserCommand ...
type DeleteUserCommand struct {
	UserID uuid.UUID `json:"user_id"`
}

// GetUserCredentialsQuery ...
type GetUserCredentialsQuery struct {
	UserID uuid.UUID `json:"user_id"`
}

// GetUserQuery ...
type GetUserQuery struct {
	UserID uuid.UUID `json:"user_id"`
}

// ListUsersQuery ...
type ListUsersQuery struct {
	AccountID uuid.UUID `json:"account_id"`
}

var _ UsersController = (*UsersControllerImpl)(nil)

// UsersController is the interface that wraps the methods to access users.
type UsersController interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, cmd CreateUserCommand) (models.User, error)
	// GetCredentials returns the credentials for a user.
	GetCredentials(ctx context.Context, query GetUserCredentialsQuery) ([]byte, error)
	// GetUser retrieves a user by its ID.
	GetUser(ctx context.Context, query GetUserQuery) (models.User, error)
	// ListUsers retrieves a list of users.
	ListUsers(ctx context.Context, query ListUsersQuery) (models.Pagination[models.User], error)
	// DeleteUser deletes a user by its ID.
	DeleteUser(ctx context.Context, cmd DeleteUserCommand) error
}

type UsersControllerImpl struct {
	db ports.Repositories
}

// NewUsersController ...
func NewUsersController(db ports.Repositories) *UsersControllerImpl {
	return &UsersControllerImpl{db}
}

// CreateUser ...
func (c *UsersControllerImpl) CreateUser(ctx context.Context, cmd CreateUserCommand) (models.User, error) {
	user := models.User{Name: cmd.Name, Description: cmd.Description}
	account := models.Account{ID: cmd.AccountID}

	err := c.db.GetAccount(ctx, &account)
	if err != nil {
		return user, err
	}
	user.Account = account

	pk, err := nkeys.CreateUser()
	if err != nil {
		return user, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return user, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return user, err
	}
	user.Key = models.NKey{ID: id, Seed: seed}

	if len(account.SigningKeyGroups) < 1 {
		return user, fmt.Errorf("account %s has no signing keys", account.ID)
	}

	ask, err := nkeys.FromSeed(account.SigningKeyGroups[0].Key.Seed)
	if err != nil {
		return user, err
	}

	askpk, err := ask.PublicKey()
	if err != nil {
		return user, err
	}

	// // Create a token for the user
	u := jwt.NewUserClaims(id)
	u.Name = cmd.Name
	u.IssuerAccount = account.Key.ID
	u.Issuer = askpk

	token, err := u.Encode(ask)
	if err != nil {
		return user, err
	}
	user.Token = models.Token{ID: id, Token: token}

	err = c.db.CreateUser(ctx, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUser ...
func (c *UsersControllerImpl) GetUser(ctx context.Context, query GetUserQuery) (models.User, error) {
	user := models.User{ID: query.UserID}

	err := c.db.GetUser(ctx, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetCredentials ...
func (c *UsersControllerImpl) GetCredentials(ctx context.Context, query GetUserCredentialsQuery) ([]byte, error) {
	user := models.User{ID: query.UserID}

	err := c.db.GetUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return user.Credentials()
}

// ListUsers ...
func (c *UsersControllerImpl) ListUsers(ctx context.Context, query ListUsersQuery) (models.Pagination[models.User], error) {
	results := models.Pagination[models.User]{}

	results, err := c.db.ListUsers(ctx, results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// DeleteUser ...
func (c *UsersControllerImpl) DeleteUser(ctx context.Context, cmd DeleteUserCommand) error {
	user := models.User{ID: cmd.UserID}

	err := c.db.DeleteUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
