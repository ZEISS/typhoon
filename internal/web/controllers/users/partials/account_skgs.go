package partials

import (
	"context"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// AccountSkgsOptionsImpl ...
type AccountSkgsOptionsImpl struct {
	AccountID uuid.UUID `json:"account_id" form:"account_id" query:"account_id" validate:"required,uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewAccountSkgsOptions ...
func NewAccountSkgsOptions(store ports.Datastore) *AccountSkgsOptionsImpl {
	return &AccountSkgsOptionsImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *AccountSkgsOptionsImpl) Prepare() error {
	err := l.Ctx().QueryParser(l)
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *AccountSkgsOptionsImpl) Get() error {
	return l.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				account := models.Account{ID: l.AccountID}
				err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetAccount(ctx, &account)
				})
				if err != nil {
					panic(err)
				}

				return forms.SelectBordered(
					forms.SelectProps{},
					forms.Option(
						forms.OptionProps{
							Selected: true,
							Disabled: true,
						},
						htmx.Text("Select an signing key group"),
					),
					htmx.ID("account-skgs"),
					htmx.Name("account_skgs_id"),
					htmx.Group(
						htmx.ForEach(tables.RowsPtr(account.SigningKeyGroups), func(e *models.SigningKeyGroup, idx int) htmx.Node {
							return htmx.Option(
								htmx.Attribute("value", e.ID.String()),
								htmx.Text(e.Name),
							)
						})...,
					),
				)
			}),
			func(err error) htmx.Node {
				return forms.SelectBordered(
					forms.SelectProps{},
					forms.Option(
						forms.OptionProps{
							Selected: true,
							Disabled: true,
						},
						htmx.Text("Select an signing key group"),
					),
					htmx.ID("account-skgs"),
					htmx.Name("account_skgs_id"),
					htmx.Option(
						htmx.Attribute("value", ""),
						htmx.Disabled(),
						htmx.Text("No signing key groups found"),
					),
				)
			},
		),
	)
}
