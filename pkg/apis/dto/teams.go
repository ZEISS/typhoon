package dto

import (
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateTeamRequest ...
func FromCreateTeamRequest(req openapi.CreateTeamRequestObject) controllers.CreateTeamCommand {
	return controllers.CreateTeamCommand{
		Name:        req.Body.Name,
		Description: utils.PtrStr(req.Body.Description),
	}
}

// ToCreateTeamResponse ...
func ToCreateTeamResponse(team adapters.GothTeam) openapi.CreateTeam201JSONResponse {
	res := openapi.CreateTeam201JSONResponse{}
	res.Id = utils.PtrUUID(team.ID)
	res.Name = team.Name
	res.CreatedAt = utils.PtrTime(team.CreatedAt)
	res.UpdatedAt = utils.PtrTime(team.UpdatedAt)
	res.DeletedAt = utils.PtrTime(team.DeletedAt.Time)

	return res
}

// FromGetTeamRequest ...
func FromGetTeamRequest(req openapi.GetTeamRequestObject) controllers.GetTeamQuery {
	return controllers.GetTeamQuery{
		ID: req.TeamId,
	}
}

// ToGetTeamResponse ...
func ToGetTeamResponse(team adapters.GothTeam) openapi.GetTeam200JSONResponse {
	res := openapi.GetTeam200JSONResponse{}
	res.Id = utils.PtrUUID(team.ID)
	res.Name = team.Name
	res.Description = utils.StrPtr(team.Description)
	res.CreatedAt = utils.PtrTime(team.CreatedAt)
	res.UpdatedAt = utils.PtrTime(team.UpdatedAt)
	res.DeletedAt = utils.PtrTime(team.DeletedAt.Time)

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
func ToListTeamsResponse(pagination models.Pagination[adapters.GothTeam]) openapi.ListTeams200JSONResponse {
	res := openapi.ListTeams200JSONResponse{}
	res.Limit = utils.PtrInt(pagination.Limit)
	res.Offset = utils.PtrInt(pagination.Offset)

	teams := make([]openapi.Team, 0, len(pagination.Rows))
	for _, team := range pagination.Rows {
		teams = append(teams, openapi.Team{
			Id:        utils.PtrUUID(team.ID),
			Name:      team.Name,
			CreatedAt: utils.PtrTime(team.CreatedAt),
			UpdatedAt: utils.PtrTime(team.UpdatedAt),
			DeletedAt: utils.PtrTime(team.DeletedAt.Time),
		})
	}
	res.Results = &teams

	return res
}
