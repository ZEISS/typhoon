package ports

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
)

// Teams ...
type Teams interface {
	CreateTeam(ctx context.Context, team *openapi.Team) error
	ListTeams(ctx context.Context) ([]*openapi.Team, error)
}
