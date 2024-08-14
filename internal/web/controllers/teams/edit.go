package teams

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// TeamEditControllerImpl ...
type TeamEditControllerImpl struct {
	team  models.Team
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamEditController ...
func NewTeamEditController(store ports.Datastore) *TeamEditControllerImpl {
	return &TeamEditControllerImpl{store: store}
}

// Prepare ...
func (p *TeamEditControllerImpl) Prepare() error {
	var params struct {
		ID uuid.UUID `json:"id" form:"id" validate:"required"`
	}

	err := p.BindParams(&params)
	if err != nil {
		return err
	}
	p.team.ID = params.ID

	return p.store.ReadTx(p.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetTeam(ctx, &p.team)
	})
}

// Post ...
func (p *TeamEditControllerImpl) Post() error {
	return p.Render(
		htmx.Fallback(
			htmx.ErrorBoundary(
				func() htmx.Node {
					var body struct {
						Name        string `json:"name" form:"name" validate:"required,min=3,max=128,alphanum"`
						Description string `json:"description" form:"description" validate:"omitempty,min=3,max=1024"`
					}

					errorx.Panic(p.BindBody(&body))
					p.team.Name = body.Name
					p.team.Description = body.Description

					errorx.Panic(p.store.ReadWriteTx(p.Context(), func(ctx context.Context, tx ports.ReadWriteTx) error {
						return tx.UpdateTeam(ctx, &p.team)
					}))

					return cards.CardBordered(
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
					)
				},
			),
			func(err error) htmx.Node {
				return htmx.Text(err.Error())
			},
		),
	)
}

// Get ...
func (p *TeamEditControllerImpl) Get() error {
	return p.Render(
		htmx.FormElement(
			htmx.HxPost(fmt.Sprintf(utils.EdtiTeamsUrlFormat, p.team.ID)),
			htmx.HxTarget("this"),
			htmx.HxSwap("outerHTML"),
			cards.CardBordered(
				cards.CardProps{
					ClassNames: htmx.ClassNames{
						tailwind.M2: true,
					},
				},
				cards.Body(
					cards.BodyProps{},
					cards.Title(
						cards.TitleProps{},
						htmx.Text("Overview"),
					),
					htmx.Div(
						forms.FormControl(
							forms.FormControlProps{},
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{},
									htmx.Text("Name"),
								),
							),
							forms.TextInputBordered(
								forms.TextInputProps{
									Name:        "name",
									Placeholder: "Ghostbusters, Avengers, Masters of the Universe ...",
									Value:       p.team.Name,
								},
							),
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{
										ClassNames: htmx.ClassNames{
											"text-neutral-500": true,
										},
									},
									htmx.Text("The name must be from 3 to 128 characters. At least 3 characters must be non-whitespace, only alphanumeric characters are allowed."),
								),
							),
						),
						forms.FormControl(
							forms.FormControlProps{},
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{},
									htmx.Text("Description"),
								),
							),
							forms.TextareaBordered(
								forms.TextareaProps{
									Name:        "description",
									Placeholder: "A team of superheroes that save the world from evil ...",
								},
								htmx.Text(p.team.Description),
							),
							forms.FormControlLabel(
								forms.FormControlLabelProps{},
								forms.FormControlLabelText(
									forms.FormControlLabelTextProps{
										ClassNames: htmx.ClassNames{
											"text-neutral-500": true,
										},
									},
									htmx.Text("This is optional. The description must be from 3 to 1024 characters."),
								),
							),
						),
					),
					cards.Actions(
						cards.ActionsProps{},
						buttons.Button(
							buttons.ButtonProps{
								Type: "button",
							},
							htmx.Text("Save"),
							htmx.HxPost(fmt.Sprintf(utils.EdtiTeamsUrlFormat, p.team.ID)),
							htmx.HxSwap("outerHTML"),
						),
					),
				),
			),
		),
	)
}
