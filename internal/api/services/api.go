package services

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

var _ openapi.StrictServerInterface = (*ApiHandlers)(nil)

// ApiHandlers ...
type ApiHandlers struct {
	systems   *controllers.SystemsController
	teams     *controllers.TeamsController
	version   *controllers.VersionController
	operators *controllers.OperatorsController

	openapi.Unimplemented
}

// NewApiHandlers ...
func NewApiHandlers(systems *controllers.SystemsController, teams *controllers.TeamsController, version *controllers.VersionController, operators *controllers.OperatorsController) *ApiHandlers {
	return &ApiHandlers{systems: systems, teams: teams, version: version, operators: operators}
}

// CreateOperator ...
func (a *ApiHandlers) CreateOperator(ctx context.Context, req openapi.CreateOperatorRequestObject) (openapi.CreateOperatorResponseObject, error) {
	operator, err := a.operators.CreateOperator(ctx, req.Body.Name)
	if err != nil {
		return nil, err
	}

	return openapi.CreateOperator201JSONResponse(openapi.Operator{Id: &operator.ID, Name: operator.Name}), nil
}

// GetOperator ...
func (a *ApiHandlers) GetOperator(ctx context.Context, req openapi.GetOperatorRequestObject) (openapi.GetOperatorResponseObject, error) {
	operator, err := a.operators.GetOperator(ctx, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return openapi.GetOperator200JSONResponse(openapi.Operator{Id: &operator.ID, Name: operator.Name}), nil
}

// CreateAccount ...
func (a *ApiHandlers) CreateOperatorAccount(ctx context.Context, req openapi.CreateOperatorAccountRequestObject) (openapi.CreateOperatorAccountResponseObject, error) {
	account, err := a.operators.CreateOperatorAccount(ctx, req.Body.Name, req.OperatorId)
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

// DeleteOperator ...
func (a *ApiHandlers) DeleteOperator(ctx context.Context, req openapi.DeleteOperatorRequestObject) (openapi.DeleteOperatorResponseObject, error) {
	err := a.operators.DeleteOperator(ctx, req.OperatorId)
	if err != nil {
		return nil, err
	}

	return openapi.DeleteOperator204Response(openapi.DeleteOperator204Response{}), nil
}

// ListOperator ...
func (a *ApiHandlers) ListOperator(ctx context.Context, req openapi.ListOperatorRequestObject) (openapi.ListOperatorResponseObject, error) {
	pagination := models.Pagination[*models.Operator]{
		Limit:  *req.Params.Limit,
		Offset: *req.Params.Offset,
	}

	result, err := a.operators.ListOperator(ctx, pagination)
	if err != nil {
		return nil, err
	}

	operators := make([]openapi.Operator, 0, len(result.Rows))
	for _, operator := range result.Rows {
		operators = append(operators, openapi.Operator{Id: &operator.ID, Name: operator.Name})
	}

	return openapi.ListOperator200JSONResponse(openapi.ListOperator200JSONResponse{Results: &operators}), nil
}

// Version ...
func (a *ApiHandlers) Version(ctx context.Context, req openapi.VersionRequestObject) (openapi.VersionResponseObject, error) {
	version, err := a.version.Version()
	if err != nil {
		return nil, err
	}

	date, err := a.version.Date()
	if err != nil {
		return nil, err
	}

	return openapi.Version200JSONResponse(openapi.Version{Date: date, Version: version}), nil
}
