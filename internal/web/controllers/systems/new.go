package systems

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewSystemControllerImpl ...
type NewSystemControllerImpl struct {
	Results tables.Results[models.Operator]

	store ports.Datastore
	htmx.DefaultController
}

// NewSystemController ...
func NewSystemController(store ports.Datastore) *NewSystemControllerImpl {
	return &NewSystemControllerImpl{
		Results:           tables.Results[models.Operator]{},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
}

// Prepare ...
func (l *NewSystemControllerImpl) Prepare() error {
	err := l.BindQuery(&l.Results)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &l.Results)
	})
}

// Get ...
func (l *NewSystemControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "New System",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				return htmx.FormElement(
					htmx.HxPost("/systems"),
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
								forms.FormControlProps{},
								forms.SelectBordered(
									forms.SelectProps{},
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select operator"),
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
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelAltText(
										forms.FormControlLabelAltTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("An operator needs to be created before adding a system."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{
										ClassNames: htmx.ClassNames{
											"flex":        true,
											"flex-col":    true,
											"items-start": true,
										},
									},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full":           true,
												"text-neutral-500": true,
											},
										},
										htmx.Text("A unique identifier for the system."),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name: "name",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelAltText(
										forms.FormControlLabelAltTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("The name must be from 3 to 100 characters. At least 3 characters must be non-whitespace."),
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
												"text-neutral-500": true,
											},
										},
										htmx.Text("A brief description of the system to provide context."),
									),
								),
								forms.TextareaBordered(
									forms.TextareaProps{
										Name: "description",
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
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Create System"),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"my-2": true,
								"mx-2": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Tags - Optional"),
							),
							htmx.Div(
								alpine.XData(`{
						tags: [],
						addTag(tag) {
						  this.tags.push({name: '', value: ''});
						},
						removeTag(index) {
						  this.tags.splice(index, 1);
						}
					  }`),
								htmx.Template(
									alpine.XFor("(tag, index) in tags"),
									htmx.Attribute(":key", "index"),
									htmx.Div(
										htmx.ClassNames{
											tailwind.Flex:    true,
											tailwind.SpaceX4: true,
										},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{},
											},
											forms.TextInputBordered(
												forms.TextInputProps{},
												alpine.XModel("tag.name"),
												alpine.XBind("name", "`tags.${index}.name`"),
											),
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{
														ClassNames: htmx.ClassNames{
															"text-neutral-500": true,
														},
													},
													htmx.Text("Key is a unique identifier. At least 3 characters must be non-whitespace."),
												),
											),
										),
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{},
											},
											forms.TextInputBordered(
												forms.TextInputProps{},
												alpine.XModel("tag.value"),
												alpine.XBind("name", "`tags.${index}.value`"),
											),
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{
														ClassNames: htmx.ClassNames{
															"text-neutral-500": true,
														},
													},
													htmx.Text("Value is a unique value for the key."),
												),
											),
										),
										buttons.Button(
											buttons.ButtonProps{
												Type: "button",
											},
											alpine.XOn("click", "removeTag(index)"),
											icons.TrashOutline(
												icons.IconProps{},
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
										alpine.XOn("click", "addTag()"),
										htmx.Text("Add Tag"),
									),
								),
							),
						),
					),
				)
			},
		),
	)
}
