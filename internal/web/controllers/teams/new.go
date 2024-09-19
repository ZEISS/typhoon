package teams

import (
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// NewTeamControllerImpl ...
type NewTeamControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewTeamController ...
func NewTeamController(store ports.Datastore) *NewTeamControllerImpl {
	return &NewTeamControllerImpl{store: store}
}

// New ...
func (p *NewTeamControllerImpl) Get() error {
	return p.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: p.Path(),
				},
				htmx.FormElement(
					htmx.HxPost(""),
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
								htmx.Text("New Team"),
							),
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
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Create Team"),
								),
							),
						),
					),
				),
			),
		),
	)
}
