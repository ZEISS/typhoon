package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

var _ SystemsController = (*systemsController)(nil)

// SystemsController is the controller for systems.
type SystemsController interface {
	// CreateSystem creates a new system.
	CreateSystem(ctx context.Context, system *models.System) (*models.System, error)
	// DeleteSystem deletes a system.
	DeleteSystem(ctx context.Context, id uuid.UUID) error
	// GetSystem retrieves a system by its ID.
	GetSystem(ctx context.Context, id uuid.UUID) (*models.System, error)
	// ListSystems retrieves a list of systems.
	ListSystems(ctx context.Context, pagination models.Pagination[models.System]) (models.Pagination[models.System], error)
}

type systemsController struct {
	db ports.Systems
}

// NewSystemsController ...
func NewSystemsController(db ports.Systems) *systemsController {
	return &systemsController{db}
}

// CreateSystem ...
func (s *systemsController) CreateSystem(ctx context.Context, system *models.System) (*models.System, error) {
	if err := s.db.CreateSystem(ctx, system); err != nil {
		return nil, err
	}

	return system, nil
}

// DeleteSystem ...
func (s *systemsController) DeleteSystem(ctx context.Context, id uuid.UUID) error {
	return s.db.DeleteSystem(ctx, id)
}

// GetSystem ...
func (s *systemsController) GetSystem(ctx context.Context, id uuid.UUID) (*models.System, error) {
	return s.db.GetSystem(ctx, id)
}

// ListSystems ...
func (s *systemsController) ListSystems(ctx context.Context, pagination models.Pagination[models.System]) (models.Pagination[models.System], error) {
	return s.db.ListSystems(ctx, pagination)
}

// UpdateSystemOperator ...
func (s *systemsController) UpdateSystemOperator(ctx context.Context, systemId, operatorID uuid.UUID) (*models.System, error) {
	system, err := s.db.GetSystem(ctx, systemId)
	if err != nil {
		return nil, err
	}

	system.OperatorID = &operatorID

	return s.db.UpdateSystem(ctx, system)
}
