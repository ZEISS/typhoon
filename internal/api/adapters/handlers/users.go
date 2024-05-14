package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/dto"
)

// ListUsers ...
func (a *ApiHandlers) ListUsers(ctx context.Context, req openapi.ListUsersRequestObject) (openapi.ListUsersResponseObject, error) {
	query := dto.FromListUsersRequest(req)

	pagination, err := a.users.ListUsers(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToListUsersResponse(pagination), nil
}

// CreateUser ...
func (a *ApiHandlers) CreateUser(ctx context.Context, req openapi.CreateUserRequestObject) (openapi.CreateUserResponseObject, error) {
	cmd := dto.FromCreateUserRequest(req)

	user, err := a.users.CreateUser(ctx, cmd)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToCreateUserResponse(user), nil
}

// GetUser ...
func (a *ApiHandlers) GetUser(ctx context.Context, req openapi.GetUserRequestObject) (openapi.GetUserResponseObject, error) {
	query := dto.FromGetUserRequest(req)

	user, err := a.users.GetUser(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToGetUserResponse(user), nil
}

// GetUserCredentials ...
func (a *ApiHandlers) GetUserCredentials(ctx context.Context, req openapi.GetUserCredentialsRequestObject) (openapi.GetUserCredentialsResponseObject, error) {
	query := dto.FromGetUserCredentialsRequest(req)

	creds, err := a.users.GetCredentials(ctx, query)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return dto.ToGetUserCredentialsResponse(creds), nil
}
