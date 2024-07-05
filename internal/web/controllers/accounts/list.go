package accounts

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/ports"
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
	return &ListAccountsController{
		Results:           tables.Results[models.Account]{Limit: 10},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
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
		components.Page(
			components.PageProps{
				Title: "Accounts",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
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
				),
			),
		),
	)
}
