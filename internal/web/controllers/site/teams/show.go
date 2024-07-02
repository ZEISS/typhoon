package teams

import (
	"context"
	"fmt"

	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/teams"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// TeamShowControllerImpl ...
type TeamShowControllerImpl struct {
	team  tables.Paginated[adapters.GothTeam]
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
	err := p.BindParams(&p.team.Value)
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
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: p.Path(),
					User: p.Session().User,
				},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
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
										htmx.Text("ID"),
									),
									htmx.H3(
										htmx.Text(p.team.Value.ID.String()),
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
										htmx.Text("Name"),
									),
									htmx.H3(
										htmx.Text(p.team.Value.Name),
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
										htmx.Text(p.team.Value.Description),
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
										htmx.Text("Created at"),
									),
									htmx.H3(
										htmx.Text(
											p.team.Value.CreatedAt.Format("2006-01-02 15:04:05"),
										),
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
										htmx.Text("Updated at"),
									),
									htmx.H3(
										htmx.Text(
											p.team.Value.UpdatedAt.Format("2006-01-02 15:04:05"),
										),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								links.Button(
									links.LinkProps{
										ClassNames: htmx.ClassNames{
											"btn-outline": true,
										},
										Href: fmt.Sprintf("%s/edit", p.team.Value.ID),
									},
									htmx.Text("Edit"),
								),
								buttons.Outline(
									buttons.ButtonProps{},
									htmx.HxDelete(""),
									htmx.HxConfirm("Are you sure you want to delete this team?"),
									htmx.Text("Delete"),
								),
							),
						),
					),
				),
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						teams.UsersTable(
							teams.UsersTableProps{
								Users:  tables.RowsPtr(p.team.Value.Users),
								Offset: p.team.GetOffset(),
								Limit:  p.team.GetLimit(),
								// Total:  p.team.GetLen(),
							},
						),
					),
				),
			),
		),
	)
}
