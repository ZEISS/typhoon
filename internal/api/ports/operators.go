package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
)

// Operators is the interface that wraps the methods to access data.
type Operators interface {
	// CreateOperator creates a new operator.
	CreateOperator(ctx context.Context, operator *models.Operator) error
	// UpdateOperator updates an existing operator.
	UpdateOperator(ctx context.Context, operator *models.Operator) error
	// GetOperator returns the operator with the given ID.
	GetOperator(ctx context.Context, operator *models.Operator) error
	// ListOperators returns a list of operators.
	ListOperators(ctx context.Context, pagination *models.Pagination[models.Operator]) error
	// DeleteOperator deletes the operator with the given ID.
	DeleteOperator(ctx context.Context, operator *models.Operator) error
}
