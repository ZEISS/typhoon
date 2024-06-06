package systems

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/systems"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
)

var _ = htmx.Controller(&ListSystemsController{})

// ListSystemsControllerParams ...
type ListSystemsControllerParams struct {
	Offset int    `json:"offset" form:"offset" params:"offset"`
	Limit  int    `json:"limit" form:"limit" params:"limit"`
	Search string `json:"search" form:"search" params:"search"`
	Sort   string `json:"sort" form:"sort" params:"sort"`
}

// ListSystemsController ...
type ListSystemsController struct {
	Pagination models.Pagination[models.System]

	store ports.Datastore
	htmx.DefaultController
}

// NewListSystemsController ...
func NewListSystemsController(store ports.Datastore) *ListSystemsController {
	return &ListSystemsController{
		Pagination:        models.Pagination[models.System]{Limit: 10},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
}

// Prepare ...
func (l *ListSystemsController) Prepare() error {
	err := l.BindQuery(&l.Pagination)
	if err != nil {
		return err
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListSystems(ctx, &l.Pagination)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (l *ListSystemsController) Get() error {
	return l.Render( // render the html using htmx
		components.Page(
			components.PageProps{
				Title: "Systems",
				Boost: true,
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(), // get the current path
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						systems.SystemsTable(
							systems.SystemsTableProps{
								Limit:   l.Pagination.GetLimit(),
								Offset:  l.Pagination.GetOffset(),
								Total:   l.Pagination.GetTotalRows(),
								Systems: l.Pagination.GetRows(),
							},
						),
					),
				),
			),
		),
	)
}
