package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// ToListAccountResponse ...
func ToListAccountResponse(output controllers.ListAccountsResponse) openapi.ListAccounts200JSONResponse {
	response := openapi.ListAccounts200JSONResponse{}

	response.Limit = utils.PtrInt(output.Limit)
	response.Offset = utils.PtrInt(output.Offset)
	response.Total = utils.PtrInt(output.Total)

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

	response.Results = &results

	return response
}

// FromListAccountRequest ...
func FromListAccountRequest(req openapi.ListAccountsRequestObject) controllers.ListAccountsRequest {
	return controllers.ListAccountsRequest{
		Limit:  utils.IntPtr(req.Params.Limit),
		Offset: utils.IntPtr(req.Params.Offset),
	}
}
