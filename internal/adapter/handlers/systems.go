package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
)

// SystemsHandler ...
type SystemsHandler struct{}

// NewSystemsHandler ...
func NewSystemsHandler() *SystemsHandler {
	return &SystemsHandler{}
}

// ListTeams ...
func (h *SystemsHandler) ListSystems(ctx context.Context, request openapi.ListSystemsRequestObject) (openapi.ListSystemsResponseObject, error) {
	return nil, nil
}

// ShowSystem...
func (h *SystemsHandler) ShowSystem(ctx context.Context, request openapi.ShowSystemRequestObject) (openapi.ShowSystemResponseObject, error) {
	return nil, nil
}
