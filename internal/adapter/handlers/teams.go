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

// ListTeams ...
func (h *TeamsHandler) ListTeam(ctx context.Context, request openapi.ListTeamRequestObject) (openapi.ListTeamResponseObject, error) {
	teams, err := h.ctrl.ListTeams(ctx)
	if err != nil {
		return nil, err
	}

	tt := make([]openapi.Team, 0, 0)
	for _, t := range teams {
		tt = append(tt, *t)
	}

	return openapi.ListTeam200JSONResponse(tt), nil
}

func (h *TeamsHandler) CreateTeam(ctx context.Context, request openapi.CreateTeamRequestObject) (openapi.CreateTeamResponseObject, error) {
	return nil, nil
}
