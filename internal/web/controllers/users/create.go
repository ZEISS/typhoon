package users

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var validate *validator.Validate

// CreateUserControllerImpl ...
type CreateUserControllerImpl struct {
	AccountID                uuid.UUID `json:"account_id" form:"account_id" validate:"required,uuid"`
	AccountSigningKeyGroupID uuid.UUID `json:"account_skgs_id" form:"account_skgs_id" validate:"required,uuid"`
	Name                     string    `json:"name" form:"name" validate:"required,min=3,max=100"`
	Description              string    `json:"description" form:"description" validate:"required,min=3,max=1024"`

	store ports.Datastore
	htmx.DefaultController
}

// NewCreateUserController ...
func NewCreateUserController(store ports.Datastore) *CreateUserControllerImpl {
	return &CreateUserControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateUserControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.Ctx().BodyParser(l)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}

// Error ...
func (l *CreateUserControllerImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Post ...
// nolint:gocyclo
func (l *CreateUserControllerImpl) Post() error {
	user := models.User{
		Name:        l.Name,
		Description: l.Description,
	}
	account := models.Account{ID: l.AccountID, SigningKeyGroups: []models.SigningKeyGroup{{ID: l.AccountSigningKeyGroupID}}}

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetAccount(ctx, &account)
	})
	if err != nil {
		return err
	}
	user.Account = account

	pk, err := nkeys.CreateUser()
	if err != nil {
		return err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return err
	}

	seed, err := pk.Seed()
	if err != nil {
		return err
	}
	user.Key = models.NKey{ID: id, Seed: seed}

	if len(account.SigningKeyGroups) < 1 {
		return fmt.Errorf("account %s has no signing keys", account.ID)
	}

	askg := account.FindSigningKeyGroupByID(l.AccountSigningKeyGroupID)
	if askg == nil {
		return fmt.Errorf("account %s does not have signing key group %s", account.ID, l.AccountSigningKeyGroupID)
	}

	ask, err := nkeys.FromSeed(askg.Key.Seed)
	if err != nil {
		return err
	}

	askpk, err := ask.PublicKey()
	if err != nil {
		return err
	}

	// // Create a token for the user
	u := jwt.NewUserClaims(id)
	u.Name = l.Name
	u.IssuerAccount = account.Key.ID
	u.Issuer = askpk

	token, err := u.Encode(ask)
	if err != nil {
		return err
	}
	user.Token = models.Token{ID: id, Token: token}

	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.CreateUser(ctx, &user)
	})
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/users")

	return nil
}
