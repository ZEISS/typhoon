package controllers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/utils"

	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// CreateOperatorAccountRequestObject ...
type UpdateOperatorAccountRequestObject = openapi.UpdateOperatorAccountRequestObject

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

// UpdateAccount ...
func (c *AccountsController) UpdateAccount(ctx context.Context, req UpdateOperatorAccountRequestObject) (*models.Account, error) {
	account, err := c.db.GetAccount(ctx, req.AccountId)
	if err != nil {
		return nil, err
	}

	// TODO: support multiple signing keys
	if len(account.Operator.SigningKeys) < 1 {
		return nil, fmt.Errorf("operator %s has no signing keys", account.Operator.KeyID)
	}

	osk, err := nkeys.FromSeed(account.Operator.SigningKeys[0].Seed)
	if err != nil {
		return nil, err
	}

	ac, err := jwt.DecodeAccountClaims(account.Token.Token)
	if err != nil {
		return nil, err
	}

	if len(*req.Body.Claims.Exports) > 0 {
		for _, e := range *req.Body.Claims.Exports {
			export := &jwt.Export{
				Name:                 utils.PtrStr(e.Name),
				Subject:              jwt.Subject(utils.PtrStr(e.Subject)),
				Type:                 jwt.ExportType(*e.Type),
				ResponseType:         jwt.ResponseType(*e.ResponseType),
				AccountTokenPosition: *e.AccountTokenPosition,
				Info: jwt.Info{
					Description: utils.PtrStr(e.Info.Description),
					InfoURL:     utils.PtrStr(e.Info.InfoUrl),
				},
			}

			ac.Exports.Add(export)
		}
	}

	token, err := ac.Encode(osk)
	if err != nil {
		return nil, err
	}
	account.Token.Token = token

	err = c.db.UpdateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
