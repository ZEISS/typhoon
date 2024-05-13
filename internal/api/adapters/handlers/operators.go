package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/dto"
)

// CreateOperator handles the request to create a new operator.
func (a *ApiHandlers) CreateOperator(ctx context.Context, req openapi.CreateOperatorRequestObject) (openapi.CreateOperatorResponseObject, error) {
	cmd := dto.FromCreateOperatorRequest(req)

	operator, err := a.operators.CreateOperator(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToCreateOperatorResponse(operator), nil
}

// GetOperator handles the request to get an operator.
func (a *ApiHandlers) GetOperator(ctx context.Context, req openapi.GetOperatorRequestObject) (openapi.GetOperatorResponseObject, error) {
	query := dto.FromGetOperatorRequest(req)

	result, err := a.operators.GetOperator(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToGetOperatorResponse(result), nil
}

// CreateOperatorSigningKeyGroup handles the request to create a new signing key group.
func (a *ApiHandlers) CreateOperatorSigningKeyGroup(ctx context.Context, req openapi.CreateOperatorSigningKeyGroupRequestObject) (openapi.CreateOperatorSigningKeyGroupResponseObject, error) {
	cmd := dto.FromCreateSigningKeyGroupRequest(req)

	signingKeyGroup, err := a.operators.CreateOperatorSigningKeyGroup(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToCreateOperatorSigningKeyGroupResponse(signingKeyGroup), nil
}

// ListOperators handles the request to list operators.
func (a *ApiHandlers) ListOperators(ctx context.Context, req openapi.ListOperatorsRequestObject) (openapi.ListOperatorsResponseObject, error) {
	query := dto.FromListOperatorsRequest(req)

	result, err := a.operators.ListOperators(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToListOperatorsResponse(result), nil
}

// GetOperatorToken handles the request to get a token for an operator.
func (a *ApiHandlers) GetOperatorToken(ctx context.Context, req openapi.GetOperatorTokenRequestObject) (openapi.GetOperatorTokenResponseObject, error) {
	query := dto.FromGetOperatorTokenRequest(req)

	result, err := a.operators.GetOperatorToken(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToGetOperatorTokenResponse(result), nil
}

// DeleteOperator handles the request to delete an operator.
func (a *ApiHandlers) DeleteOperator(ctx context.Context, req openapi.DeleteOperatorRequestObject) (openapi.DeleteOperatorResponseObject, error) {
	cmd := dto.FromDeleteOperatorRequest(req)

	err := a.operators.DeleteOperator(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToDeleteOperatorResponse(), nil
}

// GetOperatorSystemAccount handles the request to get the system account for an operator.
func (a *ApiHandlers) GetOperatorSystemAccount(ctx context.Context, req openapi.GetOperatorSystemAccountRequestObject) (openapi.GetOperatorSystemAccountResponseObject, error) {
	query := dto.FromGetOperatorSystemAccountRequest(req)

	result, err := a.operators.GetOperatorSystemAccount(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToGetOperatorSystemAccountResponse(result), nil
}

// UpdateOperatorSystemAccount handles the request to update the system account for an operator.
func (a *ApiHandlers) UpdateOperatorSystemAccount(ctx context.Context, req openapi.UpdateOperatorSystemAccountRequestObject) (openapi.UpdateOperatorSystemAccountResponseObject, error) {
	cmd := dto.FromUpdateOperatorSystemAccountRequest(req)

	result, err := a.operators.UpdateOperatorSystemAccount(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToUpdateOperatorSystemAccountResponse(result), nil
}
