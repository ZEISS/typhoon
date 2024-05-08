package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

var _ UsersController = (*UsersControllerImpl)(nil)

// UsersController is the interface that wraps the methods to access users.
type UsersController interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, name string, accountId uuid.UUID) (*models.User, error)
	// GetCredentials returns the credentials for a user.
	GetCredentials(ctx context.Context, id uuid.UUID) ([]byte, error)
}

type UsersControllerImpl struct {
	db ports.Users
}

// NewUsersController ...
func NewUsersController(db ports.Users) *UsersControllerImpl {
	return &UsersControllerImpl{db}
}

// CreateUser ...
func (c *UsersControllerImpl) CreateUser(ctx context.Context, name string, accountId uuid.UUID) (*models.User, error) {
	return nil, nil

	// pk, err := nkeys.CreateUser()
	// if err != nil {
	// 	return nil, err
	// }

	// id, err := pk.PublicKey()
	// if err != nil {
	// 	return nil, err
	// }

	// seed, err := pk.Seed()
	// if err != nil {
	// 	return nil, err
	// }

	// ac, err := c.db.GetAccount(ctx, accountId)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(ac.SigningKeys) < 1 {
	// 	return nil, fmt.Errorf("account %s has no signing keys", ac.ID)
	// }

	// ask, err := nkeys.FromSeed(ac.SigningKeys[0].Seed)
	// if err != nil {
	// 	return nil, err
	// }

	// askpk, err := ask.PublicKey()
	// if err != nil {
	// 	return nil, err
	// }

	// // Create a token for the user
	// u := jwt.NewUserClaims(id)
	// u.Name = name
	// u.IssuerAccount = ac.KeyID
	// u.Issuer = askpk

	// token, err := u.Encode(ask)
	// if err != nil {
	// 	return nil, err
	// }

	// user := &models.User{
	// 	Name: name,
	// 	// SigningKeyGroupID: ac.SigningKeys[0].ID,
	// 	Key: models.NKey{
	// 		ID:   id,
	// 		Seed: seed,
	// 	},
	// 	Token: models.Token{
	// 		ID:    id,
	// 		Token: token,
	// 	},
	// }

	// err = c.db.CreateUser(ctx, user)
	// if err != nil {
	// 	return nil, err
	// }

	// return user, nil
}

// GetCredentials ...
func (c *UsersControllerImpl) GetCredentials(ctx context.Context, id uuid.UUID) ([]byte, error) {
	user, err := c.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// generate a creds formatted file that can be used by a NATS client
	creds, err := jwt.FormatUserConfig(user.Token.Token, user.Key.Seed)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
