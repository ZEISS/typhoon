package accounts

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
)

var _ = htmx.Controller(&ListAccountsController{})

// ListAccountsController ...
type ListAccountsController struct {
	Results tables.Results[models.Account]

	store ports.Datastore
	htmx.DefaultController
}

// NewListAccountsController ...
func NewListAccountsController(store ports.Datastore) *ListAccountsController {
	return &ListAccountsController{store: store}
}

// Prepare ...
func (l *ListAccountsController) Prepare() error {
	err := l.BindQuery(&l.Results)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListAccounts(ctx, &l.Results)
	})
}

// Prepare ...
func (l *ListAccountsController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Accounts",
				Path:  l.Path(),
			},
			func() htmx.Node {
				return cards.CardBordered(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							tailwind.M2: true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						accounts.AccountsTable(
							accounts.AccountsTableProps{
								Accounts: l.Results.GetRows(),
								Offset:   l.Results.GetOffset(),
								Limit:    l.Results.GetLimit(),
								Total:    l.Results.GetLen(),
							},
						),
					),
				)
			},
		),
	)
}
