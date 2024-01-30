package controllers

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/ports"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

var _ ports.Teams = (*Teams)(nil)

// Teams ...
type Teams struct {
	port ports.Teams
}

// NewTeamsController ...
func NewTeamsController(port ports.Teams) *Teams {
	return &Teams{port}
}

// CreateTeam ...
func (t *Teams) CreateTeam(ctx context.Context, team *openapi.Team) (openapi.Team, error) {
	return t.port.CreateTeam(ctx, team)
}

// ListTeams ...
func (t *Teams) ListTeams(ctx context.Context, params openapi.ListTeamParams) (models.PaginatedListTeams, error) {
	return t.port.ListTeams(ctx, params)
}

// GetTeam ...
func (t *Teams) GetTeamByID(ctx context.Context, id openapi_types.UUID) (openapi.Team, error) {
	return t.port.GetTeamByID(ctx, id)
}
