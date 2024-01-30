package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/ports"
)

// TeamsHandler ...
type TeamsHandler struct {
	ctrl ports.Teams
}

// NewTeamsHandler ...
func NewTeamsHandler(ctrl ports.Teams) *TeamsHandler {
	return &TeamsHandler{ctrl}
}

// Create Team
func (h *TeamsHandler) CreateTeam(ctx context.Context, request openapi.CreateTeamRequestObject) (openapi.CreateTeamResponseObject, error) {
	team, err := h.ctrl.CreateTeam(ctx, request.Body)
	if err != nil {
		return nil, err
	}

	return openapi.CreateTeam201JSONResponse(team), nil
}

// ListTeams ...
func (h *TeamsHandler) ListTeam(ctx context.Context, request openapi.ListTeamRequestObject) (openapi.ListTeamResponseObject, error) {
	teams, err := h.ctrl.ListTeams(ctx, request.Params)
	if err != nil {
		return nil, err
	}

	return openapi.ListTeam200JSONResponse(teams), nil
}

// GetTeamTeamId ...
func (h *TeamsHandler) GetTeamTeamId(ctx context.Context, request openapi.GetTeamTeamIdRequestObject) (openapi.GetTeamTeamIdResponseObject, error) {
	team, err := h.ctrl.GetTeamByID(ctx, request.TeamId)
	if err != nil {
		return nil, err
	}

	return openapi.GetTeamTeamId200JSONResponse(team), err
}
