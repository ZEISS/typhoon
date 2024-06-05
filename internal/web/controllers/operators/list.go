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
	Limit  int    `json:"limit" xml:"limit" form:"limit"`
	Offset int    `json:"offset" xml:"offset" form:"offset"`
	Search string `json:"search" xml:"search" form:"search"`

	store ports.Datastore
	htmx.DefaultController
}

// NewListOperatorsController ...
func NewListOperatorsController(store ports.Datastore) *ListOperatorsController {
	return &ListOperatorsController{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *ListOperatorsController) Prepare() error {
	err := l.Ctx().QueryParser(&l)
	if err != nil {
		return nil
	}

	return nil
}

// Prepare ...
func (l *ListOperatorsController) Get() error {
	pagination := models.Pagination[models.Operator]{}

	pagination.Limit = l.Limit
	pagination.Offset = l.Offset
	pagination.Search = l.Search

	err := l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &pagination)
	})
	if err != nil {
		return err
	}

	ops := make([]*models.Operator, 0, len(pagination.Rows))
	for _, row := range pagination.Rows {
		ops = append(ops, &row)
	}

	return htmx.RenderComp(
		l.Ctx(),
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
								Operators: ops,
								Offset:    pagination.Offset,
								Limit:     pagination.Limit,
								Total:     pagination.TotalRows,
							},
						),
					),
				),
			),
		),
	)
}
