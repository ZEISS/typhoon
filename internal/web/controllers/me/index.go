package me

import (
	"context"

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

	store ports.Datastore
	htmx.DefaultController
}

// NewMeIndexController ...
func NewMeController(store ports.Datastore) *MeController {
	return &MeController{
		User:              adapters.GothUser{},
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (m *MeController) Prepare() error {
	return m.store.ReadTx(m.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetProfile(ctx, &m.User)
	})
}

// Get ...
func (m *MeController) Get() error {
	return m.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Profile",
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"m-2": true,
						},
					},
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
				)
			},
		),
	)
}
