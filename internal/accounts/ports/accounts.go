package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/models"
)

// Accounts ...
type Accounts interface {
	// GetToken returns the token for the given account.
	GetToken(ctx context.Context, account models.AccountPublicKey) (models.AccountToken, error)
}
