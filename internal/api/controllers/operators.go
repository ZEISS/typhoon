package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// OperatorsController ...
type OperatorsController struct {
	db ports.Repositories
}

// NewOperatorsController ...
func NewOperatorsController(db ports.Repositories) *OperatorsController {
	return &OperatorsController{db}
}

// CreateOperator ...
func (c *OperatorsController) CreateOperator(ctx context.Context, name string) (*models.Operator, error) {
	pk, err := nkeys.CreateOperator()
	if err != nil {
		return nil, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return nil, err
	}

	op := &models.Operator{
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

// CreateOperatorAccount ...
func (c *OperatorsController) CreateOperatorAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error) {
	key, err := nkeys.CreateAccount()
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

	account := &models.Account{
		Name:       name,
		OperatorID: operatorID,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
	}

	err = c.db.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateOperatorSigningKey ...
func (c *OperatorsController) CreateOperatorSigningKey(ctx context.Context, operatorID uuid.UUID) (*models.Operator, error) {
	return nil, nil
}

// DeleteOperator ...
func (c *OperatorsController) DeleteOperator(ctx context.Context, id uuid.UUID) error {
	return c.db.DeleteOperator(ctx, id)
}

// GetOperator ...
func (c *OperatorsController) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	return c.db.GetOperator(ctx, id)
}

// ListOperator ...
func (c *OperatorsController) ListOperator(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error) {
	return c.db.ListOperator(ctx, pagination)
}
