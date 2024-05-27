package users

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var validate *validator.Validate

// CreateUserControllerImpl ...
type CreateUserControllerImpl struct {
	AccountID                uuid.UUID `json:"account_id" form:"account_id" validate:"required,uuid"`
	AccountSigningKeyGroupID uuid.UUID `json:"account_skgs_id" form:"account_skgs_id" validate:"required,uuid"`
	Name                     string    `json:"name" form:"name" validate:"required,min=3,max=100"`
	Description              string    `json:"description" form:"description" validate:"required,min=3,max=1024"`

	ports.Repository
	htmx.DefaultController
}

// NewCreateUserController ...
func NewCreateUserController(db ports.Repository) *CreateUserControllerImpl {
	return &CreateUserControllerImpl{
		Repository:        db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *CreateUserControllerImpl) Prepare() error {
	validate = validator.New()

	err := l.Ctx().BodyParser(l)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}

// Post ...
func (l *CreateUserControllerImpl) Post() error {
	user := models.User{Name: l.Name, Description: l.Description}
	account := models.Account{ID: l.AccountID, SigningKeyGroups: []models.SigningKeyGroup{{ID: l.AccountSigningKeyGroupID}}}

	err := l.GetAccount(l.Context(), &account)
	if err != nil {
		return err
	}
	user.Account = account

	pk, err := nkeys.CreateUser()
	if err != nil {
		return err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return err
	}

	seed, err := pk.Seed()
	if err != nil {
		return err
	}
	user.Key = models.NKey{ID: id, Seed: seed}

	if len(account.SigningKeyGroups) < 1 {
		return fmt.Errorf("account %s has no signing keys", account.ID)
	}

	askg := account.FindSigningKeyGroupByID(l.AccountSigningKeyGroupID)
	if askg == nil {
		return fmt.Errorf("account %s does not have signing key group %s", account.ID, l.AccountSigningKeyGroupID)
	}

	ask, err := nkeys.FromSeed(askg.Key.Seed)
	if err != nil {
		return err
	}

	askpk, err := ask.PublicKey()
	if err != nil {
		return err
	}

	// // Create a token for the user
	u := jwt.NewUserClaims(id)
	u.Name = l.Name
	u.IssuerAccount = account.Key.ID
	u.Issuer = askpk

	token, err := u.Encode(ask)
	if err != nil {
		return err
	}
	user.Token = models.Token{ID: id, Token: token}

	err = l.CreateUser(l.Context(), &user)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/users")

	return nil
}

// Get ...
func (l *CreateUserControllerImpl) Get() error {
	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				htmx.FormElement(
					htmx.HxPost("/users/new"),
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
