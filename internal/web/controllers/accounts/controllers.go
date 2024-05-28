package accounts

import (
	"github.com/zeiss/typhoon/internal/web/controllers/accounts/tokens"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewGetAccountTokenController ...
func NewGetAccountTokenController(db ports.Accounts) *tokens.GetAccountTokenControllerImpl {
	return tokens.NewGetAccountTokenController(db)
}
