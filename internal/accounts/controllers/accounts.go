package controllers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/dto"
	"github.com/zeiss/typhoon/internal/accounts/models"
	"github.com/zeiss/typhoon/internal/accounts/ports"
	api "github.com/zeiss/typhoon/internal/api/models"
)

var _ ports.AccountsController = (*accountsController)(nil)

// AccountsController ...
type accountsController struct {
	db ports.AccountsRepository
}

// NewAccountsController ...
func NewAccountsController(db ports.AccountsRepository) *accountsController {
	return &accountsController{db: db}
}

// GetToken ...
func (c *accountsController) GetToken(ctx context.Context, query dto.GetAccountQuery) (models.AccountToken, error) {
	token := api.Token{ID: query.ID.String()}

	err := c.db.GetToken(ctx, &token)
	if err != nil {
		return models.AccountToken(token.Token), err
	}

	return models.AccountToken(token.Token), nil
}
