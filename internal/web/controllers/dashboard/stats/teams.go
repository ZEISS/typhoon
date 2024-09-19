package stats

import (
	"context"

	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/errorx"
)

// TeamsStatsControllerImpl ...
type TeamsStatsControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamsStatsController ...
func NewTeamsStatsController(store ports.Datastore) *TeamsStatsControllerImpl {
	return &TeamsStatsControllerImpl{store: store}
}

// Get ...
func (d *TeamsStatsControllerImpl) Get() error {
	return d.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := d.store.ReadTx(d.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfTeams(ctx, &total)
				})
				errorx.Panic(err)

				return stats.Value(
					stats.ValueProps{},
					htmx.Text(conv.String(total)),
				)
			}),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}
