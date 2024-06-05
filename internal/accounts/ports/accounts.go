package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/dto"
	"github.com/zeiss/typhoon/internal/accounts/models"
	api "github.com/zeiss/typhoon/internal/api/models"
)

// AccountsRepository ...
type AccountsRepository interface {
	// GetToken returns the token for the given account.
	GetToken(ctx context.Context, account *api.Token) error
}

// AccountsController ...
type AccountsController interface {
	// GetToken returns the token for the given account.
	GetToken(ctx context.Context, req dto.GetAccountQuery) (models.AccountToken, error)
}
