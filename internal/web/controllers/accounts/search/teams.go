package search

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/pkg/conv"
)

var _ = htmx.Controller(&SearchTeamsControllerImpl{})

// Search ...
type SearchTeamsControllerImpl struct {
	teams tables.Results[models.Team]
	store ports.Datastore
	htmx.DefaultController
}

// NewSearchTeamsController ...
func NewSearchTeamsController(store ports.Datastore) *SearchTeamsControllerImpl {
	return &SearchTeamsControllerImpl{
		teams: tables.Results[models.Team]{SearchFields: []string{"name"}},
		store: store,
	}
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
func (l *SearchTeamsControllerImpl) Prepare() error {
	var params struct {
		TeamID string `json:"team" form:"team" query:"team" validate:"required"`
	}

	err := l.BindQuery(&params)
	if err != nil {
		return err
	}
	l.teams.Search = params.TeamID

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListTeams(ctx, &l.teams)
	})
}

// Get ...
func (l *SearchTeamsControllerImpl) Get() error {
	return l.Render(
		htmx.Fragment(
			htmx.ForEach(l.teams.GetRows(), func(e *models.Team, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(e.Name),
					htmx.Text(conv.String(e.ID)),
					htmx.Input(
						htmx.Type("hidden"),
						htmx.Name("team_id"),
						htmx.Value(conv.String(e.ID)),
					),
				)
			})...,
		),
	)
}
