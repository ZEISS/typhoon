package operators

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteOperatorController{})

// DeleteOperatorsController ...
type DeleteOperatorController struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewDeleteOperatorController ...
func NewDeleteOperatorController(store ports.Datastore) *DeleteOperatorController {
	return &DeleteOperatorController{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Error ...
func (l *DeleteOperatorController) Error(err error) error {
	return nil
}

// Delete ...
func (l *DeleteOperatorController) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	op := models.Operator{ID: l.ID}
	err = l.store.ReadWriteTx(l.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
		return tx.DeleteOperator(l.Context(), &op)
	})
	if err != nil {
		return err
	}

	return l.Redirect("/operators")
}
