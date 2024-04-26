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
	// CreateOperatorToken creates a new token for the operator.
	CreateOperatorToken(ctx context.Context, operatorID uuid.UUID, token *models.Token) error
	// CreateOperatorAccount creates a new account for the operator.
	CreateOperatorAccount(ctx context.Context, account *models.Account) error
	// GetOperatorAccount returns the account with the given ID.
	GetOperatorAccount(ctx context.Context, id uuid.UUID) (*models.Account, error)
	// CreateOperatorAccountToken creates a new token for the operator account.
	CreateOperatorAccountToken(ctx context.Context, accountID uuid.UUID, token *models.Token) error
	// CreateOperatorAccountSigningKey creates a new signing key for the operator account.
	CreateOperatorAccountSigningKey(ctx context.Context, accountID uuid.UUID, key *models.NKey) error
	// ListOperatorAccounts ...
	ListOperatorAccounts(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error)
	// ListOperatorAccountsSigningKey returns a list of accounts for the operator.
	ListOperatorAccountsSigningKey(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error)
	// CreateOperatorAccountUser creates a new user for the operator account.
	CreateOperatorAccountUser(ctx context.Context, user *models.User) error
	// GetOperatorAccountUser returns the user with the given ID.
	GetOperatorAccountUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	// CreateOperatorAccountUserToken creates a new token for the operator account user.
	CreateOperatorAccountUserToken(ctx context.Context, userID uuid.UUID, token *models.Token) error
}
