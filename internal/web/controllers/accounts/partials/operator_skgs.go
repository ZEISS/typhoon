package partials

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// OperatorSkgsOptionsImpl ...
type OperatorSkgsOptionsImpl struct {
	operator models.Operator
	store    ports.Datastore
	htmx.DefaultController
}

// NewOperatorSkgsOptions ...
func NewOperatorSkgsOptions(store ports.Datastore) *OperatorSkgsOptionsImpl {
	return &OperatorSkgsOptionsImpl{store: store}
}

// Error ...
func (l *OperatorSkgsOptionsImpl) Error(err error) error {
	return toasts.Error(err.Error())
}

// Prepare ...
func (l *OperatorSkgsOptionsImpl) Prepare() error {
	var operator struct {
		ID uuid.UUID `json:"operator_id" form:"operator_id" query:"operator_id" validate:"required,uuid"`
	}

	err := l.Ctx().QueryParser(&operator)
	if err != nil {
		return err
	}

	l.operator.ID = operator.ID

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetOperator(ctx, &l.operator)
	})
}

// Get ...
func (l *OperatorSkgsOptionsImpl) Get() error {
	return l.Render(
		forms.SelectBordered(
			forms.SelectProps{},
			htmx.HxGet("/accounts/partials/operator-skgs"),
			htmx.HxTrigger("change from:input[name=operator]"),
			forms.Option(
				forms.OptionProps{
					Selected: true,
					Disabled: true,
				},
				htmx.Text("Select an signing key group"),
			),
			htmx.Target("this"),
			htmx.HxValidate(true),
			htmx.HxInclude("[name='operator_id']"),
			htmx.ID("operator-skgs"),
			htmx.Name("operator_skgs_id"),
			htmx.Group(
				htmx.ForEach(tables.RowsPtr(l.operator.SigningKeyGroups), func(e *models.SigningKeyGroup, idx int) htmx.Node {
					return htmx.Option(
						htmx.Attribute("value", e.KeyID),
						htmx.Text(e.Name),
						htmx.Value(e.KeyID),
					)
				})...,
			),
		),
	)
}
