package handlers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/controllers"
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
func (h *AccountLookupRequestHandler) HandleLookupRequest(ctx context.Context, accountPublicKey string) (string, error) {
	query := controllers.GetTokenQuery{AccountPublicKey: accountPublicKey}

	result, err := h.ctrl.GetToken(ctx, query)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}
