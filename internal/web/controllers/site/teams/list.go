package teams

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/teams"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListTeamsControllerImpl{})

// ListTeamsControllerImpl ...
type ListTeamsControllerImpl struct {
	teams tables.Results[adapters.GothTeam]
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamsListOperatorController ...
func NewTeamsListOperatorController(store ports.Datastore) *ListTeamsControllerImpl {
	return &ListTeamsControllerImpl{
		teams: tables.Results[adapters.GothTeam]{Limit: 10},
		store: store,
	}
}

// Prepare ...
func (l *ListTeamsControllerImpl) Prepare() error {
	err := l.Ctx().QueryParser(&l.teams)
	if err != nil {
		return nil
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTeams(ctx, &l.teams)
	})
}

// Prepare ...
func (l *ListTeamsControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{
				Title: "Operators",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						teams.TeamsTable(
							teams.TeamsTableProps{
								Teams:  l.teams.GetRows(),
								Offset: l.teams.GetOffset(),
								Limit:  l.teams.GetLimit(),
								Total:  l.teams.GetLen(),
							},
						),
					),
				),
			),
		),
	)
}
