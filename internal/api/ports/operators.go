package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Operators is the interface that wraps the methods to access data.
type Operators interface {
	// GetOperator returns the operator with the given ID.
	GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error)
	// CreateOperator creates a new operator.
	CreateOperator(ctx context.Context, operator *models.Operator) error
	// ListOperator returns a list of operators.
	ListOperator(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error)
	// DeleteOperator deletes the operator with the given ID.
	DeleteOperator(ctx context.Context, id uuid.UUID) error
	// CreateOperatorSigningKey creates a new signing key for the operator.
	CreateOperatorSigningKey(ctx context.Context, operatorID uuid.UUID, key *models.NKey) error
}
