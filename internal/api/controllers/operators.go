package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/utils"
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

	if op.SystemAdminAccount != nil {
		ac = *op.SystemAdminAccount
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

	op.SystemAdminAccountID = utils.PtrUUID(cmd.AccountID)

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

// CreateOperatorAccount ...
// func (c *OperatorsController) CreateOperatorAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error) {
// 	key, err := nkeys.CreateAccount()
// 	if err != nil {
// 		return nil, err
// 	}

// 	id, err := key.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	seed, err := key.Seed()
// 	if err != nil {
// 		return nil, err
// 	}

// 	account := &models.Account{
// 		Name:       name,
// 		OperatorID: operatorID,
// 		Key: models.NKey{
// 			ID:   id,
// 			Seed: seed,
// 		},
// 	}

// 	err = c.db.CreateAccount(ctx, account)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return account, nil
// }

// CreateOperatorSigningKey ...
// func (c *OperatorsController) CreateOperatorSigningKey(ctx context.Context, operatorID uuid.UUID) (*models.NKey, error) {
// 	nkey, err := nkeys.CreateOperator()
// 	if err != nil {
// 		return nil, err
// 	}

// 	pk, err := nkey.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	seed, err := nkey.Seed()
// 	if err != nil {
// 		return nil, err
// 	}

// 	key := &models.NKey{
// 		ID:   pk,
// 		Seed: seed,
// 	}

// 	err = c.db.CreateOperatorSigningKey(ctx, operatorID, key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }

// CreateOperatorToken ...
// func (c *OperatorsController) CreateOperatorToken(ctx context.Context, operatorID uuid.UUID) (*models.Token, error) {
// 	o, err := c.db.GetOperator(ctx, operatorID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	okp, err := nkeys.FromSeed(o.Key.Seed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	opk, err := okp.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	oc := jwt.NewOperatorClaims(opk)
// 	oc.Name = o.Name

// 	oc.SigningKeys = make([]string, len(o.SigningKeys))
// 	for i, k := range o.SigningKeys {
// 		oc.SigningKeys[i] = k.ID
// 	}

// 	token, err := oc.Encode(okp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	t := &models.Token{
// 		TokenID: opk,
// 		Token:   token,
// 	}

// 	err = c.db.CreateOperatorToken(ctx, operatorID, t)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return t, nil
// }

// // ListOperatorAccount ...
// func (c *OperatorsController) ListOperatorAccount(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error) {
// 	return c.db.ListOperatorAccounts(ctx, operatorID, pagination)
// }

// // ListOperatorAccountUsers ...
// func (c *OperatorsController) ListOperatorAccountUsers(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[*models.User]) (*models.Pagination[*models.User], error) {
// 	return c.db.ListOperatorAccountUsers(ctx, accountID, pagination)
// }

// // CreateOperatorAccountToken ...
// func (c *OperatorsController) CreateOperatorAccountToken(ctx context.Context, accountID uuid.UUID) (*models.Token, error) {
// 	a, err := c.db.GetOperatorAccount(ctx, accountID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	akp, err := nkeys.FromSeed(a.Key.Seed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	apk, err := akp.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ac := jwt.NewAccountClaims(apk)
// 	ac.Name = a.Name

// 	ac.SigningKeys = make([]string, len(a.SigningKeys))
// 	for i, k := range a.SigningKeys {
// 		ac.SigningKeys[i] = k.ID
// 	}

// 	token, err := ac.Encode(akp)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	t := &models.Token{
// 		TokenID: apk,
// 		Token:   token,
// 	}

// 	err = c.db.CreateOperatorAccountToken(ctx, accountID, t)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return t, nil
// }

// // CreateOperatorAccountUser ...
// func (c *OperatorsController) CreateOperatorAccountUser(ctx context.Context, accountID uuid.UUID, name string) (*models.User, error) {
// 	key, err := nkeys.CreateUser()
// 	if err != nil {
// 		return nil, err
// 	}

// 	id, err := key.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	seed, err := key.Seed()
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := &models.User{
// 		Name:      name,
// 		AccountID: accountID,
// 		Key: models.NKey{
// 			ID:   id,
// 			Seed: seed,
// 		},
// 	}

// 	err = c.db.CreateOperatorAccountUser(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// // CreateOperatorAccountUserToken ...
// func (c *OperatorsController) CreateOperatorAccountUserToken(ctx context.Context, userID uuid.UUID) (*models.Token, error) {
// 	u, err := c.db.GetOperatorAccountUser(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ukp, err := nkeys.FromSeed(u.Key.Seed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	upk, err := ukp.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(u.Account.SigningKeys) < 1 {
// 		return nil, fmt.Errorf("account %s has no signing keys", u.AccountID)
// 	}

// 	ask, err := nkeys.FromSeed(u.Account.SigningKeys[0].Seed)
// 	if err != nil {
// 		return nil, err
// 	}

// 	uc := jwt.NewUserClaims(upk)
// 	uc.Name = u.Name
// 	uc.IssuerAccount = u.Account.KeyID

// 	token, err := uc.Encode(ask)
// 	if err != nil {
// 		return nil, err
// 	}

// 	t := &models.Token{
// 		TokenID: upk,
// 		Token:   token,
// 	}

// 	err = c.db.CreateOperatorAccountUserToken(ctx, userID, t)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return t, nil
// }

// // DeleteOperator ...
// func (c *OperatorsController) DeleteOperator(ctx context.Context, id uuid.UUID) error {
// 	return c.db.DeleteOperator(ctx, id)
// }

// // GetOperator ...
// func (c *OperatorsController) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
// 	return c.db.GetOperator(ctx, id)
// }

// // CreateOperatorAccountSigningKey ...
// func (c *OperatorsController) CreateOperatorAccountSigningKey(ctx context.Context, accountID uuid.UUID) (*models.NKey, error) {
// 	nkey, err := nkeys.CreateAccount()
// 	if err != nil {
// 		return nil, err
// 	}

// 	pk, err := nkey.PublicKey()
// 	if err != nil {
// 		return nil, err
// 	}

// 	seed, err := nkey.Seed()
// 	if err != nil {
// 		return nil, err
// 	}

// 	key := &models.NKey{
// 		ID:   pk,
// 		Seed: seed,
// 	}

// 	err = c.db.CreateOperatorAccountSigningKey(ctx, accountID, key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }
