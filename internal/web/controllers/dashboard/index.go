package dashboard

import (
	"github.com/zeiss/typhoon/internal/web/components"

	htmx "github.com/zeiss/fiber-htmx"
)

// IndexDashboardController ...
type IndexDashboardController struct {
	htmx.DefaultController
}

// NewIndexDashboardController ...
func NewIndexDashboardController() *IndexDashboardController {
	return &IndexDashboardController{}
}

// Prepare ...
func (d *IndexDashboardController) Prepare() error {
	return nil
}

// Get ...
func (l *IndexDashboardController) Get() error {
	return l.Hx().RenderComp(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				components.Wrap(
					components.WrapProps{},
				),
				htmx.Div(
					htmx.ID("messages"),
				),
			),
		),
	)
}
