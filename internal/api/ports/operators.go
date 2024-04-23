package ports

import "github.com/zeiss/typhoon/internal/api/models"

// Operators is the interface that wraps the methods to access data.
type Operators interface {
	// GetOperator returns the operator with the given ID.
	GetOperator(id string) (*models.Operator, error)
	// CreateOperator creates a new operator.
	CreateOperator(operator *models.Operator) error
}
