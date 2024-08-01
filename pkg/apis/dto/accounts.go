package dto

import (
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateAccountRequest ...
func FromCreateAccountRequest(req openapi.CreateAccountRequestObject) controllers.CreateAccountCommand {
	return controllers.CreateAccountCommand{
		OperatorID:  req.Body.OperatorId,
		Name:        req.Body.Name,
		Description: cast.Value(req.Body.Description),
		TeamID:      cast.Value(req.Body.TeamId),
	}
}

// ToCreateAccountResponse ...
func ToCreateAccountResponse(account models.Account) openapi.CreateAccount201JSONResponse {
	res := openapi.CreateAccount201JSONResponse{}

	res.Id = cast.Ptr(account.ID)
	res.Name = account.Name
	res.CreatedAt = cast.Ptr(account.CreatedAt)
	res.UpdatedAt = cast.Ptr(account.UpdatedAt)
	res.DeletedAt = cast.Ptr(account.DeletedAt.Time)

	return res
}

// ToListAccountResponse ...
func ToListAccountResponse(accounts models.Pagination[models.Account]) openapi.ListAccounts200JSONResponse {
	res := openapi.ListAccounts200JSONResponse{}

	res.Limit = cast.Ptr(accounts.Limit)
	res.Offset = cast.Ptr(accounts.Offset)
	res.Total = cast.Ptr(accounts.TotalRows)

	results := make([]openapi.Account, 0, len(accounts.Rows))

	for _, account := range accounts.Rows {
		row := openapi.Account{
			Id:        cast.Ptr(account.ID),
			Name:      account.Name,
			CreatedAt: cast.Ptr(account.CreatedAt),
			UpdatedAt: cast.Ptr(account.UpdatedAt),
			DeletedAt: cast.Ptr(account.DeletedAt.Time),
		}
		results = append(results, row)
	}

	res.Results = &results

	return res
}

// FromListAccountRequest ...
func FromListAccountRequest(req openapi.ListAccountsRequestObject) controllers.ListAccountsQuery {
	return controllers.ListAccountsQuery{
		OperatorID: cast.Value(req.Body.OperatorId),
		Limit:      cast.Value(req.Params.Limit),
		Offset:     cast.Value(req.Params.Offset),
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

	res.Id = cast.Ptr(account.ID)
	res.Name = account.Name
	res.CreatedAt = cast.Ptr(account.CreatedAt)
	res.UpdatedAt = cast.Ptr(account.UpdatedAt)
	res.DeletedAt = cast.Ptr(account.DeletedAt.Time)

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

	res.Token = cast.Ptr(token.Token)

	return res
}
