package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Teams is the interface that wraps the methods to access data.
type Teams interface {
	// CreateTeam creates a new team.
	CreateTeam(ctx context.Context, team models.Team) (models.Team, error)
	// GetTeam returns the team with the given id.
	GetTeam(ctx context.Context, id uuid.UUID) (models.Team, error)
	// DeleteTeam deletes the team with the given id.
	DeleteTeam(ctx context.Context, id uuid.UUID) error
}
