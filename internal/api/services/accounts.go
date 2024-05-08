package services

import (
	"context"

	"github.com/gofiber/fiber/v2"
	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/transformers"
)

// ListAccounts ...
func (a *ApiHandlers) ListAccounts(ctx context.Context, request openapi.ListAccountsRequestObject) (openapi.ListAccountsResponseObject, error) {
	input := transformers.FromListAccountRequest(request)

	output, err := a.accounts.ListAccounts(ctx, input)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return transformers.ToListAccountResponse(output), nil
}
