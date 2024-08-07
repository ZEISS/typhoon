package operators

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewOperatorControllerImpl ...
type NewOperatorControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewOperatorsController ...
func NewOperatorController(store ports.Datastore) *NewOperatorControllerImpl {
	return &NewOperatorControllerImpl{store: store}
}

// Get ...
func (l *NewOperatorControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "New Operator",
				Path:  l.Path(),
			},
			func() htmx.Node {
				return htmx.FormElement(
					htmx.HxPost("/operators/new"),
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
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
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
										Placeholder: "Indiana Jones, Luke Skywalker, etc.",
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
										htmx.Text("The name must be from 3 to 100 characters. At least 3 characters must be non-whitespace."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Description"),
									),
								),
								forms.TextareaBordered(
									forms.TextareaProps{
										Name:        "description",
										Placeholder: "In a galaxy far, far away...",
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
										htmx.Text("The description must be from 3 to 1024 characters."),
									),
								),
							),
						),
					),
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
								htmx.Text("Account Server"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:        "account_server_url",
										Placeholder: "https://example.com:8080",
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
										htmx.Text("A valid URL with a scheme of http or https. Certificates need be valid."),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Outline(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Create Operator"),
								),
							),
						),
					),
				)
			},
		),
	)
}
