package tokens

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// IndexOperatorTokenControllerImpl ...
type IndexOperatorTokenControllerImpl struct {
	ID uuid.UUID `json:"id" form:"id" param:"id" validate:"required:uuid"`

	ports.Operators
	htmx.DefaultController
}

// NewIndexOperatorTokenController ...
func NewIndexOperatorTokenController(db ports.Operators) *IndexOperatorTokenControllerImpl {
	return &IndexOperatorTokenControllerImpl{
		Operators:         db,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *IndexOperatorTokenControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	operator := models.Operator{ID: l.ID}

	err = l.GetOperator(l.Context(), &operator)
	if err != nil {
		return err
	}

	r := strings.NewReader(operator.Token.Token)

	l.Ctx().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jwt", operator.Name))

	return l.Ctx().SendStream(r)
}
