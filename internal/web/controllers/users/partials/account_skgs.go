package partials

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// AccountSkgsOptionsImpl ...
type AccountSkgsOptionsImpl struct {
	AccountID uuid.UUID `json:"account_id" form:"account_id" query:"account_id" validate:"required,uuid"`

	ports.Accounts
	htmx.DefaultController
}

// NewAccountSkgsOptions ...
func NewAccountSkgsOptions(db ports.Accounts) *AccountSkgsOptionsImpl {
	return &AccountSkgsOptionsImpl{
		Accounts:          db,
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
	account := models.Account{
		ID: l.AccountID,
	}
	err := l.GetAccount(l.Context(), &account)
	if err != nil {
		return err
	}

	skgs := make([]*models.SigningKeyGroup, 0)
	for _, skg := range account.SigningKeyGroups {
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
			htmx.ID("account-skgs"),
			htmx.Name("account_skgs_id"),
			htmx.Group(
				htmx.ForEach(skgs, func(e *models.SigningKeyGroup) htmx.Node {
					return htmx.Option(
						htmx.Attribute("value", e.ID.String()),
						htmx.Text(e.Name),
					)
				})...,
			),
		),
	)
}
