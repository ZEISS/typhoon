package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Systems is the interface that wraps the methods to access data.
type Systems interface {
	// CreateSystem creates a new system.
	CreateSystem(ctx context.Context, system *models.System) error
	// UpdateSystem updates a system.
	UpdateSystem(ctx context.Context, system *models.System) (*models.System, error)
	// GetSystem retrieves a system by its ID.
	GetSystem(ctx context.Context, id uuid.UUID) (*models.System, error)
	// ListSystems retrieves all systems.
	ListSystems(ctx context.Context, pagination models.Pagination[models.System]) (models.Pagination[models.System], error)
	// DeleteSystem deletes a system by its ID.
	DeleteSystem(ctx context.Context, id uuid.UUID) error
}
