package handlers

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/dto"
	"github.com/zeiss/typhoon/internal/accounts/ports"
	openapi "github.com/zeiss/typhoon/pkg/apis/accounts"
)

var _ openapi.StrictServerInterface = (*AccountsHandler)(nil)

// AccountsHandler ...
type AccountsHandler struct {
	ac ports.AccountsController
}

// NewAccountsHandler ...
func NewAccountsHandler(ac ports.AccountsController) *AccountsHandler {
	return &AccountsHandler{ac: ac}
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
