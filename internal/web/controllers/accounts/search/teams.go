package accounts

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components/alerts"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/pkg/errorx"
)

var _ = htmx.Controller(&SearchTeamsControllerImpl{})

// Search ...
type SearchTeamsControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewSearchTeamsController ...
func NewSearchTeamsController(store ports.Datastore) *SearchTeamsControllerImpl {
	return &SearchTeamsControllerImpl{store: store}
}

// Prepare ...
func (l *SearchTeamsControllerImpl) Get() error {
	return l.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(
				func() htmx.Node {
					results := tables.Results[models.Team]{}
					errorx.Panic(l.BindQuery(&results))

					errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
						return tx.ListTeams(ctx, &results)
					}))

					return htmx.Fragment(
						htmx.ForEach(tables.RowsPtr(results.Rows), func(e *models.Team, idx int) htmx.Node {
							return htmx.Option(
								htmx.Text(e.Name),
								htmx.Value(e.ID.String()),
							)
						})...,
					)
				},
			),
			func(err error) htmx.Node {
				return alerts.Error(alerts.ErrorProps{})
			},
		),
	)
}
