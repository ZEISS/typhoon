package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateAccountRequest ...
func FromCreateAccountRequest(req openapi.CreateAccountRequestObject) controllers.CreateAccountRequest {
	return controllers.CreateAccountRequest{
		SystemID:    req.Body.SystemId,
		Name:        req.Body.Name,
		Description: utils.PtrStr(req.Body.Description),
	}
}

// ToCreateAccountResponse ...
func ToCreateAccountResponse(account models.Account) openapi.CreateAccount201JSONResponse {
	res := openapi.CreateAccount201JSONResponse{}

	res.Id = utils.PtrUUID(account.ID)
	res.Name = account.Name
	res.CreatedAt = utils.PtrTime(account.CreatedAt)
	res.UpdatedAt = utils.PtrTime(account.UpdatedAt)
	res.DeletedAt = utils.PtrTime(account.DeletedAt.Time)

	return res
}

// ToListAccountResponse ...
func ToListAccountResponse(output controllers.ListAccountsResponse) openapi.ListAccounts200JSONResponse {
	res := openapi.ListAccounts200JSONResponse{}

	res.Limit = utils.PtrInt(output.Limit)
	res.Offset = utils.PtrInt(output.Offset)
	res.Total = utils.PtrInt(output.Total)

	results := make([]openapi.Account, 0, len(output.Accounts))

	for _, account := range output.Accounts {
		row := openapi.Account{
			Id:        utils.PtrUUID(account.ID),
			Name:      account.Name,
			CreatedAt: utils.PtrTime(account.CreatedAt),
			UpdatedAt: utils.PtrTime(account.UpdatedAt),
			DeletedAt: utils.PtrTime(account.DeletedAt.Time),
		}
		results = append(results, row)
	}

	res.Results = &results

	return res
}

// FromListAccountRequest ...
func FromListAccountRequest(req openapi.ListAccountsRequestObject) controllers.ListAccountsRequest {
	return controllers.ListAccountsRequest{
		Limit:  utils.IntPtr(req.Params.Limit),
		Offset: utils.IntPtr(req.Params.Offset),
	}
}
