package accounts

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/components/nkeys"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
)

// ShowAccountControllerImpl ...
type ShowAccountControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewShowAccountController ...
func NewShowAccountController(store ports.Datastore) *ShowAccountControllerImpl {
	return &ShowAccountControllerImpl{store: store}
}

// Get ...
func (l *ShowAccountControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{},
			htmx.ControllerComponent(l, func(ctrl htmx.Controller) htmx.Node {
				var params struct {
					ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`
				}

				err := l.BindParams(&params)
				if err != nil {
					panic(err)
				}

				account := models.Account{ID: params.ID}

				err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetAccount(ctx, &account)
				})
				if err != nil {
					panic(err)
				}

				skgs := []*models.SigningKeyGroup{}
				for _, skg := range account.SigningKeyGroups {
					skgs = append(skgs, &skg)
				}

				return htmx.Fragment(
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
								htmx.Text("Overview"),
							),
							htmx.Div(
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
										htmx.Text("DescripNtion"),
									),
									htmx.H3(
										htmx.Text(account.Name),
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
										htmx.Text("Description"),
									),
									htmx.H3(
										htmx.Text(cast.Value(account.Description)),
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
											account.CreatedAt.Format("2006-01-02 15:04:05"),
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
											account.UpdatedAt.Format("2006-01-02 15:04:05"),
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
												htmx.Href(fmt.Sprintf("/accounts/%s/token", account.ID)),
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
												htmx.HxDelete(fmt.Sprintf("/accounts/%s", account.ID)),
												htmx.Text("Delete Account"),
											),
										),
									),
								),
							),
						),
					),
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
								htmx.Text("Details"),
							),
							nkeys.NKey(
								nkeys.NKeyProps{
									Title:     "ID",
									PublicKey: account.Key.ID,
								},
							),
							nkeys.NKey(
								nkeys.NKeyProps{
									Title: "Issuer",
									// PublicKey: acc.Operator.Key.ID,
								},
							),
						),
					),
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
									htmx.Href(fmt.Sprintf("/accounts/%s/skgs/new", account.ID)),
									icons.PlusOutline(
										icons.IconProps{},
									),
									htmx.Text("Signing Key Group"),
								),
							),
						),
					),
				)
			}),
		),
	)
}
