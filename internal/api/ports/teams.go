package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Teams is the interface that wraps the methods to access data.
type Teams interface {
	// CreateTeam creates a new team.
	CreateTeam(ctx context.Context, team *models.Team) error
	// GetTeam returns the team with the given id.
	GetTeam(ctx context.Context, team *models.Team) error
	// DeleteTeam deletes the team with the given id.
	DeleteTeam(ctx context.Context, team *models.Team) error
	// ListTeams returns all teams.
	ListTeams(ctx context.Context, pagination *models.Pagination[models.Team]) error
}
