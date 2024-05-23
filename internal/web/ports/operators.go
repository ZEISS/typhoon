package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Operators is a port that defines the methods for operators
type Operators interface {
	// ListOperators is a method that returns a list of operators
	ListOperators(ctx context.Context, pagination *models.Pagination[models.Operator]) error
	// CreateOperator is a method that creates a new operator
	CreateOperator(ctx context.Context, operator *models.Operator) error
	// GetOperator is a method that returns an operator by ID
	GetOperator(ctx context.Context, operator *models.Operator) error
	// UpdateOperator is a method that updates an operator
	UpdateOperator(ctx context.Context, operator *models.Operator) error
	// DeleteOperator is a method that deletes an operator
	DeleteOperator(ctx context.Context, operator *models.Operator) error
}
