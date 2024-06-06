package systems

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewSystemControllerImpl ...
type NewSystemControllerImpl struct {
	Operators []*models.Operator

	store ports.Datastore
	htmx.DefaultController
}

// NewSystemController ...
func NewSystemController(store ports.Datastore) *NewSystemControllerImpl {
	return &NewSystemControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *NewSystemControllerImpl) Prepare() error {
	pagination := models.Pagination[models.Operator]{}

	err := l.BindQuery(&pagination)
	if err != nil {
		return err
	}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListOperators(ctx, &pagination)
	})
	if err != nil {
		return err
	}

	for _, op := range pagination.Rows {
		l.Operators = append(l.Operators, &op)
	}

	return nil
}

// Get ...
func (l *NewSystemControllerImpl) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(),
				},
				htmx.FormElement(
					htmx.HxPost("/systems"),
					cards.CardBordered(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
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
										htmx.ForEach(l.Operators, func(operator *models.Operator) htmx.Node {
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
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
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
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
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
									htmx.Attribute("type", "submit"),
									htmx.Text("Create System"),
								),
							),
						),
					),
				),
			),
		),
	)
}
