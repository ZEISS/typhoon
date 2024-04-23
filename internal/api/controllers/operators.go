package controllers

import (
	"context"

	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// OperatorsController ...
type OperatorsController struct {
	db ports.Operators
}

// NewOperatorsController ...
func NewOperatorsController(db ports.Operators) *OperatorsController {
	return &OperatorsController{db}
}

// CreateOperator ...
func (c *OperatorsController) CreateOperator(ctx context.Context, name string) (*models.Operator, error) {
	key, err := nkeys.CreateOperator()
	if err != nil {
		return nil, err
	}

	id, err := key.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := key.Seed()
	if err != nil {
		return nil, err
	}

	op := &models.Operator{
		ID: id,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
	}

	err = c.db.CreateOperator(ctx, op)
	if err != nil {
		return nil, err
	}

	return op, nil
}

// ListOperator ...
func (c *OperatorsController) ListOperator(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error) {
	return c.db.ListOperator(ctx, pagination)
}
