package teams

import (
	"context"
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/teams"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// TeamShowControllerImpl ...
type TeamShowControllerImpl struct {
	team  models.Team
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamShowController ...
func NewTeamShowController(store ports.Datastore) *TeamShowControllerImpl {
	return &TeamShowControllerImpl{
		store: store,
	}
}

// Prepare ...
func (p *TeamShowControllerImpl) Prepare() error {
	err := p.BindParams(&p.team)
	if err != nil {
		return err
	}

	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &p.team)
	})
}

// Get ...
func (p *TeamShowControllerImpl) Get() error {
	return p.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Team",
				Path:  p.Path(),
			},
			func() htmx.Node {
				return htmx.Fragment(
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						htmx.HxTarget("this"),
						htmx.HxSwap("outerHTML"),
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Overview"),
							),
							htmx.Div(
								htmx.Div(
									htmx.ClassNames{
										"flex":     true,
										"flex-col": true,
										"py-2":     true,
									},
									htmx.H4(
										htmx.ClassNames{
											"text-gray-500": true,
										},
										htmx.Text("Name"),
									),
									htmx.H3(
										htmx.Text(p.team.Name),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex":     true,
										"flex-col": true,
										"py-2":     true,
									},
									htmx.H4(
										htmx.ClassNames{
											"text-gray-500": true,
										},
										htmx.Text("Description"),
									),
									htmx.H3(
										htmx.Text(p.team.Description),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{
										Type: "button",
									},
									htmx.Text("Edit"),
									htmx.HxGet(fmt.Sprintf(utils.EdtiTeamsUrlFormat, p.team.ID)),
									htmx.HxSwap("outerHTML"),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this team?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
					teams.MetaCard(
						teams.MetaCardProps{
							Team: p.team,
						},
					),
				)
			},
		),
	)
}
