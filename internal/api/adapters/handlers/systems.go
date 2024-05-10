package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/dto"
)

// GetSystem ...
func (a *ApiHandlers) GetSystem(ctx context.Context, req openapi.GetSystemRequestObject) (openapi.GetSystemResponseObject, error) {
	query := dto.FromGetSystemRequest(req)

	system, err := a.systems.GetSystem(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToGetSystemResponse(system), nil
}

// DeleteSystem ...
func (a *ApiHandlers) DeleteSystem(ctx context.Context, req openapi.DeleteSystemRequestObject) (openapi.DeleteSystemResponseObject, error) {
	cmd := dto.FromDeleteSystemRequest(req)

	err := a.systems.DeleteSystem(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToDeleteSystemResponse(), nil
}

// ListSystems ...
func (a *ApiHandlers) ListSystems(ctx context.Context, req openapi.ListSystemsRequestObject) (openapi.ListSystemsResponseObject, error) {
	query := dto.FromListSystemsRequest(req)

	result, err := a.systems.ListSystems(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToListSystemsResponse(result), nil
}

// CreateSystem ...
func (a *ApiHandlers) CreateSystem(ctx context.Context, req openapi.CreateSystemRequestObject) (openapi.CreateSystemResponseObject, error) {
	cmd := dto.FromCreateSystemRequest(req)

	system, err := a.systems.CreateSystem(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToCreateSystemResponse(system), nil
}
