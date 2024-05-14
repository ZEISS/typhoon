package controllers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/ports"
	"github.com/zeiss/typhoon/internal/api/models"
)

// GetTokenQuery ...
type GetTokenQuery struct {
	AccountPublicKey string `json:"account_public_key" validate:"required"`
}

// AccountsController ...
type AccountsController interface {
	GetToken(ctx context.Context, query GetTokenQuery) (models.Token, error)
}

// AccountsController ...
type accountsController struct {
	db ports.Accounts
}

// NewAccountsController ...
func NewAccountsController(db ports.Accounts) *accountsController {
	return &accountsController{db: db}
}

// GetToken ...
func (c *accountsController) GetToken(ctx context.Context, query GetTokenQuery) (models.Token, error) {
	token := models.Token{ID: query.AccountPublicKey}

	err := c.db.GetToken(ctx, &token)
	if err != nil {
		return token, err
	}

	return token, nil
}
