package partials

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// OperatorSkgsOptionsImpl ...
type OperatorSkgsOptionsImpl struct {
	OperatorID uuid.UUID `json:"operator_id" form:"operator_id" query:"operator_id" validate:"required,uuid"`

	ports.Operators
	htmx.DefaultController
}

// NewOperatorSkgsOptions ...
func NewOperatorSkgsOptions(db ports.Operators) *OperatorSkgsOptionsImpl {
	return &OperatorSkgsOptionsImpl{
		Operators:         db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *OperatorSkgsOptionsImpl) Prepare() error {
	err := l.Ctx().QueryParser(l)
	if err != nil {
		return err
	}

	return nil
}

// Error ...
func (l *OperatorSkgsOptionsImpl) Error(err error) error {
	fmt.Println(err)
	return nil
}

// Get ...
func (l *OperatorSkgsOptionsImpl) Get() error {
	operator := models.Operator{
		ID: l.OperatorID,
	}
	err := l.GetOperator(l.Context(), &operator)
	if err != nil {
		fmt.Println(err)
		return err
	}

	skgs := make([]*models.SigningKeyGroup, 0)
	for _, skg := range operator.SigningKeyGroups {
		skgs = append(skgs, &skg)
	}

	return htmx.RenderComp(
		l.Ctx(),
		forms.SelectBordered(
			forms.SelectProps{},
			forms.Option(
				forms.OptionProps{
					Selected: true,
					Disabled: true,
				},
				htmx.Text("Select an signing key group"),
			),
			htmx.ID("operator-skgs"),
			htmx.Name("operator_skgs_id"),
			htmx.Group(
				htmx.ForEach(skgs, func(e *models.SigningKeyGroup) htmx.Node {
					return htmx.Option(
						htmx.Attribute("value", e.KeyID),
						htmx.Text(e.Name),
					)
				})...,
			),
		),
	)
}
