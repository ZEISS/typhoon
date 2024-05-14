package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Accounts ...
type Accounts interface {
	// GetToken returns the token for the given account.
	GetToken(ctx context.Context, account *models.Token) error
}
