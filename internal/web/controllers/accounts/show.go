package accounts

import (
	"context"
	"fmt"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/components/nkeys"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/cast"
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
			components.DefaultLayoutProps{
				Title: "Account",
				Path:  l.Path(),
			},
			func() htmx.Node {
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
										htmx.Text("Name"),
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
								htmx.A(
									htmx.ClassNames{
										"btn": true,
									},
									htmx.Href(fmt.Sprintf(utils.DownloadTokenAccountUrl, account.ID)),
									icons.ArrowDownOnSquareOutline(
										icons.IconProps{},
									),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(fmt.Sprintf(utils.DeleteAccountUrlFormat, account.ID)),
									htmx.HxConfirm("Are you sure you want to delete this account?"),
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
									SigningKeyGroups: cast.PtrSlice(account.SigningKeyGroups...),
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
	)
}
