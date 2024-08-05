package systems

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteSystemControllerImpl{})

// DeleteSystemControllerParams ...
type DeleteSystemControllerParams struct {
	ID uuid.UUID `json:"id" params:"id" form:"id" validate:"required:uuid"`
}

// DeleteSystemControllerImpl ...
type DeleteSystemControllerImpl struct {
	Params DeleteSystemControllerParams
	store  ports.Datastore
	htmx.DefaultController
}

// NewDeleteSystemController ...
func NewDeleteSystemController(store ports.Datastore) *DeleteSystemControllerImpl {
	return &DeleteSystemControllerImpl{
		Params:            DeleteSystemControllerParams{},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
}

// Prepare ...
func (l *DeleteSystemControllerImpl) Prepare() error {
	return l.BindParams(&l.Params)
}

// Delete ...
func (l *DeleteSystemControllerImpl) Delete() error {
	sys := models.System{ID: l.Params.ID}

	err := l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteSystem(l.Context(), &sys)
	})
	if err != nil {
		return err
	}

	return l.Redirect("/systems")
}
