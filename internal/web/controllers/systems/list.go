package systems

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/systems"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListSystemsController{})

// ListSystemsController ...
type ListSystemsController struct {
	systems tables.Results[models.System]
	store   ports.Datastore
	htmx.DefaultController
}

// NewListSystemsController ...
func NewListSystemsController(store ports.Datastore) *ListSystemsController {
	return &ListSystemsController{
		systems: tables.Results[models.System]{SearchFields: []string{"name"}},
		store:   store,
	}
}

// Prepare ...
func (l *ListSystemsController) Prepare() error {
	err := l.BindQuery(&l.systems)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListSystems(ctx, &l.systems)
	})
}

// Get ...
func (l *ListSystemsController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Systems",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						systems.SystemsTable(
							systems.SystemsTableProps{
								Limit:   l.systems.GetLimit(),
								Offset:  l.systems.GetOffset(),
								Total:   l.systems.GetLen(),
								Systems: l.systems.GetRows(),
							},
						),
					),
				)
			},
		),
	)
}
