package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/models"
)

var _ OperatorsController = (*OperatorsControllerImpl)(nil)

// CreateOperatorCommand ...
type CreateOperatorCommand struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=1024"`
}

// CreateOperatorSigningKeyGroupCommand ...
type CreateOperatorSigningKeyGroupCommand struct {
	OperatorID  uuid.UUID `json:"operator_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=3,max=255"`
	Description string    `json:"description" validate:"max=1024"`
}

// GetOperatorQuery ...
type GetOperatorQuery struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// GetOperatorTokenQuery ...
type GetOperatorTokenQuery struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// DeleteOperatorCommand ...
type DeleteOperatorCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// ListOperatorsQuery ...
type ListOperatorsQuery struct {
	Limit  int    `json:"limit" validate:"required"`
	Offset int    `json:"offset" validate:"required"`
	Search string `json:"search"`
	Sort   string `json:"sort"`
}

// UpdateOperatorSystemAccountCommand ..,
type UpdateOperatorSystemAccountCommand struct {
	OperatorID uuid.UUID `json:"operator_id" validate:"required"`
	AccountID  uuid.UUID `json:"system_id" validate:"required"`
}

// GetOperatorSystemAccountQuery ...
type GetOperatorSystemAccountQuery struct {
	OperatorID uuid.UUID `json:"operator_id" validate:"required"`
}

// OperatorsController is the interface that wraps the methods to access operators.
type OperatorsController interface {
	// CreateOperator creates a new operator.
	CreateOperator(ctx context.Context, cmd CreateOperatorCommand) (models.Operator, error)
	// GetOperator gets an operator.
	GetOperator(ctx context.Context, query GetOperatorQuery) (models.Operator, error)
	// CreateOperatorSigningKeyGroup creates a new signing key group.
	CreateOperatorSigningKeyGroup(ctx context.Context, cmd CreateOperatorSigningKeyGroupCommand) (models.SigningKeyGroup, error)
	// GetOperatorToken gets an operator token.
	GetOperatorToken(ctx context.Context, query GetOperatorTokenQuery) (models.Token, error)
	// ListOperators lists operators.
	ListOperators(ctx context.Context, query ListOperatorsQuery) (models.Pagination[models.Operator], error)
	// DeleteOperator deletes an operator.
	DeleteOperator(ctx context.Context, cmd DeleteOperatorCommand) error
	// UpdateOperatorSystemAccount ...
	UpdateOperatorSystemAccount(ctx context.Context, cmd UpdateOperatorSystemAccountCommand) (models.Account, error)
	// GetOperatorSystemAccount ...
	GetOperatorSystemAccount(ctx context.Context, query GetOperatorSystemAccountQuery) (models.Account, error)
}

// OperatorsControllerImpl is the controller for operators.
type OperatorsControllerImpl struct {
	db ports.Repositories
}

// NewOperatorsController returns a new OperatorsController.
func NewOperatorsController(db ports.Repositories) *OperatorsControllerImpl {
	return &OperatorsControllerImpl{db}
}

// CreateOperator is the method to create a new operator.
func (c *OperatorsControllerImpl) CreateOperator(ctx context.Context, cmd CreateOperatorCommand) (models.Operator, error) {
	op := models.Operator{
		Name:        cmd.Name,
		Description: cmd.Description,
	}

	pk, err := nkeys.CreateOperator()
	if err != nil {
		return op, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return op, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return op, err
	}

	// Create a token for the operator
	oc := jwt.NewOperatorClaims(id)
	oc.Name = cmd.Name

	token, err := oc.Encode(pk)
	if err != nil {
		return op, err
	}

	op.Key = models.NKey{ID: id, Seed: seed}
	op.Token = models.Token{ID: id, Token: token}

	err = c.db.CreateOperator(ctx, &op)
	if err != nil {
		return op, err
	}

	return op, nil
}

// GetOperator ...
func (c *OperatorsControllerImpl) GetOperator(ctx context.Context, query GetOperatorQuery) (models.Operator, error) {
	op := models.Operator{ID: query.ID}

	err := c.db.GetOperator(ctx, &op)
	if err != nil {
		return op, err
	}

	return op, nil
}

