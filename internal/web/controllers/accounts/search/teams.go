package search

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components/alerts"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/pkg/slices"
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

// Error ...
func (l *SearchTeamsControllerImpl) Error(err error) error {
	return toasts.RenderToasts(
		l.Ctx(),
		toasts.Toasts(
			toasts.ToastsProps{},
			toasts.ToastAlertError(
				toasts.ToastProps{},
				htmx.Text(err.Error()),
			),
		),
	)
}

// Prepare ...
func (l *SearchTeamsControllerImpl) Post() error {
	return l.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(
				func() htmx.Node {
					results := tables.Results[models.Team]{}
					errorx.Panic(l.BindQuery(&results))

					errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
						return tx.ListTeams(ctx, &results)
					}))

					return htmx.IfElse(
						!slices.Size(0, results.Rows...),
						htmx.Group(
							htmx.ForEach(tables.RowsPtr(results.Rows), func(e *models.Team, idx int) htmx.Node {
								return dropdowns.DropdownMenuItem(
									dropdowns.DropdownMenuItemProps{},
									htmx.A(
										htmx.Text(e.Name),
										htmx.Value(e.ID.String()),
										alpine.XOn("click", "onOptionClick($event.target.getAttribute('value'), $event.target.innerText)"),
									),
								)
							})...,
						),
						dropdowns.DropdownMenuItem(
							dropdowns.DropdownMenuItemProps{},
							htmx.A(
								htmx.Text("No teams found"),
							),
						),
					)
				},
			),
			func(err error) htmx.Node {
				return alerts.Error(alerts.ErrorProps{})
			},
		),
	)
}
