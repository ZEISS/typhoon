package operators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var validate *validator.Validate

// NewOperatorControllerImpl ...
type NewOperatorControllerImpl struct {
	ports.Operators
	htmx.DefaultController
}

// NewOperatorsController ...
func NewOperatorController(db ports.Operators) *NewOperatorControllerImpl {
	return &NewOperatorControllerImpl{db, htmx.DefaultController{}}
}

// Prepare ...
func (l *NewOperatorControllerImpl) Prepare() error {
	validate = validator.New()

	return nil
}

// Post ...
func (l *NewOperatorControllerImpl) Post() error {
	query := struct {
		Name        string `json:"name" form:"name" validate:"required,min=3,max=100"`
		Description string `json:"description" form:"description" validate:"required,min=3,max=1024"`
	}{}

	err := l.Ctx().BodyParser(&query)
	if err != nil {
		return err
	}

	err = validate.Struct(query)
	if err != nil {
		return err
	}

	op, err := models.NewOperator(query.Name, query.Description)
	if err != nil {
		return err
	}

	err = l.CreateOperator(l.Context(), &op)
	if err != nil {
		return err
	}

	htmx.HxRedirect(l.Ctx(), "/operators")

	return nil
}

// Error ...
func (l *NewOperatorControllerImpl) Error(err error) error {
	fmt.Println(err)

	return nil
}

// Get ...
func (l *NewOperatorControllerImpl) Get() error {
	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost("/operators/new"),
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
										htmx.Text("Name"),
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
										htmx.Text("A unique identifier for operator."),
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
											htmx.Text("A brief description of the operator to provide context."),
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
									htmx.Text("Create Operator"),
								),
							),
						),
					),
				),
			),
		),
	)
}
