package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Accounts ...
type Accounts interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, account *models.Account) error
}
