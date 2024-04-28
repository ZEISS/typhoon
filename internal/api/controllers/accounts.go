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

// AccountsController ...
type AccountsController struct {
	db ports.Repositories
}

// NewAccountsController ...
func NewAccountsController(db ports.Repositories) *AccountsController {
	return &AccountsController{db}
}

// CreateAccount ...
func (c *AccountsController) CreateAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error) {
	operator, err := c.db.GetOperator(ctx, operatorID)
	if err != nil {
		return nil, err
	}

	pk, err := nkeys.CreateAccount()
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

	// Create a signing key for the account
	sk, err := nkeys.CreateAccount()
	if err != nil {
		return nil, err
	}

	spk, err := sk.PublicKey()
	if err != nil {
		return nil, err
	}

	skSeed, err := sk.Seed()
	if err != nil {
		return nil, err
	}

	if len(operator.SigningKeys) < 1 {
		return nil, fmt.Errorf("operator %s has no signing keys", operator.ID)
	}

	osk, err := nkeys.FromSeed(operator.SigningKeys[0].Seed)
	if err != nil {
		return nil, err
	}

	// Create a token for the account
	oc := jwt.NewAccountClaims(id)
	oc.Name = name
	oc.Issuer = operator.KeyID
	oc.SigningKeys.Add(spk)

	token, err := oc.Encode(osk)
	if err != nil {
		return nil, err
	}

	ac := &models.Account{
		Name:       name,
		OperatorID: operatorID,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
		SigningKeys: []models.NKey{
			{
				ID:   spk,
				Seed: skSeed,
			},
		},
		Token: models.Token{
			ID:    id,
			Token: token,
		},
	}

	err = c.db.CreateAccount(ctx, ac)
	if err != nil {
		return nil, err
	}

	return ac, nil
}
