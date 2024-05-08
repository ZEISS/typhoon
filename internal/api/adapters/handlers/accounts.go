package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/dto"
)

// ListAccounts ...
func (a *ApiHandlers) ListAccounts(ctx context.Context, request openapi.ListAccountsRequestObject) (openapi.ListAccountsResponseObject, error) {
	req := dto.FromListAccountRequest(request)

	output, err := a.accounts.ListAccounts(ctx, req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToListAccountResponse(output), nil
}
