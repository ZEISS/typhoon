package controllers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// OperatorsController ...
type OperatorsController struct {
	db ports.Repositories
}

// NewOperatorsController ...
func NewOperatorsController(db ports.Repositories) *OperatorsController {
	return &OperatorsController{db}
}

// CreateOperator ...
func (c *OperatorsController) CreateOperator(ctx context.Context, name string) (*models.Operator, error) {
	pk, err := nkeys.CreateOperator()
	if err != nil {
		return nil, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return nil, err
	}

	op := &models.Operator{
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
	}

	err = c.db.CreateOperator(ctx, op)
	if err != nil {
		return nil, err
	}

	return op, nil
}

// CreateOperatorAccount ...
func (c *OperatorsController) CreateOperatorAccount(ctx context.Context, name string, operatorID uuid.UUID) (*models.Account, error) {
	key, err := nkeys.CreateAccount()
	if err != nil {
		return nil, err
	}

	id, err := key.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := key.Seed()
	if err != nil {
		return nil, err
	}

	account := &models.Account{
		Name:       name,
		OperatorID: operatorID,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
	}

	err = c.db.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// CreateOperatorSigningKey ...
func (c *OperatorsController) CreateOperatorSigningKey(ctx context.Context, operatorID uuid.UUID) (*models.NKey, error) {
	nkey, err := nkeys.CreateOperator()
	if err != nil {
		return nil, err
	}

	pk, err := nkey.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := nkey.Seed()
	if err != nil {
		return nil, err
	}

	key := &models.NKey{
		ID:   pk,
		Seed: seed,
	}

	err = c.db.CreateOperatorSigningKey(ctx, operatorID, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// CreateOperatorToken ...
func (c *OperatorsController) CreateOperatorToken(ctx context.Context, operatorID uuid.UUID) (*models.Token, error) {
	o, err := c.db.GetOperator(ctx, operatorID)
	if err != nil {
		return nil, err
	}

	okp, err := nkeys.FromSeed(o.Key.Seed)
	if err != nil {
		return nil, err
	}

	opk, err := okp.PublicKey()
	if err != nil {
		return nil, err
	}

	oc := jwt.NewOperatorClaims(opk)
	oc.Name = o.Name

	oc.SigningKeys = make([]string, len(o.SigningKeys))
	for i, k := range o.SigningKeys {
		oc.SigningKeys[i] = k.ID
	}

	token, err := oc.Encode(okp)
	if err != nil {
		return nil, err
	}

	t := &models.Token{
		ID:    opk,
		Token: token,
	}

	err = c.db.CreateOperatorToken(ctx, operatorID, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// ListOperatorAccount ...
func (c *OperatorsController) ListOperatorAccount(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error) {
	return c.db.ListOperatorAccounts(ctx, operatorID, pagination)
}

// CreateOperatorAccountToken ...
func (c *OperatorsController) CreateOperatorAccountToken(ctx context.Context, accountID uuid.UUID) (*models.Token, error) {
	a, err := c.db.GetOperatorAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	akp, err := nkeys.FromSeed(a.Key.Seed)
	if err != nil {
		return nil, err
	}

	apk, err := akp.PublicKey()
	if err != nil {
		return nil, err
	}

	ac := jwt.NewAccountClaims(apk)
	ac.Name = a.Name

	ac.SigningKeys = make([]string, len(a.SigningKeys))
	for i, k := range a.SigningKeys {
		ac.SigningKeys[i] = k.ID
	}

	token, err := ac.Encode(akp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	t := &models.Token{
		ID:    apk,
		Token: token,
	}

	err = c.db.CreateOperatorAccountToken(ctx, accountID, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// CreateOperatorAccountUser ...
func (c *OperatorsController) CreateOperatorAccountUser(ctx context.Context, accountID uuid.UUID, name string) (*models.User, error) {
	key, err := nkeys.CreateUser()
	if err != nil {
		return nil, err
	}

	id, err := key.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := key.Seed()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      name,
		AccountID: accountID,
		Key: models.NKey{
			ID:   id,
			Seed: seed,
		},
	}

	err = c.db.CreateOperatorAccountUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateOperatorAccountUserToken ...
func (c *OperatorsController) CreateOperatorAccountUserToken(ctx context.Context, userID uuid.UUID) (*models.Token, error) {
	u, err := c.db.GetOperatorAccountUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ukp, err := nkeys.FromSeed(u.Key.Seed)
	if err != nil {
		return nil, err
	}

	upk, err := ukp.PublicKey()
	if err != nil {
		return nil, err
	}

	uc := jwt.NewUserClaims(upk)
	uc.Name = u.Name

	uc.IssuerAccount = u.AccountID.String()

	token, err := uc.Encode(ukp)
	if err != nil {
		return nil, err
	}

	t := &models.Token{
		ID:    upk,
		Token: token,
	}

	err = c.db.CreateOperatorAccountUserToken(ctx, userID, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// DeleteOperator ...
func (c *OperatorsController) DeleteOperator(ctx context.Context, id uuid.UUID) error {
	return c.db.DeleteOperator(ctx, id)
}

// GetOperator ...
func (c *OperatorsController) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	return c.db.GetOperator(ctx, id)
}

// ListOperator ...
func (c *OperatorsController) ListOperator(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error) {
	return c.db.ListOperator(ctx, pagination)
}

// CreateOperatorAccountSigningKey ...
func (c *OperatorsController) CreateOperatorAccountSigningKey(ctx context.Context, accountID uuid.UUID) (*models.NKey, error) {
	nkey, err := nkeys.CreateAccount()
	if err != nil {
		return nil, err
	}

	pk, err := nkey.PublicKey()
	if err != nil {
		return nil, err
	}

	seed, err := nkey.Seed()
	if err != nil {
		return nil, err
	}

	key := &models.NKey{
		ID:   pk,
		Seed: seed,
	}

	err = c.db.CreateOperatorAccountSigningKey(ctx, accountID, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}
