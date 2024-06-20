package ports

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Teams is the interface that wraps the methods to access data.
type Teams interface {
	// CreateTeam creates a new team.
	CreateTeam(ctx context.Context, team *adapters.GothTeam) error
	// GetTeam returns the team with the given id.
	GetTeam(ctx context.Context, team *adapters.GothTeam) error
	// DeleteTeam deletes the team with the given id.
	DeleteTeam(ctx context.Context, team *adapters.GothTeam) error
	// ListTeams returns all teams.
	ListTeams(ctx context.Context, pagination *models.Pagination[adapters.GothTeam]) error
}
