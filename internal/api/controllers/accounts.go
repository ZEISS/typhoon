package controllers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/utils"
)

// CreateAccountCommand ...
type CreateAccountCommand struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OperatorID  uuid.UUID `json:"operator_id"`
	TeamID      uuid.UUID `json:"team_id"`
}

// ListAccountsQuery ...
type ListAccountsQuery struct {
	OperatorID uuid.UUID `json:"system_id"`
	Limit      int       `json:"limit"`
	Offset     int       `json:"offset"`
}

// GetAccountQuery ...
type GetAccountQuery struct {
	AccountID uuid.UUID `json:"account_id"`
}

// GetAccountTokenQuery ...
type GetAccountTokenQuery struct {
	AccountID uuid.UUID `json:"account_id"`
}

// AccountsController is the interface that wraps the methods to access accounts.
type AccountsController interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, cmd CreateAccountCommand) (models.Account, error)
	// DeleteToken deletes a token.
	DeleteToken(ctx context.Context, accountID uuid.UUID) error
	// CreateSigningKeyGroup creates a new signing key group.
	CreateSigningKeyGroup(ctx context.Context) (*models.Account, error)
	// ListSigningKeys of an account.
	ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error)
	// ListAccounts ...
	ListAccounts(ctx context.Context, req ListAccountsQuery) (models.Pagination[models.Account], error)
	// GetAccount ...
	GetAccount(ctx context.Context, query GetAccountQuery) (models.Account, error)
	// GetAccountToken ...
	GetAccountToken(ctx context.Context, query GetAccountTokenQuery) (models.Token, error)
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
func (c *accountsController) CreateAccount(ctx context.Context, cmd CreateAccountCommand) (models.Account, error) {
	account := models.Account{
		Name:        cmd.Name,
		Description: utils.StrPtr(cmd.Description),
	}

	fmt.Println(account)

	operator := models.Operator{
		ID: cmd.OperatorID,
	}
	err := c.db.GetOperator(ctx, &operator)
	if err != nil {
		return account, err
	}

	pk, err := nkeys.CreateAccount()
	if err != nil {
		return account, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return account, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return account, err
	}
	account.Key = models.NKey{ID: id, Seed: seed}

	skg := models.SigningKeyGroup{Name: "Default", Description: "Default signing key group"}

	skgpk, err := nkeys.CreateAccount()
	if err != nil {
		return account, err
	}

	skgid, err := skgpk.PublicKey()
	if err != nil {
		return account, err
	}

	skgseed, err := skgpk.Seed()
	if err != nil {
		return account, err
	}
	skg.Key = models.NKey{ID: skgid, Seed: skgseed}
	account.SigningKeyGroups = append(account.SigningKeyGroups, skg)

	// @katallaxie: this is a bit weird, but I think it's a good idea to have a default signing key group
	osk, err := nkeys.FromSeed(operator.SigningKeyGroups[0].Key.Seed)
	if err != nil {
		return account, err
	}

	ac := jwt.NewAccountClaims(id)
	ac.Name = cmd.Name
	ac.Issuer = operator.Key.ID
	ac.SigningKeys.Add(skg.Key.ID)

	token, err := ac.Encode(osk)
	if err != nil {
		return account, err
	}
	account.Token = models.Token{ID: id, Token: token}

	err = c.db.CreateAccount(ctx, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

// ListAccounts ...
func (c *accountsController) ListAccounts(ctx context.Context, query ListAccountsQuery) (models.Pagination[models.Account], error) {
	accounts := models.Pagination[models.Account]{}

	accounts.Limit = query.Limit
	accounts.Offset = query.Offset

	err := c.db.ListAccounts(ctx, &accounts)
	if err != nil {
		return accounts, err
	}

	return accounts, nil
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

// GetAccount ...
func (c *accountsController) GetAccount(ctx context.Context, query GetAccountQuery) (models.Account, error) {
	account := models.Account{ID: query.AccountID}

	err := c.db.GetAccount(ctx, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

// GetAccountToken ...
func (c *accountsController) GetAccountToken(ctx context.Context, query GetAccountTokenQuery) (models.Token, error) {
	account := models.Account{ID: query.AccountID}

	err := c.db.GetAccount(ctx, &account)
	if err != nil {
		return models.Token{}, err
	}

	return account.Token, nil
}
