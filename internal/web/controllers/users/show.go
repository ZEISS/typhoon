package users

import (
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// ShowUserControllerImpl ...
type ShowUserControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`

	ports.Users
	htmx.DefaultController
}

// NewShowUserController ...
func NewShowUserController(db ports.Users) *ShowUserControllerImpl {
	return &ShowUserControllerImpl{
		Users:             db,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *ShowUserControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	user := models.User{ID: l.ID}

	err = l.GetUser(l.Context(), &user)
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
								htmx.Text(user.Name),
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
									htmx.Text("Description"),
								),
								htmx.H3(
									htmx.Text(user.Description),
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
									htmx.Text("Created at"),
								),
								htmx.H3(
									htmx.Text(
										user.CreatedAt.Format("2006-01-02 15:04:05"),
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
										user.UpdatedAt.Format("2006-01-02 15:04:05"),
									),
								),
							),
						),
						cards.Actions(
							cards.ActionsProps{},
							dropdowns.Dropdown(
								dropdowns.DropdownProps{
									ClassNames: htmx.ClassNames{
										"dropdown-end": true,
									},
								},
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									icons.BoltOutline(
										icons.IconProps{},
									),
									htmx.Text("Actions"),
								),
								dropdowns.DropdownMenuItems(
									dropdowns.DropdownMenuItemsProps{},
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										htmx.A(
											htmx.Href(fmt.Sprintf("/users/%s/credentials", user.ID)),
											htmx.Text("Get Credentials"),
										),
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