package users

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteUserControllerImpl{})

// DeleteUserControllerImpl ...
type DeleteUserControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" param:"id" validate:"required,uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewDeleteUserController ...
func NewDeleteUserController(store ports.Datastore) *DeleteUserControllerImpl {
	return &DeleteUserControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Delete ...
func (l *DeleteUserControllerImpl) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	user := models.User{ID: l.ID}
	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteUser(l.Context(), &user)
	})
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/users")

	return nil
}
