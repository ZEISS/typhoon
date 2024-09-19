package operators

import (
	"context"
	"fmt"

	"github.com/zeiss/pkg/errorx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/nkeys"
	"github.com/zeiss/typhoon/internal/web/components/operators"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

// ShowOperatorControllerImpl ...
type ShowOperatorControllerImpl struct {
	store ports.Datastore
	htmx.DefaultController
}

// NewShowOperatorController ...
func NewShowOperatorController(store ports.Datastore) *ShowOperatorControllerImpl {
	return &ShowOperatorControllerImpl{
		store: store,
	}
}

// Get ...
func (l *ShowOperatorControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Operator",
				Path:  l.Path(),
				User:  l.Session().User,
			},
			func() htmx.Node {
				op := models.Operator{}
				err := l.BindParams(&op)
				errorx.Panic(err)

				errorx.Panic(l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
					return tx.GetOperator(ctx, &op)
				}))

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
								htmx.H1(
									htmx.Text(op.Name),
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
										htmx.Text(op.Description),
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
											op.CreatedAt.Format("2006-01-02 15:04:05"),
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
											op.UpdatedAt.Format("2006-01-02 15:04:05"),
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
									htmx.Href(fmt.Sprintf(utils.DownloadTokenOperatorUrl, op.ID)),
									icons.ArrowDownOnSquareOutline(
										icons.IconProps{},
									),
								),
								buttons.Button(
									buttons.ButtonProps{},
									htmx.HxDelete(fmt.Sprintf(utils.DeleteOperatorUrlFormat, op.ID)),
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
									PublicKey: op.Key.ID,
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
								htmx.Text("System Account"),
							),
							nkeys.NKey(
								nkeys.NKeyProps{
									Title:     "ID",
									PublicKey: op.SystemAccount.Key.ID,
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
								htmx.Text("System Users"),
							),
							operators.UsersTable(
								operators.UsersTableProps{
									Users: tables.RowsPtr(op.SystemAccount.Users),
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

							operators.SigningKeyGroupsTable(
								operators.SigningKeyGroupsTableProps{
									SigningKeyGroups: tables.RowsPtr(op.SigningKeyGroups),
								},
							),
							cards.Actions(
								cards.ActionsProps{},

								htmx.A(
									htmx.ClassNames{
										"btn": true,
									},
									htmx.Href(fmt.Sprintf("/operators/%s/skgs/new", op.ID)),
									icons.PlusOutline(
										icons.IconProps{},
									),
									htmx.Text("Signing Key Group"),
								),
							),
						),
					),
				)
			},
		),
	)
}
