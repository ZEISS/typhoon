package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// CreateSystemCommand ...
type CreateSystemCommand struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=1024"`
}

// DeleteSystemCommand ...
type DeleteSystemCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// GetSystemQuery ...
type GetSystemQuery struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// ListSystemsQuery ...
type ListSystemsQuery struct {
	Limit  int    `json:"limit" validate:"required"`
	Offset int    `json:"offset" validate:"required"`
	Search string `json:"search"`
	Sort   string `json:"sort"`
}

var _ SystemsController = (*SystemsControllerImpl)(nil)

// SystemsController is the controller for systems.
type SystemsController interface {
	// CreateSystem creates a new system.
	CreateSystem(ctx context.Context, cmd CreateSystemCommand) (models.System, error)
	// DeleteSystem deletes a system.
	DeleteSystem(ctx context.Context, cmd DeleteSystemCommand) error
	// GetSystem retrieves a system by its ID.
	GetSystem(ctx context.Context, query GetSystemQuery) (models.System, error)
	// ListSystems retrieves a list of systems.
	ListSystems(ctx context.Context, query ListSystemsQuery) (models.Pagination[models.System], error)
}

// SystemsControllerImpl ...
type SystemsControllerImpl struct {
	db ports.Systems
}

// NewSystemsController ...
func NewSystemsController(db ports.Systems) *SystemsControllerImpl {
	return &SystemsControllerImpl{db}
}

// CreateSystem is the implementation of the CreateSystem method.
func (s *SystemsControllerImpl) CreateSystem(ctx context.Context, cmd CreateSystemCommand) (models.System, error) {
	system := models.System{}
	system.Name = cmd.Name
	system.Description = cmd.Description

	if err := s.db.CreateSystem(ctx, &system); err != nil {
		return system, err
	}

	return system, nil
}

// DeleteSystem is deleting a system.
func (s *SystemsControllerImpl) DeleteSystem(ctx context.Context, cmd DeleteSystemCommand) error {
	return s.db.DeleteSystem(ctx, cmd.ID)
}

// GetSystem ...
func (s *SystemsControllerImpl) GetSystem(ctx context.Context, query GetSystemQuery) (models.System, error) {
	op := models.System{ID: query.ID}

	err := s.db.GetSystem(ctx, &op)
	if err != nil {
		return op, err
	}

	return op, nil
}

// ListSystems ...
func (s *SystemsControllerImpl) ListSystems(ctx context.Context, query ListSystemsQuery) (models.Pagination[models.System], error) {
	sys := models.Pagination[models.System]{}

	sys.Limit = query.Limit
	sys.Offset = query.Offset
	sys.Search = query.Search

	err := s.db.ListSystems(ctx, &sys)
	if err != nil {
		return sys, err
	}

	return sys, nil
}

// UpdateSystemOperator ...
func (s *SystemsControllerImpl) UpdateSystemOperator(ctx context.Context, systemId, operatorID uuid.UUID) (*models.System, error) {
	system := models.System{}
	system.ID = systemId

	err := s.db.GetSystem(ctx, &system)
	if err != nil {
		return nil, err
	}

	system.OperatorID = &operatorID

	return s.db.UpdateSystem(ctx, &system)
}
