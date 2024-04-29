package services

import (
	"bytes"
	"context"

	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

var _ openapi.StrictServerInterface = (*ApiHandlers)(nil)

// ApiHandlers ...
type ApiHandlers struct {
	systems   *controllers.SystemsController
	teams     *controllers.TeamsController
	version   *controllers.VersionController
	operators *controllers.OperatorsController
	accounts  *controllers.AccountsController
	users     *controllers.UsersController

	openapi.Unimplemented
}

// NewApiHandlers ...
func NewApiHandlers(systems *controllers.SystemsController, teams *controllers.TeamsController, version *controllers.VersionController, operators *controllers.OperatorsController, accounts *controllers.AccountsController, users *controllers.UsersController) *ApiHandlers {
	return &ApiHandlers{systems: systems, teams: teams, version: version, operators: operators, accounts: accounts, users: users}
}

// CreateSystem ...
func (a *ApiHandlers) CreateSystem(ctx context.Context, req openapi.CreateSystemRequestObject) (openapi.CreateSystemResponseObject, error) {
	system, err := a.systems.CreateSystem(ctx, req.Body.Name, *req.Body.Description)
	if err != nil {
		return nil, err
	}

	return openapi.CreateSystem201JSONResponse(openapi.System{Id: &system.ID, Name: system.Name, Description: utils.StrPtr(system.Description)}), nil
}

// GetSystem ...
func (a *ApiHandlers) GetSystem(ctx context.Context, req openapi.GetSystemRequestObject) (openapi.GetSystemResponseObject, error) {
	system, err := a.systems.GetSystem(ctx, req.SystemId)
	if err != nil {
		return nil, err
	}

	return openapi.GetSystem200JSONResponse(openapi.System{Id: &system.ID, Name: system.Name, Description: utils.StrPtr(system.Description), Operator: &openapi.Operator{Id: system.OperatorID}}), nil
}

// GetSystemOperator ...
func (a *ApiHandlers) GetSystemOperator(ctx context.Context, req openapi.GetSystemOperatorRequestObject) (openapi.GetSystemOperatorResponseObject, error) {
	system, err := a.systems.GetSystem(ctx, req.SystemId)
	if err != nil {
		return nil, err
	}

	return openapi.GetSystemOperator200JSONResponse(openapi.Operator{Id: utils.PtrUUID(system.Operator.ID), Name: system.Operator.Name}), nil
}

// DeleteSystem ...
func (a *ApiHandlers) DeleteSystem(ctx context.Context, req openapi.DeleteSystemRequestObject) (openapi.DeleteSystemResponseObject, error) {
	err := a.systems.DeleteSystem(ctx, req.SystemId)
	if err != nil {
		return nil, err
	}

	return openapi.DeleteSystem204Response(openapi.DeleteSystem204Response{}), nil
}

// ListSystems ...
func (a *ApiHandlers) ListSystems(ctx context.Context, req openapi.ListSystemsRequestObject) (openapi.ListSystemsResponseObject, error) {
	pagination := models.Pagination[models.System]{}

	result, err := a.systems.ListSystems(ctx, pagination)
	if err != nil {
		return nil, err
	}

	systems := make([]openapi.System, 0, len(result.Rows))
	for _, system := range result.Rows {
		systems = append(systems, openapi.System{
			Id:          &system.ID,
			Name:        system.Name,
			Description: utils.StrPtr(system.Description),
			CreatedAt:   &system.CreatedAt,
			UpdatedAt:   &system.UpdatedAt,
			DeletedAt:   &system.DeletedAt.Time,
		})
	}

	return openapi.ListSystems200JSONResponse(openapi.ListSystems200JSONResponse{Results: &systems}), nil
}

// UpdateSystemOperator ...
func (a *ApiHandlers) UpdateSystemOperator(ctx context.Context, req openapi.UpdateSystemOperatorRequestObject) (openapi.UpdateSystemOperatorResponseObject, error) {
	system, err := a.systems.UpdateSystemOperator(ctx, req.SystemId, *&req.Body.OperatorId)
	if err != nil {
		return nil, err
	}

	return openapi.UpdateSystemOperator200JSONResponse(openapi.UpdateSystemOperator200JSONResponse{Id: utils.PtrUUID(system.ID)}), nil
}

// CreateOperator ...
func (a *ApiHandlers) CreateOperator(ctx context.Context, req openapi.CreateOperatorRequestObject) (openapi.CreateOperatorResponseObject, error) {
	operator, err := a.operators.CreateOperator(ctx, req.Body.Name)
	if err != nil {
		return nil, err
	}

	return openapi.CreateOperator201JSONResponse(openapi.Operator{Id: &operator.ID, Name: operator.Name}), nil
}

// CreateOperatorAccount ...
func (a *ApiHandlers) CreateOperatorAccount(ctx context.Context, req openapi.CreateOperatorAccountRequestObject) (openapi.CreateOperatorAccountResponseObject, error) {
	account, err := a.accounts.CreateAccount(ctx, req.Body.Name, req.OperatorId)
	if err != nil {
		return nil, err
	}

	resp := openapi.CreateOperatorAccount201JSONResponse(
		openapi.Account{
			Id:        &account.ID,
			Name:      account.Name,
			CreatedAt: &account.CreatedAt,
			UpdatedAt: &account.UpdatedAt,
			DeletedAt: &account.DeletedAt.Time,
		},
	)

	return openapi.CreateOperatorAccount201JSONResponse(resp), nil
}

// CreateOperatorAccountUser ...
func (a *ApiHandlers) CreateOperatorAccountUser(ctx context.Context, req openapi.CreateOperatorAccountUserRequestObject) (openapi.CreateOperatorAccountUserResponseObject, error) {
	user, err := a.users.CreateUser(ctx, req.Body.Name, req.AccountId)
	if err != nil {
		return nil, err
	}

	return openapi.CreateOperatorAccountUser201JSONResponse(openapi.User{Id: &user.ID, Name: user.Name}), nil
}

// GetOperatorAccountUserCredentials ...
func (a *ApiHandlers) GetOperatorAccountUserCredentials(ctx context.Context, req openapi.GetOperatorAccountUserCredentialsRequestObject) (openapi.GetOperatorAccountUserCredentialsResponseObject, error) {
	credentials, err := a.users.GetCredentials(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(credentials)

	return openapi.GetOperatorAccountUserCredentials200ApplicationoctetStreamResponse(openapi.GetOperatorAccountUserCredentials200ApplicationoctetStreamResponse{Body: body, ContentLength: int64(body.Len())}), nil
}

// UpdateOperatorAccount ...
func (a *ApiHandlers) UpdateOperatorAccount(ctx context.Context, req openapi.UpdateOperatorAccountRequestObject) (openapi.UpdateOperatorAccountResponseObject, error) {
	account, err := a.accounts.UpdateAccount(ctx, controllers.UpdateOperatorAccountRequestObject(req))
	if err != nil {
		return nil, err
	}

	return openapi.UpdateOperatorAccount200JSONResponse(openapi.Account{Name: account.Name}), nil
}

// // GetOperator ...
// func (a *ApiHandlers) GetOperator(ctx context.Context, req openapi.GetOperatorRequestObject) (openapi.GetOperatorResponseObject, error) {
// 	operator, err := a.operators.GetOperator(ctx, req.OperatorId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.GetOperator200JSONResponse(openapi.Operator{Id: &operator.ID, Name: operator.Name}), nil
// }

// // CreateAccount ...
// func (a *ApiHandlers) CreateOperatorAccount(ctx context.Context, req openapi.CreateOperatorAccountRequestObject) (openapi.CreateOperatorAccountResponseObject, error) {
// 	account, err := a.operators.CreateOperatorAccount(ctx, req.Body.Name, req.OperatorId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := openapi.CreateOperatorAccount201JSONResponse(
// 		openapi.Account{
// 			Id:        &account.ID,
// 			Name:      account.Name,
// 			CreatedAt: &account.CreatedAt,
// 			UpdatedAt: &account.UpdatedAt,
// 			DeletedAt: &account.DeletedAt.Time,
// 		},
// 	)

// 	return openapi.CreateOperatorAccount201JSONResponse(resp), nil
// }

// // DeleteOperator ...
// func (a *ApiHandlers) DeleteOperator(ctx context.Context, req openapi.DeleteOperatorRequestObject) (openapi.DeleteOperatorResponseObject, error) {
// 	err := a.operators.DeleteOperator(ctx, req.OperatorId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.DeleteOperator204Response(openapi.DeleteOperator204Response{}), nil
// }

// // ListOperatorAccounts ...
// func (a *ApiHandlers) ListOperatorAccounts(ctx context.Context, req openapi.ListOperatorAccountsRequestObject) (openapi.ListOperatorAccountsResponseObject, error) {
// 	pagination := models.Pagination[*models.Account]{
// 		Limit:  *req.Params.Limit,
// 		Offset: *req.Params.Offset,
// 	}

// 	result, err := a.operators.ListOperatorAccount(ctx, req.OperatorId, pagination)
// 	if err != nil {
// 		return nil, err
// 	}

// 	accounts := make([]openapi.Account, 0, len(result.Rows))
// 	for _, account := range result.Rows {
// 		accounts = append(accounts, openapi.Account{
// 			Id:        &account.ID,
// 			Name:      account.Name,
// 			CreatedAt: &account.CreatedAt,
// 			UpdatedAt: &account.UpdatedAt,
// 			DeletedAt: &account.DeletedAt.Time,
// 		})
// 	}

// 	return openapi.ListOperatorAccounts200JSONResponse(openapi.ListOperatorAccounts200JSONResponse{Results: &accounts}), nil
// }

// // ListOperatorAccountUsers ...
// func (a *ApiHandlers) ListOperatorAccountUsers(ctx context.Context, req openapi.ListOperatorAccountUsersRequestObject) (openapi.ListOperatorAccountUsersResponseObject, error) {
// 	pagination := models.Pagination[*models.User]{
// 		Limit:  *req.Params.Limit,
// 		Offset: *req.Params.Offset,
// 	}

// 	result, err := a.operators.ListOperatorAccountUsers(ctx, req.AccountId, pagination)
// 	if err != nil {
// 		return nil, err
// 	}

// 	users := make([]openapi.User, 0, len(result.Rows))
// 	for _, user := range result.Rows {
// 		users = append(users, openapi.User{Id: &user.ID, Name: user.Name})
// 	}

// 	return openapi.ListOperatorAccountUsers200JSONResponse(openapi.ListOperatorAccountUsers200JSONResponse{Results: &users}), nil
// }

// // ListOperator ...
// func (a *ApiHandlers) ListOperator(ctx context.Context, req openapi.ListOperatorRequestObject) (openapi.ListOperatorResponseObject, error) {
// 	pagination := models.Pagination[*models.Operator]{
// 		Limit:  *req.Params.Limit,
// 		Offset: *req.Params.Offset,
// 	}

// 	result, err := a.operators.ListOperator(ctx, pagination)
// 	if err != nil {
// 		return nil, err
// 	}

// 	operators := make([]openapi.Operator, 0, len(result.Rows))
// 	for _, operator := range result.Rows {
// 		operators = append(operators, openapi.Operator{Id: &operator.ID, Name: operator.Name})
// 	}

// 	return openapi.ListOperator200JSONResponse(openapi.ListOperator200JSONResponse{Results: &operators}), nil
// }

// // CreateOperatorToken ...
// func (a *ApiHandlers) CreateOperatorToken(ctx context.Context, req openapi.CreateOperatorTokenRequestObject) (openapi.CreateOperatorTokenResponseObject, error) {
// 	token, err := a.operators.CreateOperatorToken(ctx, req.OperatorId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorToken201JSONResponse(openapi.JWTToken(token.Token)), nil
// }

// // CreateOperatorSigningKey ...
// func (a *ApiHandlers) CreateOperatorSigningKey(ctx context.Context, req openapi.CreateOperatorSigningKeyRequestObject) (openapi.CreateOperatorSigningKeyResponseObject, error) {
// 	key, err := a.operators.CreateOperatorSigningKey(ctx, req.OperatorId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorSigningKey201JSONResponse(openapi.KeyPair{PublicKey: key.ID}), nil
// }

// // CreateOperatorAccountSigningKey ...
// func (a *ApiHandlers) CreateOperatorAccountSigningKey(ctx context.Context, req openapi.CreateOperatorAccountSigningKeyRequestObject) (openapi.CreateOperatorAccountSigningKeyResponseObject, error) {
// 	key, err := a.operators.CreateOperatorAccountSigningKey(ctx, req.AccountId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorAccountSigningKey201JSONResponse(openapi.KeyPair{PublicKey: key.ID}), nil
// }

// // CreateOperatorAccountToken ...
// func (a *ApiHandlers) CreateOperatorAccountToken(ctx context.Context, req openapi.CreateOperatorAccountTokenRequestObject) (openapi.CreateOperatorAccountTokenResponseObject, error) {
// 	token, err := a.operators.CreateOperatorAccountToken(ctx, req.AccountId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorAccountToken201JSONResponse(openapi.JWTToken(token.Token)), nil
// }

// // CreateOperatorAccountUser ...
// func (a *ApiHandlers) CreateOperatorAccountUser(ctx context.Context, req openapi.CreateOperatorAccountUserRequestObject) (openapi.CreateOperatorAccountUserResponseObject, error) {
// 	user, err := a.operators.CreateOperatorAccountUser(ctx, req.AccountId, req.Body.Name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorAccountUser201JSONResponse(openapi.User{Id: &user.ID, Name: user.Name}), nil
// }

// // CreateOperatorAccountUserToken ...
// func (a *ApiHandlers) CreateOperatorAccountUserToken(ctx context.Context, req openapi.CreateOperatorAccountUserTokenRequestObject) (openapi.CreateOperatorAccountUserTokenResponseObject, error) {
// 	token, err := a.operators.CreateOperatorAccountUserToken(ctx, req.UserId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.CreateOperatorAccountUserToken201JSONResponse(openapi.JWTToken(token.Token)), nil
// }

// // Version ...
// func (a *ApiHandlers) Version(ctx context.Context, req openapi.VersionRequestObject) (openapi.VersionResponseObject, error) {
// 	version, err := a.version.Version()
// 	if err != nil {
// 		return nil, err
// 	}

// 	date, err := a.version.Date()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return openapi.Version200JSONResponse(openapi.Version{Date: date, Version: version}), nil
// }
