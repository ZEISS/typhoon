package operators

import (
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

	ports.Operators
	htmx.DefaultController
}

// NewListOperatorsController ...
func NewListOperatorsController(db ports.Operators) *ListOperatorsController {
	return &ListOperatorsController{
		Operators:         db,
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

	err := l.ListOperators(l.Context(), &pagination)
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
