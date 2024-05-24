package accounts

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// ShowAccountControllerImpl ...
type ShowAccountControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`

	ports.Accounts
	htmx.DefaultController
}

// NewShowAccountController ...
func NewShowAccountController(db ports.Accounts) *ShowAccountControllerImpl {
	return &ShowAccountControllerImpl{
		Accounts:          db,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *ShowAccountControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	acc := models.Account{ID: l.ID}

	err = l.GetAccount(l.Context(), &acc)
	if err != nil {
		return err
	}

	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{},
			components.Layout(
				components.LayoutProps{},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Overview"),
						),
						htmx.Div(
							htmx.H1(
								htmx.Text(acc.Name),
							),
							htmx.P(
								htmx.Text(utils.PtrStr(acc.Description)),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Created at"),
								),
								htmx.H3(
									htmx.Text(
										acc.CreatedAt.Format("2006-01-02 15:04:05"),
									),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"flex":     true,
									"flex-col": true,
									"py-2":     true,
								},
								htmx.H4(
									htmx.ClassNames{
										"text-gray-500": true,
									},
									htmx.Text("Updated at"),
								),
								htmx.H3(
									htmx.Text(
										acc.UpdatedAt.Format("2006-01-02 15:04:05"),
									),
								),
							),
						),
						cards.Actions(
							cards.ActionsProps{},
							buttons.Outline(
								buttons.ButtonProps{},
								htmx.HxDelete(fmt.Sprintf("/account/%s", acc.ID)),
								htmx.HxConfirm("Are you sure you want to delete this lens?"),
								htmx.Text("Delete"),
							),
						),
					),
				),
			),
		),
	)

}
