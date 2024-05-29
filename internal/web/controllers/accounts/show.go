package accounts

import (
	"fmt"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/components/nkeys"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
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

	skgs := []*models.SigningKeyGroup{}
	for _, skg := range acc.SigningKeyGroups {
		skgs = append(skgs, &skg)
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
											htmx.Href(fmt.Sprintf("/accounts/%s/token", acc.ID)),
											htmx.Text("Download JWT Token"),
										),
									),
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										buttons.Error(
											buttons.ButtonProps{
												ClassNames: htmx.ClassNames{
													"btn-sm": true,
												},
											},
											htmx.HxDelete(fmt.Sprintf("/accounts/%s", acc.ID)),
											htmx.Text("Delete Account"),
										),
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
							htmx.Text("Details"),
						),
						nkeys.NKey(
							nkeys.NKeyProps{
								Title:     "ID",
								PublicKey: acc.Key.ID,
							},
						),
						nkeys.NKey(
							nkeys.NKeyProps{
								Title:     "Issuer",
								PublicKey: acc.Operator.Key.ID,
							},
						),
					),
				),
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Signing Key Groups"),
						),

						accounts.SigningKeyGroupsTable(
							accounts.SigningKeyGroupsTableProps{
								SigningKeyGroups: skgs,
							},
						),
						cards.Actions(
							cards.ActionsProps{},

							htmx.A(
								htmx.ClassNames{
									"btn": true,
								},
								htmx.Href(fmt.Sprintf("/accounts/%s/skgs/new", acc.ID)),
								icons.PlusOutline(
									icons.IconProps{},
								),
								htmx.Text("Signing Key Group"),
							),
						),
					),
				),
			),
		),
	)
}
