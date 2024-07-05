package accounts

import (
	"github.com/zeiss/typhoon/internal/web/controllers/accounts/tokens"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewGetAccountTokenController ...
func NewGetAccountTokenController(store ports.Datastore) *tokens.GetAccountTokenControllerImpl {
	return tokens.NewGetAccountTokenController(store)
}
