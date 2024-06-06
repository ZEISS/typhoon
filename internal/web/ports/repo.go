package ports

import (
	"context"
	"io"

	"github.com/zeiss/typhoon/internal/api/models"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// Migration is a method that runs the migration.
type Migration interface {
	// Migrate is a method that runs the migration.
	Migrate(context.Context) error
}

// Datastore provides methods for transactional operations.
type Datastore interface {
	// ReadTx starts a read only transaction.
	ReadTx(context.Context, func(context.Context, ReadTx) error) error
	// ReadWriteTx starts a read write transaction.
	ReadWriteTx(context.Context, func(context.Context, ReadWriteTx) error) error

	io.Closer
	Migration
}

// ReadTx provides methods for transactional read operations.
type ReadTx interface {
	// GetOperator is a method that returns an operator by ID
	GetOperator(ctx context.Context, operator *models.Operator) error
	// ListOperators is a method that returns a list of operators
	ListOperators(ctx context.Context, pagination *tables.Results[models.Operator]) error
	// GetAccount ...
	GetAccount(ctx context.Context, account *models.Account) error
	// ListAccounts ...
	ListAccounts(ctx context.Context, pagination *models.Pagination[models.Account]) error
	// GetUser is a method that returns a user by ID
	GetUser(ctx context.Context, user *models.User) error
	// ListUsers is a method that returns a list of users
	ListUsers(ctx context.Context, pagination *models.Pagination[models.User]) error
	// GetProfile is a method that returns the profile of the current user
	GetProfile(ctx context.Context, user *adapters.GothUser) error
	// GetSystem is a method that returns a system by ID
	GetSystem(ctx context.Context, system *models.System) error
	// ListSystems is a method that returns a list of systems
	ListSystems(ctx context.Context, pagination *models.Pagination[models.System]) error
}

// ReadWriteTx provides methods for transactional read and write operations.
type ReadWriteTx interface {
	ReadTx

	// CreateOperator is a method that creates a new operator
	CreateOperator(ctx context.Context, operator *models.Operator) error
	// UpdateOperator is a method that updates an operator
	UpdateOperator(ctx context.Context, operator *models.Operator) error
	// DeleteOperator is a method that deletes an operator
	DeleteOperator(ctx context.Context, operator *models.Operator) error
	// CreateAccount is creating a new account.
	CreateAccount(ctx context.Context, account *models.Account) error
	// UpdateAccount ...
	UpdateAccount(ctx context.Context, account *models.Account) error
	// DeleteAccount ...
	DeleteAccount(ctx context.Context, account *models.Account) error
	// CreateUser is a method that creates a user
	CreateUser(ctx context.Context, user *models.User) error
	// UpdateUser is a method that updates a user
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser is a method that deletes a user
	DeleteUser(ctx context.Context, user *models.User) error
	// CreateSystem is a method that creates a new system
	CreateSystem(ctx context.Context, system *models.System) error
	// UpdateSystem is a method that updates a system
	UpdateSystem(ctx context.Context, system *models.System) error
	// DeleteSystem is a method that deletes a system
	DeleteSystem(ctx context.Context, system *models.System) error
}