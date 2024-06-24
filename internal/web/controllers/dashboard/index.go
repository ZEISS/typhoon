package dashboard

import (
	"github.com/zeiss/typhoon/internal/web/components"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/stats"
)

// IndexDashboardController ...
type IndexDashboardController struct {
	htmx.DefaultController
}

// NewIndexDashboardController ...
func NewIndexDashboardController() *IndexDashboardController {
	return &IndexDashboardController{}
}

// Get ...
func (l *IndexDashboardController) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{
				Title: "Dashboard",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				components.Wrap(
					components.WrapProps{},
					stats.Stats(
						stats.StatsProps{},
						stats.Title(
							stats.TitleProps{},
							htmx.Text("Dashboard"),
						),
					),
				),
			),
		),
	)
}
