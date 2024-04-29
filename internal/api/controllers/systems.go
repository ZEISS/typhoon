package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// SystemsController ...
type SystemsController struct {
	db ports.Systems
}

// NewSystemsController ...
func NewSystemsController(db ports.Systems) *SystemsController {
	return &SystemsController{db}
}

// CreateSystem ...
func (c *SystemsController) CreateSystem(ctx context.Context, name, description string) (*models.System, error) {
	system := &models.System{
		Name:        name,
		Description: description,
	}

	if err := c.db.CreateSystem(ctx, system); err != nil {
		return nil, err
	}

	return system, nil
}

// DeleteSystem ...
func (c *SystemsController) DeleteSystem(ctx context.Context, id uuid.UUID) error {
	return c.db.DeleteSystem(ctx, id)
}

// GetSystem ...
func (c *SystemsController) GetSystem(ctx context.Context, id uuid.UUID) (*models.System, error) {
	return c.db.GetSystem(ctx, id)
}

// ListSystems ...
func (c *SystemsController) ListSystems(ctx context.Context, pagination models.Pagination[models.System]) (models.Pagination[models.System], error) {
	return c.db.ListSystems(ctx, pagination)
}

// UpdateSystemOperator ...
func (c *SystemsController) UpdateSystemOperator(ctx context.Context, systemId, operatorID uuid.UUID) (*models.System, error) {
	system, err := c.db.GetSystem(ctx, systemId)
	if err != nil {
		return nil, err
	}

	system.OperatorID = &operatorID

	return c.db.UpdateSystem(ctx, system)
}
