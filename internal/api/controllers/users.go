package controllers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// UsersController ...
type UsersController struct {
	db ports.Repositories
}

// NewUsersController ...
func NewUsersController(db ports.Repositories) *UsersController {
	return &UsersController{db}
}

// CreateUser ...
func (c *UsersController) CreateUser(ctx context.Context, name string, accountId uuid.UUID) (*models.User, error) {
	pk, err := nkeys.CreateUser()
	if err != nil {
		return nil, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return nil, err
	}

	ac, err := c.db.GetAccount(ctx, accountId)
	if err != nil {
		return nil, err
	}

	if len(ac.SigningKeys) < 1 {
		return nil, fmt.Errorf("account %s has no signing keys", ac.ID)
	}

	ask, err := nkeys.FromSeed(ac.SigningKeys[0].Seed)
	if err != nil {
		return nil, err
	}

	// Create a token for the user
	u := jwt.NewUserClaims(id)
	u.Name = name
	u.IssuerAccount = ac.KeyID

	token, err := u.Encode(ask)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      name,
		AccountID: accountId,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
		Token: models.Token{
			ID:    id,
			Token: token,
		},
	}

	err = c.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
