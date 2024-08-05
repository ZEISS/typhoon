package dashboard

import (
	"github.com/zeiss/fiber-htmx/components/cards"
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

// Get ...
func (l *IndexDashboardController) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{
				Title: "Dashboard",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Welcome to Typhoon"),
						),
					),
				),
			),
		),
	)
}
