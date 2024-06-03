package me

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// MeController ...
type MeController struct {
	User adapters.GothUser
	ports.Me
	htmx.DefaultController
}

// NewMeIndexController ...
func NewMeController(db ports.Me) *MeController {
	return &MeController{
		User:              adapters.GothUser{},
		Me:                db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (m *MeController) Prepare() error {
	return m.GetProfile(m.Context(), &m.User)
}

// Get ...
func (m *MeController) Get() error {
	return htmx.RenderComp(
		m.Ctx(),
		components.Page(
			components.PageProps{
				Title: "Profile",
			},
			components.Layout(
				components.LayoutProps{
					Path: m.Ctx().Path(),
				},
				components.Wrap(
					components.WrapProps{},
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Profile"),
							),
							htmx.Form(
								htmx.HxPost("/me"),
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
											Name:     "username",
											Value:    m.User.Name,
											Disabled: true,
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
											htmx.Text("Your full nane as it will appear in the system."),
										),
									),
								),
								forms.FormControl(
									forms.FormControlProps{},
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{},
											htmx.Text("Email"),
										),
									),
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:     "email",
											Value:    m.User.Email,
											Disabled: true,
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
											htmx.Text("Your email address. This is where we will send notifications."),
										),
									),
								),

								cards.Actions(
									cards.ActionsProps{},
									buttons.OutlinePrimary(
										buttons.ButtonProps{
											Disabled: true,
										},
										htmx.Attribute("type", "submit"),
										htmx.Text("Update Profile"),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}
