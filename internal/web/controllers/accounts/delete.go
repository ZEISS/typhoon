package accounts

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteAccountControllerImpl{})

// DeleteAccountControllerImpl ...
type DeleteAccountControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" params:"id" validate:"required,uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewDeleteAccountController ...
func NewDeleteAccountController(store ports.Datastore) *DeleteAccountControllerImpl {
	return &DeleteAccountControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Delete ...
func (l *DeleteAccountControllerImpl) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	account := models.Account{ID: l.ID}
	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetAccount(ctx, &account)
	})
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/accounts")

	return nil
}
