package systems

import (
	"github.com/zeiss/typhoon/internal/web/components"
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
	Params ListSystemsControllerParams

	ports.Users
	htmx.DefaultController
}

// NewListSystemsController ...
func NewListSystemsController(db ports.Users) *ListSystemsController {
	return &ListSystemsController{
		Params:            ListSystemsControllerParams{},
		Users:             db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *ListSystemsController) Prepare() error {
	err := l.BindParams(&l.Params)
	if err != nil {
		return err
	}

	return nil
}

// Prepare ...
func (l *ListSystemsController) Get() error {
	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{
				Title: "Systems",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
					),
				),
			),
		),
	)
}
