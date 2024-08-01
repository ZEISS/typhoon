package controllers

import (
	"context"

	"github.com/zeiss/pkg/dbx"
	"github.com/zeiss/typhoon/internal/accounts/dto"
	"github.com/zeiss/typhoon/internal/accounts/ports"
	"github.com/zeiss/typhoon/internal/models"
)

var _ AccountsController = (*AccountsControllerImpl)(nil)

// AccountsController ...
type AccountsController interface {
	// GetToken is a method that returns a token.
	GetToken(ctx context.Context, query dto.GetAccountQuery) (models.Token, error)
}

// AccountsControllerImpl ...
type AccountsControllerImpl struct {
	store dbx.Database[ports.ReadTx, ports.ReadWriteTx]
}

// NewAccountsController ...
func NewAccountsController(store dbx.Database[ports.ReadTx, ports.ReadWriteTx]) *AccountsControllerImpl {
	return &AccountsControllerImpl{store}
}

// GetToken ...
func (c *AccountsControllerImpl) GetToken(ctx context.Context, query dto.GetAccountQuery) (models.Token, error) {
	token := models.Token{ID: query.ID}

	err := c.store.ReadTx(ctx, func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetToken(ctx, &token)
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
