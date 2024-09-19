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
	accounts tables.Results[models.Account]
	store    ports.Datastore
	htmx.DefaultController
}

// NewListAccountsController ...
func NewListAccountsController(store ports.Datastore) *ListAccountsController {
	return &ListAccountsController{
		accounts: tables.Results[models.Account]{SearchFields: []string{"name"}},
		store:    store,
	}
}

// Prepare ...
func (l *ListAccountsController) Prepare() error {
	err := l.BindQuery(&l.accounts)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListAccounts(ctx, &l.accounts)
	})
}

// Prepare ...
func (l *ListAccountsController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Accounts",
				Path:  l.Path(),
				User:  l.Session().User,
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
								Accounts: l.accounts.GetRows(),
								Offset:   l.accounts.GetOffset(),
								Limit:    l.accounts.GetLimit(),
								Total:    l.accounts.GetLen(),
								URL:      l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
