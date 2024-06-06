package operators

import (
	"context"
	"fmt"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/nkeys"
	"github.com/zeiss/typhoon/internal/web/components/operators"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
)

// ShowOperatorControllerImpl ...
type ShowOperatorControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" validate:"required:uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewShowOperatorController ...
func NewShowOperatorController(store ports.Datastore) *ShowOperatorControllerImpl {
	return &ShowOperatorControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *ShowOperatorControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	op := models.Operator{ID: l.ID}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetOperator(ctx, &op)
	})
	if err != nil {
		return err
	}

	skgs := []*models.SigningKeyGroup{}
	for _, skg := range op.SigningKeyGroups {
		skgs = append(skgs, &skg)
	}

	accs := []*models.Account{}
	for _, acc := range op.Accounts {
		accs = append(accs, &acc)
	}

	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{
				Boost: true,
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
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
								htmx.Text(op.Name),
							),
							htmx.P(
								htmx.Text(op.Description),
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
											htmx.Href(fmt.Sprintf("/operators/%s/token", op.ID)),
											htmx.Text("Download JWT Token"),
										),
									),
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										htmx.A(
											htmx.HxDelete(fmt.Sprintf("/operators/%s", op.ID)),
											htmx.HxConfirm("Are you sure you want to delete this operator?"),
											htmx.Text("Delete"),
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
							htmx.Text("System Account"),
						),
						htmx.Form(
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
										htmx.HxPut(fmt.Sprintf("/operators/%s/system-account", op.ID)),
										htmx.HxTarget("this"),
										htmx.HxSwap("outerHTML"),
										forms.Option(
											forms.OptionProps{
												Selected: true,
												Disabled: true,
											},
											htmx.Text("Select account"),
										),
										htmx.Name("system_account_id"),
										htmx.Group(
											htmx.ForEach(accs, func(account *models.Account) htmx.Node {
												return forms.Option(
													forms.OptionProps{
														Value:    account.ID.String(),
														Selected: op.SystemAdminAccountID != nil && account.ID == utils.UUIDPtr(op.SystemAdminAccountID),
													},
													htmx.Text(account.Name),
												)
											})...,
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
								PublicKey: op.Key.ID,
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
								SigningKeyGroups: skgs,
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
			),
		),
	)
}