package adapters

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/controllers"
	"github.com/zeiss/typhoon/internal/accounts/models"
)

// AccountLookupRequestHandler ...
type AccountLookupRequestHandler struct {
	ctrl controllers.AccountsController
}

// NewAccountLookupRequestHandler ...
func NewAccountLookupRequestHandler(ctrl controllers.AccountsController) *AccountLookupRequestHandler {
	return &AccountLookupRequestHandler{ctrl: ctrl}
}

// HandleLookupRequest ...
func (h *AccountLookupRequestHandler) HandleLookupRequest(ctx context.Context, accountPublicKey models.AccountPublicKey) (models.AccountToken, error) {
	return h.ctrl.GetToken(ctx, accountPublicKey)
}
