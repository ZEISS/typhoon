package stats

import (
	"context"

	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/pkg/conv"
	"github.com/zeiss/pkg/errorx"
)

// SystemsStatsControllerImpl ...
type SystemsStatsControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewSystemsStatsController ...
func NewSystemsStatsController(store ports.Datastore) *SystemsStatsControllerImpl {
	return &SystemsStatsControllerImpl{store: store}
}

// Get ...
func (d *SystemsStatsControllerImpl) Get() error {
	return d.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(func() htmx.Node {
				var total int64

				err := d.store.ReadTx(d.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetTotalNumberOfSystems(ctx, &total)
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
