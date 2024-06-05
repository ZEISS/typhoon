package partials

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// OperatorSkgsOptionsImpl ...
type OperatorSkgsOptionsImpl struct {
	OperatorID uuid.UUID `json:"operator_id" form:"operator_id" query:"operator_id" validate:"required,uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewOperatorSkgsOptions ...
func NewOperatorSkgsOptions(store ports.Datastore) *OperatorSkgsOptionsImpl {
	return &OperatorSkgsOptionsImpl{
		store:             store,
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

// Get ...
func (l *OperatorSkgsOptionsImpl) Get() error {
	operator := models.Operator{
		ID: l.OperatorID,
	}

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetOperator(ctx, &operator)
	})
	if err != nil {
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