// ListOperators is the method to list operators.
func (c *OperatorsControllerImpl) ListOperators(ctx context.Context, query ListOperatorsQuery) (models.Pagination[models.Operator], error) {
	ops := models.Pagination[models.Operator]{}

	ops.Limit = query.Limit
	ops.Offset = query.Offset
	ops.Search = query.Search

	err := c.db.ListOperators(ctx, &ops)
	if err != nil {
		return ops, err
	}

	return ops, nil
}

// GetOperatorToken ...
func (c *OperatorsControllerImpl) GetOperatorToken(ctx context.Context, query GetOperatorTokenQuery) (models.Token, error) {
	op := models.Operator{ID: query.ID}

	err := c.db.GetOperator(ctx, &op)
	if err != nil {
		return op.Token, err
	}

	return op.Token, nil
}

// DeleteOperator ...
func (c *OperatorsControllerImpl) DeleteOperator(ctx context.Context, cmd DeleteOperatorCommand) error {
	op := models.Operator{ID: cmd.ID}

	return c.db.DeleteOperator(ctx, &op)
}

// CreateOperatorSigningKeyGroup ...
func (c *OperatorsControllerImpl) CreateOperatorSigningKeyGroup(ctx context.Context, cmd CreateOperatorSigningKeyGroupCommand) (models.SigningKeyGroup, error) {
	op := models.Operator{ID: cmd.OperatorID}
	skg := models.SigningKeyGroup{Name: cmd.Name, Description: cmd.Description}

	err := c.db.GetOperator(ctx, &op)
	if err != nil {
		return skg, err
	}

	pk, err := nkeys.CreateOperator()
	if err != nil {
		return skg, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return skg, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return skg, err
	}
	skg.Key = models.NKey{ID: id, Seed: seed}

	op.SigningKeyGroups = append(op.SigningKeyGroups, skg)

	oc := jwt.NewOperatorClaims(id)
	oc.Name = op.Name

	for _, sk := range op.SigningKeyGroups {
		oc.SigningKeys.Add(sk.Key.ID, sk.Key.ID, sk.Key.ID)
	}

	token, err := oc.Encode(pk)
	if err != nil {
		return skg, err
	}
	op.Token = models.Token{
		ID:    id,
		Token: token,
	}

	err = c.db.UpdateOperator(ctx, &op)
	if err != nil {
		return skg, err
	}

	return skg, nil
}

// GetOperatorSystemAccount ...
func (c *OperatorsControllerImpl) GetOperatorSystemAccount(ctx context.Context, query GetOperatorSystemAccountQuery) (models.Account, error) {
	op := models.Operator{ID: query.OperatorID}
	ac := models.Account{}

	err := c.db.GetOperator(ctx, &op)
	if err != nil {
		return ac, err
	}

	return ac, nil
}

// UpdateOperatorSystemAccount ...
func (c *OperatorsControllerImpl) UpdateOperatorSystemAccount(ctx context.Context, cmd UpdateOperatorSystemAccountCommand) (models.Account, error) {
	op := models.Operator{ID: cmd.OperatorID}
	ac := models.Account{ID: cmd.AccountID}

	err := c.db.GetAccount(ctx, &ac)
	if err != nil {
		return ac, err
	}

	err = c.db.GetOperator(ctx, &op)
	if err != nil {
		return ac, err
	}

	// op.SystemAdminAccountID = utils.PtrUUID(cmd.AccountID)

	pk, err := nkeys.FromSeed(op.Key.Seed)
	if err != nil {
		return ac, err
	}

	oc := jwt.NewOperatorClaims(op.Key.ID)
	oc.Name = op.Name
	oc.SystemAccount = ac.Key.ID

	for _, sk := range op.SigningKeyGroups {
		oc.SigningKeys.Add(sk.Key.ID, sk.Key.ID, sk.Key.ID)
	}

	token, err := oc.Encode(pk)
	if err != nil {
		return ac, err
	}

	op.Token = models.Token{
		ID:    op.Key.ID,
		Token: token,
	}

	err = c.db.UpdateOperator(ctx, &op)
	if err != nil {
		return ac, err
	}

	return ac, nil
}
