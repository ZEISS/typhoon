package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// CreateAccountRequest ...
type CreateAccountRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SystemID    uuid.UUID `json:"system_id"`
}

// CreateAccountResponse ...
type CreateAccountResponse struct {
	Account models.Account `json:"account"`
}

// ListAccountsRequest ...
type ListAccountsRequest struct {
	SystemID uuid.UUID `json:"system_id"`
	Limit    int       `json:"limit"`
	Offset   int       `json:"offset"`
}

// ListAccountsResponse ...
type ListAccountsResponse struct {
	Accounts []models.Account `json:"accounts"`
	Total    int              `json:"total"`
	Offset   int              `json:"offset"`
	Limit    int              `json:"limit"`
}

// AccountsController is the interface that wraps the methods to access accounts.
type AccountsController interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error)
	// DeleteToken deletes a token.
	DeleteToken(ctx context.Context, accountID uuid.UUID) error
	// CreateSigningKeyGroup creates a new signing key group.
	CreateSigningKeyGroup(ctx context.Context) (*models.Account, error)
	// ListSigningKeys of an account.
	ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error)
	// ListAccounts ...
	ListAccounts(ctx context.Context, req ListAccountsRequest) (ListAccountsResponse, error)
}

var _ AccountsController = (*accountsController)(nil)

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

// ListAccounts ...
func (c *accountsController) ListAccounts(ctx context.Context, input ListAccountsRequest) (ListAccountsResponse, error) {
	results := models.Pagination[models.Account]{}

	results, err := c.db.ListAccounts(ctx, results)
	if err != nil {
		return ListAccountsResponse{}, err
	}

	return ListAccountsResponse{Total: results.TotalRows, Limit: results.Limit, Offset: results.Limit, Accounts: results.Rows}, nil
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
