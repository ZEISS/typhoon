package services

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/controllers"
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
	operator, err := a.operators.CreateOperator(req.Body.Name)
	if err != nil {
		return nil, err
	}

	return openapi.CreateOperator201JSONResponse(openapi.Operator{Id: &operator.ID, Name: operator.Name}), nil
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
