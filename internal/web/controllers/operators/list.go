package operators

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/operators"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&ListOperatorsController{})

// ListOperatorsController ...
type ListOperatorsController struct {
	Pagination models.Pagination[models.Operator]

	store ports.Datastore
	htmx.DefaultController
}

// NewListOperatorsController ...
func NewListOperatorsController(store ports.Datastore) *ListOperatorsController {
	return &ListOperatorsController{
		Pagination:        models.Pagination[models.Operator]{Limit: 10},
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *ListOperatorsController) Prepare() error {
	err := l.Ctx().QueryParser(&l.Pagination)
	if err != nil {
		return nil
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &l.Pagination)
	})
}

// Prepare ...
func (l *ListOperatorsController) Get() error {
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
						operators.OperatorsTable(
							operators.OperatorsTableProps{
								Operators: l.Pagination.GetRows(),
								Offset:    l.Pagination.GetOffset(),
								Limit:     l.Pagination.GetLimit(),
								Total:     l.Pagination.GetTotalRows(),
							},
						),
					),
				),
			),
		),
	)
}
