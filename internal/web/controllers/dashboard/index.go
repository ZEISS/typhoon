package dashboard

import (
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tailwind"
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
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Dashboard",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Welcome to Typhoon"),
						),
						stats.Stats(
							stats.StatsProps{},
							stats.Stat(
								stats.StatProps{},
								stats.Title(
									stats.TitleProps{},
									htmx.Text("Accounts"),
								),
								stats.Value(
									stats.ValueProps{},
									htmx.Text("0"),
								),
							),
							stats.Stat(
								stats.StatProps{},
								stats.Title(
									stats.TitleProps{},
									htmx.Text("Teams"),
								),
								stats.Value(
									stats.ValueProps{},
									htmx.Text("0"),
								),
							),
							stats.Stat(
								stats.StatProps{},
								stats.Title(
									stats.TitleProps{},
									htmx.Text("Users"),
								),
								stats.Value(
									stats.ValueProps{},
									htmx.Text("0"),
								),
							),
						),
					),
				)
			},
		),
	)
}
