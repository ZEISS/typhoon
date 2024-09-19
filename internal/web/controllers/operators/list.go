package operators

import (
	"context"

	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/operators"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/tables"
)

var _ = htmx.Controller(&ListOperatorsController{})

// ListOperatorsController ...
type ListOperatorsController struct {
	operators tables.Results[models.Operator]
	store     ports.Datastore
	htmx.DefaultController
}

// NewListOperatorsController ...
func NewListOperatorsController(store ports.Datastore) *ListOperatorsController {
	return &ListOperatorsController{
		operators: tables.Results[models.Operator]{Limit: 10, SearchFields: []string{"name"}},
		store:     store,
	}
}

// Prepare ...
func (l *ListOperatorsController) Prepare() error {
	err := l.Ctx().QueryParser(&l.operators)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &l.operators)
	})
}

// Prepare ...
func (l *ListOperatorsController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Operators",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				return htmx.Fragment(
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							operators.OperatorsTable(
								operators.OperatorsTableProps{
									Operators: l.operators.GetRows(),
									Offset:    l.operators.GetOffset(),
									Limit:     l.operators.GetLimit(),
									Total:     l.operators.GetLen(),
									URL:       l.OriginalURL(),
								},
							),
						),
					),
				)
			},
		),
	)
}
