package operators

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/operators"
	"github.com/zeiss/typhoon/internal/web/ports"
	"github.com/zeiss/typhoon/pkg/resolvers"
)

var _ = htmx.Controller(&ListOperatorsController{})

// ListOperatorsController ...
type ListOperatorsController struct {
	htmx.DefaultController
}

// NewListOperatorsController ...
func NewListOperatorsController(db ports.Operators) *ListOperatorsController {
	return &ListOperatorsController{}
}

// Prepare ...
func (l *ListOperatorsController) Get() error {
	pagination := htmx.Values[models.Pagination[models.Operator]](l.Ctx().UserContext(), resolvers.ValuesKeyOperators)

	ops := make([]*models.Operator, 0, len(pagination.Rows))
	for _, row := range pagination.Rows {
		ops = append(ops, &row)
	}

	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
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
	)
}
