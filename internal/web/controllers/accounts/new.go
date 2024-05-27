package accounts

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// NewAccountControllerImpl ...
type NewAccountControllerImpl struct {
	Operators []*models.Operator

	ports.Repository
	htmx.DefaultController
}

// NewAccountController ...
func NewAccountController(db ports.Repository) *NewAccountControllerImpl {
	return &NewAccountControllerImpl{
		Repository:        db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *NewAccountControllerImpl) Prepare() error {
	pagination := models.Pagination[models.Operator]{}

	err := l.BindQuery(&pagination)
	if err != nil {
		return err
	}

	err = l.ListOperators(l.Context(), &pagination)
	if err != nil {
		return err
	}
	for _, operator := range pagination.Rows {
		l.Operators = append(l.Operators, &operator)
	}

	fmt.Println((len(l.Operators)))

	return nil
}

// Get ...
func (l *NewAccountControllerImpl) Get() error {
	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost("/accounts/create"),
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
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"-my-4": true,
											},
										},
										htmx.Text("Operator	"),
									),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										forms.SelectBordered(
											forms.SelectProps{},
											htmx.HxGet("/accounts/partials/operator-skgs"),
											htmx.HxTarget("#operator-skgs"),
											htmx.HxSwap("outerHTML"),
											forms.Option(
												forms.OptionProps{
													Selected: true,
													Disabled: true,
												},
												htmx.Text("Select an operator"),
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
													"-my-4": true,
												},
											},
											htmx.Text("Signing Key Group"),
										),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											forms.SelectBordered(
												forms.SelectProps{},
												forms.Option(
													forms.OptionProps{
														Selected: true,
														Disabled: true,
													},
													htmx.Text("Select an signing key group"),
												),
												htmx.ID("operator-skgs"),
												htmx.Name("operator_skgs_id"),
											),
										),
									),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("A unique identifier for the account."),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name: "name",
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
													"-my-4": true,
												},
											},
											htmx.Text("Description"),
										),
									),
									forms.FormControlLabel(
										forms.FormControlLabelProps{},
										forms.FormControlLabelText(
											forms.FormControlLabelTextProps{
												ClassNames: htmx.ClassNames{
													"text-neutral-500": true,
												},
											},
											htmx.Text("A brief description of the acount to provide context."),
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
