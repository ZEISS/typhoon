package accounts

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
)

// NewAccountControllerImpl ...
type NewAccountControllerImpl struct {
	Results tables.Results[models.Operator]
	Teams   tables.Results[models.Team]
	store   ports.Datastore
	htmx.DefaultController
}

// NewAccountController ...
func NewAccountController(store ports.Datastore) *NewAccountControllerImpl {
	return &NewAccountControllerImpl{store: store}
}

// Prepare ...
func (l *NewAccountControllerImpl) Prepare() error {
	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		err := tx.ListOperators(ctx, &l.Results)
		if err != nil {
			return err
		}

		return tx.ListTeams(ctx, &l.Teams)
	})
}

// Get ...
func (l *NewAccountControllerImpl) Get() error {
	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{
				Title: "New Account",
				Path:  l.Path(),
			},
			components.Layout(
				components.LayoutProps{},
				htmx.FormElement(
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Team"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									htmx.Required(),
									htmx.HxValidate(true),
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select a team"),
									),
									htmx.Name("team_id"),
									htmx.Group(
										htmx.ForEach(l.Teams.GetRows(), func(operator *models.Team, idx int) htmx.Node {
											return forms.Option(
												forms.OptionProps{
													Value: operator.ID.String(),
												},
												htmx.Text(operator.Name),
											)
										})...,
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
										htmx.Text("Operator"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									htmx.HxGet("/accounts/partials/operator-skgs"),
									htmx.HxTarget("#operator-skgs"),
									htmx.HxSwap("outerHTML"),
									htmx.HxValidate(true),
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select an operator"),
									),
									htmx.Name("operator_id"),
									htmx.Group(
										htmx.ForEach(l.Results.GetRows(), func(operator *models.Operator, idx int) htmx.Node {
											return forms.Option(
												forms.OptionProps{
													Value: operator.ID.String(),
												},
												htmx.Text(operator.Name),
											)
										})...,
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Signing Key Group"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select an signing key group"),
									),
									htmx.HxValidate(true),
									htmx.ID("operator-skgs"),
									htmx.Name("operator_skgs_id"),
								),
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
										Placeholder: "Start giving it a name ...",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												tailwind.TextNeutral500: true,
											},
										},
										htmx.Text("The name must be from 3 to 255 characters. At least 3 characters must be non-whitespace."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												tailwind.TextNeutral500: true,
											},
										},
										htmx.Text("A brief description of the acount to provide context."),
									),
								),
								forms.TextareaBordered(
									forms.TextareaProps{
										Name:        "description",
										Placeholder: "Start typing a description ...",
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
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Limits"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Enable JetStream"),
									),
								),
								forms.Toggle(
									forms.ToggleProps{
										Name:    "jetstream_enable",
										Checked: true,
										Value:   conv.String(true),
									},
									htmx.ID("jetstream_enable"),
								),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Max Disk Storage"),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex": true,
									},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:  "jetstream_max_disk_storage",
											Value: "2.5",
										},
									),
									forms.SelectBordered(
										forms.SelectProps{
											ClassNames: htmx.ClassNames{
												"w-full": false,
											},
										},
										htmx.Name("jetstream_max_disk_storage_unit"),
										forms.Option(
											forms.OptionProps{
												Value: "B",
											},
											htmx.Text("Bytes"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "KB",
											},
											htmx.Text("KiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "MB",
											},
											htmx.Text("MiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value:    "GB",
												Selected: true,
											},
											htmx.Text("GB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "TB",
											},
											htmx.Text("TiB"),
										),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Streams"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:  "jetstream_max_streams",
										Value: "10",
									},
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Consumers"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:  "jetstream_max_consumers",
										Value: "10",
									},
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Max Stream Size"),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex": true,
									},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:  "jetstream_max_stream_size",
											Value: "2.6",
										},
									),
									forms.SelectBordered(
										forms.SelectProps{
											ClassNames: htmx.ClassNames{
												"w-full": false,
											},
										},
										htmx.Name("jetstream_max_stream_size_unit"),
										forms.Option(
											forms.OptionProps{
												Value: "bytes",
											},
											htmx.Text("Bytes"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "kilobit",
											},
											htmx.Text("KiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "megabit",
											},
											htmx.Text("MiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value:    "gigabit",
												Selected: true,
											},
											htmx.Text("GiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "terabit",
											},
											htmx.Text("TiB"),
										),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Require Max Bytes"),
									),
								),
								forms.Toggle(
									forms.ToggleProps{
										Checked: true,
									},
									htmx.Name("jetstream_max_bytes_required"),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Tags - Optional"),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Outline(
									buttons.ButtonProps{},
									htmx.Action("."),
									htmx.HxPost("/accounts/create"),
									htmx.HxTargetError("#alerts"),
									htmx.Attribute("type", "submit"),
									htmx.Text("Create Account"),
								),
							),
						),
					),
				),
			),
		),
	)
}
