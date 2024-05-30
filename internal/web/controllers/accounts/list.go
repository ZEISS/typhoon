package accounts

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/accounts"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&ListAccountsController{})

// ListAccountsController ...
type ListAccountsController struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	Search string `json:"search" form:"search"`
	Sort   string `json:"sort" form:"sort"`

	ports.Accounts
	htmx.DefaultController
}

// NewListAccountsController ...
func NewListAccountsController(db ports.Accounts) *ListAccountsController {
	return &ListAccountsController{
		Accounts:          db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *ListAccountsController) Prepare() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	return nil
}

// Prepare ...
func (l *ListAccountsController) Get() error {
	pagination := models.Pagination[models.Account]{
		Offset: l.Offset,
		Limit:  l.Limit,
		Sort:   l.Sort,
		Search: l.Search,
	}

	err := l.ListAccounts(l.Context(), &pagination)
	if err != nil {
		return err
	}

	accs := make([]*models.Account, 0, len(pagination.Rows))
	for _, row := range pagination.Rows {
		accs = append(accs, &row)
	}

	return htmx.RenderComp(
		l.Ctx(),
		components.Page(
			components.PageProps{
				Title: "Accounts",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						accounts.AccountsTable(
							accounts.AccountsTableProps{
								Accounts: accs,
								Offset:   pagination.Offset,
								Limit:    pagination.Limit,
								Total:    pagination.TotalRows,
							},
						),
					),
				),
			),
		),
	)
}
