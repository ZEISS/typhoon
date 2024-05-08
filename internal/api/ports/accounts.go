package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Accounts ...
type Accounts interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, account *models.Account) error
	// UpdateAccount updates an existing account.
	UpdateAccount(ctx context.Context, account *models.Account) error
	// GetAccount returns the account with the given ID.
	GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error)
	// ListAccounts returns a list of accounts.
	ListAccounts(ctx context.Context, pagination models.Pagination[models.Account]) (models.Pagination[models.Account], error)
	// ListSigningKeys returns a list of signing keys for the account with the given ID.
	ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error)
}
