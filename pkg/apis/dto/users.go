package dto

import (
	"bytes"

	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/models"
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

	res.Id = cast.Ptr(user.ID)
	res.Name = user.Name
	res.UpdatedAt = cast.Ptr(user.UpdatedAt)
	res.CreatedAt = cast.Ptr(user.CreatedAt)
	res.DeletedAt = cast.Ptr(user.DeletedAt.Time)

	return res
}

// FromListUsersRequest ...
func FromListUsersRequest(req openapi.ListUsersRequestObject) controllers.ListUsersQuery {
	return controllers.ListUsersQuery{
		AccountID: cast.Value(req.Body.AccountId),
	}
}

// ToListUsersResponse ...
func ToListUsersResponse(pagination models.Pagination[models.User]) openapi.ListUsers200JSONResponse {
	res := openapi.ListUsers200JSONResponse{}

	results := []openapi.User{}
	for _, user := range pagination.Rows {
		results = append(results, openapi.User{
			Id:        cast.Ptr(user.ID),
			Name:      user.Name,
			CreatedAt: cast.Ptr(user.CreatedAt),
			UpdatedAt: cast.Ptr(user.UpdatedAt),
			DeletedAt: cast.Ptr(user.DeletedAt.Time),
		})
	}
	res.Limit = cast.Ptr(pagination.Limit)
	res.Offset = cast.Ptr(pagination.Offset)
	res.Total = cast.Ptr(pagination.TotalRows)
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

// FromGetUserRequest ...
func FromGetUserRequest(req openapi.GetUserRequestObject) controllers.GetUserQuery {
	return controllers.GetUserQuery{
		UserID: req.UserId,
	}
}

// ToGetUserResponse ...
func ToGetUserResponse(user models.User) openapi.GetUser200JSONResponse {
	res := openapi.GetUser200JSONResponse{}

	res.Id = cast.Ptr(user.ID)
	res.Name = user.Name
	res.CreatedAt = cast.Ptr(user.CreatedAt)
	res.UpdatedAt = cast.Ptr(user.UpdatedAt)
	res.DeletedAt = cast.Ptr(user.DeletedAt.Time)

	return res
}
