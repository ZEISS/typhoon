package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
)

// VersionHandler ...
type VersionHandler struct{}

// NewVersionHandler ...
func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}

// ListTeams ...
func (h *VersionHandler) Version(ctx context.Context, request openapi.VersionRequestObject) (openapi.VersionResponseObject, error) {
	return nil, nil
}
