package handlers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/controllers"
	"github.com/zeiss/typhoon/internal/accounts/dto"

	openapi "github.com/zeiss/typhoon/pkg/apis/accounts"
)

var _ openapi.StrictServerInterface = (*AccountsHandler)(nil)

// AccountsHandler ...
type AccountsHandler struct {
	ac controllers.AccountsController
}

// NewAccountsHandler ...
func NewAccountsHandler(ac controllers.AccountsController) *AccountsHandler {
	return &AccountsHandler{ac}
}

// GetToken ...
func (h *AccountsHandler) GetAccountToken(ctx context.Context, req openapi.GetAccountTokenRequestObject) (openapi.GetAccountTokenResponseObject, error) {
	query := dto.FromGetAccountTokenRequest(req)

	token, err := h.ac.GetToken(ctx, query)
	if err != nil {
		return openapi.GetAccountToken404Response{}, err
	}

	return dto.ToGetAccountTokenResponse(token), nil
}
