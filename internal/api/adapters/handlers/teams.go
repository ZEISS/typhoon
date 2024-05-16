package handlers

import (
	"context"

	openapi "github.com/zeiss/typhoon/pkg/apis"
	"github.com/zeiss/typhoon/pkg/apis/dto"
)

// CreateTeam ...
func (a *ApiHandlers) CreateTeam(ctx context.Context, req openapi.CreateTeamRequestObject) (openapi.CreateTeamResponseObject, error) {
	cmd := dto.FromCreateTeamRequest(req)

	team, err := a.teams.CreateTeam(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToCreateTeamResponse(team), nil
}

// GetTeam ...
func (a *ApiHandlers) GetTeam(ctx context.Context, req openapi.GetTeamRequestObject) (openapi.GetTeamResponseObject, error) {
	query := dto.FromGetTeamRequest(req)

	team, err := a.teams.GetTeam(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToGetTeamResponse(team), nil
}

// DeleteTeam ...
func (a *ApiHandlers) DeleteTeam(ctx context.Context, req openapi.DeleteTeamRequestObject) (openapi.DeleteTeamResponseObject, error) {
	cmd := dto.FromDeleteTeamRequest(req)

	err := a.teams.DeleteTeam(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return dto.ToDeleteTeamResponse(), nil
}

// ListTeams ...
func (a *ApiHandlers) ListTeams(ctx context.Context, req openapi.ListTeamsRequestObject) (openapi.ListTeamsResponseObject, error) {
	query := dto.FromListTeamsRequest(req)

	teams, err := a.teams.ListTeams(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.ToListTeamsResponse(teams), nil
}
