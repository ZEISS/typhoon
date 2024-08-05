package accounts

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteAccountControllerImpl{})

// DeleteAccountControllerImpl ...
type DeleteAccountControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewDeleteAccountController ...
func NewDeleteAccountController(store ports.Datastore) *DeleteAccountControllerImpl {
	return &DeleteAccountControllerImpl{store: store}
}

// Delete ...
func (l *DeleteAccountControllerImpl) Delete() error {
	var params struct {
		ID uuid.UUID `param:"id"`
	}

	err := l.BindParams(&params)
	if err != nil {
		return err
	}

	return l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteAccount(ctx, &models.Account{ID: params.ID})
	})
}
