package controllers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/models"
	"github.com/zeiss/typhoon/internal/accounts/ports"
)

// AccountsController ...
type AccountsController interface {
	GetToken(ctx context.Context, accountPublicKey models.AccountPublicKey) (models.AccountToken, error)
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
func (c *accountsController) GetToken(ctx context.Context, pubkey models.AccountPublicKey) (models.AccountToken, error) {
	return c.db.GetToken(ctx, pubkey)
}
