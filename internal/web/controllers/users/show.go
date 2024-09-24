package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tooltips"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// ShowUserControllerImpl ...
type ShowUserControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewShowUserController ...
func NewShowUserController(store ports.Datastore) *ShowUserControllerImpl {
	return &ShowUserControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *ShowUserControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Users",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				err := l.BindParams(l)
				if err != nil {
					panic(err)
				}

				user := models.User{ID: l.ID}
				err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetUser(ctx, &user)
				})
				if err != nil {
					panic(err)
				}

				return htmx.Fragment(
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"m-2": true,
							},
						},
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
								htmx.A(
									htmx.ClassNames{
										"btn": true,
									},
									htmx.Href(fmt.Sprintf(utils.DownloadCredentialsUserUrlFormat, user.ID)),
									icons.ArrowDownOnSquareOutline(
										icons.IconProps{},
									),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(fmt.Sprintf(utils.DeleteUserUrlFormat, user.ID)),
									htmx.HxConfirm("Are you sure you want to delete this operator?"),
									icons.TrashOutline(
										icons.IconProps{},
									),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"m-2": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Details"),
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
									htmx.Text("Public NKey"),
									tooltips.Tooltip(
										tooltips.TooltipProps{
											ClassNames: htmx.ClassNames{
												"tooltip-right": true,
											},
											DataTip: "Public NKey is the public key of the user",
										},
										icons.InformationCircleOutline(
											icons.IconProps{},
										),
									),
								),
								htmx.H3(
									htmx.Text(
										user.Key.ID,
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
