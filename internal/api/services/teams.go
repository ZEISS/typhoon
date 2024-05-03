package services

import (
	"context"

	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// CreateTeam ...
func (a *ApiHandlers) CreateTeam(ctx context.Context, req openapi.CreateTeamRequestObject) (openapi.CreateTeamResponseObject, error) {
	team, err := a.teams.CreateTeam(ctx, req.Body.Name, req.Body.Description)
	if err != nil {
		return openapi.CreateTeam400JSONResponse{}, nil
	}

	return openapi.CreateTeam201JSONResponse(openapi.CreateTeam201JSONResponse(openapi.Team{Name: team.Name, Description: team.Description})), nil
}

// GetTeam ...
func (a *ApiHandlers) GetTeam(ctx context.Context, req openapi.GetTeamRequestObject) (openapi.GetTeamResponseObject, error) {
	team, err := a.teams.GetTeam(ctx, req.TeamId)
	if err != nil {
		return openapi.GetTeam404JSONResponse{}, nil
	}

	return openapi.GetTeam200JSONResponse(openapi.Team{Id: utils.UUIDPtr(team.ID), Name: team.Name, Description: team.Description, CreatedAt: utils.PtrTime(team.CreatedAt), UpdatedAt: utils.PtrTime(team.UpdatedAt), DeletedAt: utils.PtrTime(team.DeletedAt.Time)}), nil
}
