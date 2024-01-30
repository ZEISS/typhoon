package ports

import (
	"context"

	openapi_types "github.com/oapi-codegen/runtime/types"
	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/models"
)

// Teams ...
type Teams interface {
	CreateTeam(ctx context.Context, team *openapi.Team) (openapi.Team, error)
	ListTeams(ctx context.Context, params openapi.ListTeamParams) (models.PaginatedListTeams, error)
	GetTeamByID(ctx context.Context, id openapi_types.UUID) (openapi.Team, error)
}
