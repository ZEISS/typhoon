package dto

import (
	"bytes"

	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateUserRequest ...
func FromCreateUserRequest(req openapi.CreateUserRequestObject) controllers.CreateUserCommand {
	return controllers.CreateUserCommand{
		AccountID: req.Body.AccountId,
		Name:      req.Body.Name,
	}
}

// ToCreateUserResponse ...
func ToCreateUserResponse(user models.User) openapi.CreateUser201JSONResponse {
	res := openapi.CreateUser201JSONResponse{}

	res.Id = utils.PtrUUID(user.ID)
	res.Name = user.Name
	res.UpdatedAt = utils.PtrTime(user.UpdatedAt)
	res.CreatedAt = utils.PtrTime(user.CreatedAt)
	res.DeletedAt = utils.PtrTime(user.DeletedAt.Time)

	return res
}

// FromListUsersRequest ...
func FromListUsersRequest(req openapi.ListUsersRequestObject) controllers.ListUsersQuery {
	return controllers.ListUsersQuery{
		AccountID: utils.UUIDPtr(req.Body.AccountId),
	}
}

// ToListUsersResponse ...
func ToListUsersResponse(pagination models.Pagination[models.User]) openapi.ListUsers200JSONResponse {
	res := openapi.ListUsers200JSONResponse{}

	results := []openapi.User{}
	for _, user := range pagination.Rows {
		results = append(results, openapi.User{
			Id:        utils.PtrUUID(user.ID),
			Name:      user.Name,
			CreatedAt: utils.PtrTime(user.CreatedAt),
			UpdatedAt: utils.PtrTime(user.UpdatedAt),
			DeletedAt: utils.PtrTime(user.DeletedAt.Time),
		})
	}
	res.Limit = utils.PtrInt(pagination.Limit)
	res.Offset = utils.PtrInt(pagination.Offset)
	res.Total = utils.PtrInt(pagination.TotalRows)
	res.Results = &results

	return res
}

// FromGetUserCredentialsRequest ...
func FromGetUserCredentialsRequest(req openapi.GetUserCredentialsRequestObject) controllers.GetUserCredentialsQuery {
	return controllers.GetUserCredentialsQuery{
		UserID: req.UserId,
	}
}

// ToGetUserCredentialsResponse ...
func ToGetUserCredentialsResponse(creds []byte) openapi.GetUserCredentials200ApplicationoctetStreamResponse {
	res := openapi.GetUserCredentials200ApplicationoctetStreamResponse{}

	body := bytes.NewReader(creds)
	res.Body = body
	res.ContentLength = int64(len(creds))

	return res
}
