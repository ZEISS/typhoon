package ports

import (
	"context"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/zeiss/typhoon/internal/api/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// Teams ...
type Teams interface {
	CreateTeam(ctx context.Context, team *openapi.Team) (openapi.Team, error)
	ListTeams(ctx context.Context, params openapi.ListTeamParams) (models.PaginatedListTeams, error)
	GetTeamByID(ctx context.Context, id openapi_types.UUID) (openapi.Team, error)
}
