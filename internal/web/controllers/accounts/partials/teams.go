package partials

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// TeamsOptionsImpl ...
type TeamsOptionsImpl struct {
	teams tables.Results[adapters.GothTeam]
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamsOptions ...
func NewTeamsOptions(store ports.Datastore) *TeamsOptionsImpl {
	return &TeamsOptionsImpl{
		store: store,
	}
}

// Prepare ...
func (l *TeamsOptionsImpl) Prepare() error {
	return l.store.ReadTx(context.Background(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTeams(ctx, &l.teams)
	})
}

// Get ...
func (l *TeamsOptionsImpl) Get() error {
	return l.Render(
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
				htmx.ForEach(l.teams.GetRows(), func(e *adapters.GothTeam, idx int) htmx.Node {
					return htmx.Option(
						htmx.Attribute("value", e.ID.String()),
						htmx.Text(e.Name),
					)
				})...,
			),
		),
	)
}
