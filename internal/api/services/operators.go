package services

import (
	"context"

	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// ListOperators ...
func (a *ApiHandlers) ListOperators(ctx context.Context, req openapi.ListOperatorsRequestObject) (openapi.ListOperatorsResponseObject, error) {
	return nil, nil
}

// GetOperator ...
func (a *ApiHandlers) GetOperator(ctx context.Context, req openapi.GetOperatorRequestObject) (openapi.GetOperatorResponseObject, error) {
	result, err := a.operators.GetOperator(ctx, req.OperatorId)
	if err != nil {
		return nil, err
	}

	operator := openapi.Operator{
		Id:          utils.PtrUUID(result.ID),
		Name:        result.Name,
		Description: utils.StrPtr(result.Description),
		UpdatedAt:   utils.PtrTime(result.UpdatedAt),
		CreatedAt:   utils.PtrTime(result.CreatedAt),
		DeletedAt:   utils.PtrTime(result.DeletedAt.Time),
	}

	return openapi.GetOperatorResponseObject(openapi.GetOperator200JSONResponse(operator)), nil
}

// GetOperatorToken ...
func (a *ApiHandlers) GetOperatorToken(ctx context.Context, req openapi.GetOperatorTokenRequestObject) (openapi.GetOperatorTokenResponseObject, error) {
	result, err := a.operators.GetOperatorToken(ctx, req.OperatorId)
	if err != nil {
		return nil, err
	}

	token := openapi.JWTToken{
		Token: utils.StrPtr(result.Token),
	}

	return openapi.GetOperatorTokenResponseObject(openapi.GetOperatorToken200JSONResponse(token)), nil
}
