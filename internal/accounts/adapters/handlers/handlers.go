package handlers

import (
	"context"
	"errors"

	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/typhoon/internal/accounts/controllers"
	"github.com/zeiss/typhoon/internal/accounts/dto"
	"gorm.io/gorm"

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

// GetHelp ...
func (h *AccountsHandler) GetHelp(ctx context.Context, req openapi.GetHelpRequestObject) (openapi.GetHelpResponseObject, error) {
	return openapi.GetHelp200Response{}, nil // this is a test endpoint
}

// GetToken ...
func (h *AccountsHandler) GetAccountToken(ctx context.Context, req openapi.GetAccountTokenRequestObject) (openapi.GetAccountTokenResponseObject, error) {
	query := dto.FromGetAccountTokenRequest(req)

	token, err := h.ac.GetToken(ctx, query)
	var queryError *seed.QueryError
	if errors.As(err, &queryError) && errors.Is(queryError.Err, gorm.ErrRecordNotFound) {
		return openapi.GetAccountToken404Response{}, nil
	}

	if err != nil {
		return nil, err
	}

	return dto.ToGetAccountTokenResponse(token), nil
}
