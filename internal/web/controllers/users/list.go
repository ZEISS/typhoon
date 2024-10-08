package users

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/users"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&ListUsersController{})

// ListUsersController ...
type ListUsersController struct {
	users tables.Results[models.User]
	store ports.Datastore
	htmx.DefaultController
}

// NewListUsersController ...
func NewListUsersController(store ports.Datastore) *ListUsersController {
	return &ListUsersController{
		users: tables.Results[models.User]{SearchFields: []string{"name"}},
		store: store,
	}
}

// Prepare ...
func (l *ListUsersController) Prepare() error {
	err := l.BindQuery(&l.users)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListUsers(ctx, &l.users)
	})
}

// Prepare ...
func (l *ListUsersController) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "Users",
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
						users.UsersTable(
							users.UsersTableProps{
								Users:  l.users.GetRows(),
								Offset: l.users.GetOffset(),
								Limit:  l.users.GetLimit(),
								Total:  l.users.GetLen(),
								URL:    l.OriginalURL(),
							},
						),
					),
				)
			},
		),
	)
}
