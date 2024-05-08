package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"

	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// AccountsController is the interface that wraps the methods to access accounts.
type AccountsController interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error)
	// UpdateAccount updates an account.
	UpdateAccount(ctx context.Context, req UpdateOperatorAccountRequestObject) (*models.Account, error)
	// DeleteToken deletes a token.
	DeleteToken(ctx context.Context, accountID uuid.UUID) error
	// CreateSigningKeyGroup creates a new signing key group.
	CreateSigningKeyGroup(ctx context.Context) (*models.Account, error)
	// ListSigningKeys of an account.
	ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error)
}

// CreateOperatorAccountRequestObject ...
type UpdateOperatorAccountRequestObject = openapi.UpdateOperatorAccountRequestObject

type accountsController struct {
	db ports.Repositories
}

// NewAccountsController ...
func NewAccountsController(db ports.Repositories) *accountsController {
	return &accountsController{db}
}

// CreateAccount ...
func (c *accountsController) CreateAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error) {
	return nil, nil
	// operator, err := c.db.GetOperator(ctx, operatorID)
	// if err != nil {
	// 	return nil, err
	// }

	// pk, err := nkeys.CreateAccount()
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

	// // Create a signing key for the account
	// sk, err := nkeys.CreateAccount()
	// if err != nil {
	// 	return nil, err
	// }

	// spk, err := sk.PublicKey()
	// if err != nil {
	// 	return nil, err
	// }

	// skSeed, err := sk.Seed()
	// if err != nil {
	// 	return nil, err
	// }

	// if len(operator.SigningKeys) < 1 {
	// 	return nil, fmt.Errorf("operator %s has no signing keys", operator.ID)
	// }

	// osk, err := nkeys.FromSeed(operator.SigningKeys[0].Seed)
	// if err != nil {
	// 	return nil, err
	// }

	// // Create a token for the account
	// ac := jwt.NewAccountClaims(id)
	// ac.Name = name
	// ac.Issuer = operator.KeyID
	// ac.SigningKeys.Add(spk)

	// token, err := ac.Encode(osk)
	// if err != nil {
	// 	return nil, err
	// }

	// account := &models.Account{
	// 	Name:       name,
	// 	OperatorID: operatorID,
	// 	Key: models.NKey{
	// 		ID:   id,
	// 		Seed: seed,
	// 	},
	// 	// SigningKeys: []models.NKey{
	// 	// 	{
	// 	// 		ID:   spk,
	// 	// 		Seed: skSeed,
	// 	// 	},
	// 	// },
	// 	Token: models.Token{
	// 		ID:    id,
	// 		Token: token,
	// 	},
	// }

	// err = c.db.CreateAccount(ctx, account)
	// if err != nil {
	// 	return nil, err
	// }

	// return account, nil
}

// UpdateAccount ...
func (c *accountsController) UpdateAccount(ctx context.Context, req UpdateOperatorAccountRequestObject) (*models.Account, error) {
	return nil, nil
	// account, err := c.db.GetAccount(ctx, req.AccountId)
	// if err != nil {
	// 	return nil, err
	// }

	// // TODO: support multiple signing keys
	// if len(account.Operator.SigningKeys) < 1 {
	// 	return nil, fmt.Errorf("operator %s has no signing keys", account.Operator.KeyID)
	// }

	// osk, err := nkeys.FromSeed(account.Operator.SigningKeys[0].Seed)
	// if err != nil {
	// 	return nil, err
	// }

	// ac, err := jwt.DecodeAccountClaims(account.Token.Token)
	// if err != nil {
	// 	return nil, err
	// }
	// ac.Exports = make([]*jwt.Export, 0)

	// if len(*req.Body.Claims.Exports) > 0 {
	// 	for _, e := range *req.Body.Claims.Exports {
	// 		export := &jwt.Export{
	// 			Name:                 utils.PtrStr(e.Name),
	// 			Subject:              jwt.Subject(utils.PtrStr(e.Subject)),
	// 			Type:                 jwt.ExportType(*e.Type),
	// 			ResponseType:         jwt.ResponseType(*e.ResponseType),
	// 			AccountTokenPosition: *e.AccountTokenPosition,
	// 			Info: jwt.Info{
	// 				Description: utils.PtrStr(e.Info.Description),
	// 				InfoURL:     utils.PtrStr(e.Info.InfoUrl),
	// 			},
	// 		}

	// 		ac.Exports.Add(export)
	// 	}
	// }

	// token, err := ac.Encode(osk)
	// if err != nil {
	// 	return nil, err
	// }
	// account.Token.Token = token

	// err = c.db.UpdateAccount(ctx, account)
	// if err != nil {
	// 	return nil, err
	// }

	// return account, nil
}

// DeleteToken ...
func (c *accountsController) DeleteToken(ctx context.Context, accountID uuid.UUID) error {
	return nil
	// account, err := c.db.GetAccount(ctx, accountID)
	// if err != nil {
	// 	return err
	// }

	// operator, err := c.db.GetOperator(ctx, account.OperatorID)
	// if err != nil {
	// 	return err
	// }

	// osk, err := nkeys.FromSeed(operator.SigningKeys[0].Seed)
	// if err != nil {
	// 	return err
	// }

	// ac, err := jwt.DecodeAccountClaims(account.Token.Token)
	// if err != nil {
	// 	return err
	// }

	// ac.Expires = time.Now().Add(time.Minute).Unix()

	// for _, user := range account.Users {
	// 	if ac.Revocations[user.KeyID] == 0 {
	// 		ac.Revoke(user.KeyID)
	// 	}
	// }

	// token, err := ac.Encode(osk)
	// if err != nil {
	// 	return err
	// }
	// account.Token.Token = token

	// err = c.db.UpdateAccount(ctx, account)
	// if err != nil {
	// 	return err
	// }

	// return nil
}

// CreateSigningKeyGroup ...
func (c *accountsController) CreateSigningKeyGroup(ctx context.Context) (*models.Account, error) {
	return nil, nil
}

// ListSigningKeys ...
func (c *accountsController) ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error) {
	return c.db.ListSigningKeys(ctx, accountID, pagination)
}
