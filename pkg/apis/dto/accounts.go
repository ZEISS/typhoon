package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateAccountRequest ...
func FromCreateAccountRequest(req openapi.CreateAccountRequestObject) controllers.CreateAccountCommand {
	return controllers.CreateAccountCommand{
		OperatorID:  req.Body.OperatorId,
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
func ToListAccountResponse(accounts models.Pagination[models.Account]) openapi.ListAccounts200JSONResponse {
	res := openapi.ListAccounts200JSONResponse{}

	res.Limit = utils.PtrInt(accounts.Limit)
	res.Offset = utils.PtrInt(accounts.Offset)
	res.Total = utils.PtrInt(accounts.TotalRows)

	results := make([]openapi.Account, 0, len(accounts.Rows))

	for _, account := range accounts.Rows {
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
func FromListAccountRequest(req openapi.ListAccountsRequestObject) controllers.ListAccountsQuery {
	return controllers.ListAccountsQuery{
		OperatorID: utils.UUIDPtr(req.Body.OperatorId),
		Limit:      utils.IntPtr(req.Params.Limit),
		Offset:     utils.IntPtr(req.Params.Offset),
	}
}

// FromGetAccountRequest ...
func FromGetAccountRequest(req openapi.GetAccountRequestObject) controllers.GetAccountQuery {
	return controllers.GetAccountQuery{
		AccountID: req.AccountId,
	}
}

// ToGetAccountResponse ...
func ToGetAccountResponse(account models.Account) openapi.GetAccount200JSONResponse {
	res := openapi.GetAccount200JSONResponse{}

	res.Id = utils.PtrUUID(account.ID)
	res.Name = account.Name
	res.CreatedAt = utils.PtrTime(account.CreatedAt)
	res.UpdatedAt = utils.PtrTime(account.UpdatedAt)
	res.DeletedAt = utils.PtrTime(account.DeletedAt.Time)

	return res
}

// FromGetAccountTokenRequest ...
func FromGetAccountTokenRequest(req openapi.GetAccountTokenRequestObject) controllers.GetAccountTokenQuery {
	return controllers.GetAccountTokenQuery{
		AccountID: req.AccountId,
	}
}

// ToGetAccountTokenResponse ...
func ToGetAccountTokenResponse(token models.Token) openapi.GetAccountToken200JSONResponse {
	res := openapi.GetAccountToken200JSONResponse{}

	res.Token = utils.StrPtr(token.Token)

	return res
}
