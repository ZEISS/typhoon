package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Accounts is a port that defines the methods for accounts
type Accounts interface {
	// ListAccounts ...
	ListAccounts(ctx context.Context, pagination *models.Pagination[models.Account]) error
	// CreateAccount ...
	CreateAccount(ctx context.Context, account *models.Account) error
	// GetAccount ...
	GetAccount(ctx context.Context, account *models.Account) error
	// UpdateAccount ...
	UpdateAccount(ctx context.Context, account *models.Account) error
	// DeleteAccount ...
	DeleteAccount(ctx context.Context, account *models.Account) error
}
