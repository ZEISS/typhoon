package ports

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// Operators is a port that defines the methods for operators
type Operators interface {
	// ListOperators is a method that returns a list of operators
	ListOperators(ctx context.Context, pagination *models.Pagination[models.Operator]) error
}
