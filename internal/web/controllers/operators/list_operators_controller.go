package operators

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
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
	ops := htmx.Values[models.Pagination[models.Operator]](l.Ctx().UserContext(), resolvers.ValuesKeyOperators)

	fmt.Println(ops)

	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				htmx.Text("Operators"),
			),
		),
	)
}
