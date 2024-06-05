package ports

import (
	"context"
	"io"

	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	Accounts
	Operators
	Users
	Me
	Conn() *gorm.DB
}

// Datastore provides methods for transactional operations.
type Datastore interface {
	io.Closer

	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error

	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	// GetOperator is a method that returns an operator by ID
	GetOperator(ctx context.Context, operator *models.Operator) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx

	// CreateAccount is creating a new account.
	CreateAccount(ctx context.Context, account *models.Account) error
}
