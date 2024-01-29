package ports

import (
	"context"

	"github.com/zeiss/typhoon/api"
)

// Teams ...
type Teams interface {
	CreateTeam(ctx context.Context, team *api.Team) error
}
