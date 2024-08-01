package dto

import (
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateTeamRequest ...
func FromCreateTeamRequest(req openapi.CreateTeamRequestObject) controllers.CreateTeamCommand {
	return controllers.CreateTeamCommand{
		Name:        req.Body.Name,
		Description: cast.Value(req.Body.Description),
	}
}

// ToCreateTeamResponse ...
func ToCreateTeamResponse(team models.Team) openapi.CreateTeam201JSONResponse {
	res := openapi.CreateTeam201JSONResponse{}
	res.Id = cast.Ptr(team.ID)
	res.Name = team.Name
	res.CreatedAt = cast.Ptr(team.CreatedAt)
	res.UpdatedAt = cast.Ptr(team.UpdatedAt)
	res.DeletedAt = cast.Ptr(team.DeletedAt.Time)

	return res
}

// FromGetTeamRequest ...
func FromGetTeamRequest(req openapi.GetTeamRequestObject) controllers.GetTeamQuery {
	return controllers.GetTeamQuery{
		ID: req.TeamId,
	}
}

// ToGetTeamResponse ...
func ToGetTeamResponse(team models.Team) openapi.GetTeam200JSONResponse {
	res := openapi.GetTeam200JSONResponse{}
	res.Id = cast.Ptr(team.ID)
	res.Name = team.Name
	res.Description = cast.Ptr(team.Description)
	res.CreatedAt = cast.Ptr(team.CreatedAt)
	res.UpdatedAt = cast.Ptr(team.UpdatedAt)
	res.DeletedAt = cast.Ptr(team.DeletedAt.Time)

	return res
}

// FromDeleteTeamRequest ...
func FromDeleteTeamRequest(req openapi.DeleteTeamRequestObject) controllers.DeleteTeamCommand {
	return controllers.DeleteTeamCommand{
		ID: req.TeamId,
	}
}

// ToDeleteTeamResponse ...
func ToDeleteTeamResponse() openapi.DeleteTeam204Response {
	return openapi.DeleteTeam204Response{}
}

// FromListTeamsRequest ...
func FromListTeamsRequest(req openapi.ListTeamsRequestObject) controllers.ListTeamsQuery {
	return controllers.ListTeamsQuery{}
}

// ToListTeamsResponse ...
func ToListTeamsResponse(pagination models.Pagination[models.Team]) openapi.ListTeams200JSONResponse {
	res := openapi.ListTeams200JSONResponse{}
	res.Limit = cast.Ptr(pagination.Limit)
	res.Offset = cast.Ptr(pagination.Offset)

	teams := make([]openapi.Team, 0, len(pagination.Rows))
	for _, team := range pagination.Rows {
		teams = append(teams, openapi.Team{
			Id:        cast.Ptr(team.ID),
			Name:      team.Name,
			CreatedAt: cast.Ptr(team.CreatedAt),
			UpdatedAt: cast.Ptr(team.UpdatedAt),
			DeletedAt: cast.Ptr(team.DeletedAt.Time),
		})
	}
	res.Results = &teams

	return res
}
