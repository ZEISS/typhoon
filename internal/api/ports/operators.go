package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// Operators is the interface that wraps the methods to access data.
type Operators interface {
	// CreateOperator creates a new operator.
	CreateOperator(ctx context.Context, operator *models.Operator) error
	// UpdateOperator updates an existing operator.
	UpdateOperator(ctx context.Context, operator *models.Operator) error
	// GetOperator returns the operator with the given ID.
	GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error)
	// ListOperators returns a list of operators.
	ListOperators(ctx context.Context, pagination models.OperatorPagination) (models.OperatorPagination, error)
}
