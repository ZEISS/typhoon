package operators

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteOperatorController{})

// DeleteOperatorsController ...
type DeleteOperatorController struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`
	ports.Operators
	htmx.DefaultController
}

// NewDeleteOperatorController ...
func NewDeleteOperatorController(db ports.Operators) *DeleteOperatorController {
	return &DeleteOperatorController{
		Operators:         db,
		DefaultController: htmx.DefaultController{},
	}
}

// Delete ...
func (l *DeleteOperatorController) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	op := models.Operator{ID: l.ID}
	err = l.DeleteOperator(l.Context(), &op)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/operators")

	return nil
}
