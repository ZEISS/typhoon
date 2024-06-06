package users

import (
	"context"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/users"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&ListUsersController{})

// ListUsersController ...
type ListUsersController struct {
	Results tables.Results[models.User]

	store ports.Datastore
	htmx.DefaultController
}

// NewListUsersController ...
func NewListUsersController(store ports.Datastore) *ListUsersController {
	return &ListUsersController{
		Results:           tables.Results[models.User]{Limit: 10},
		DefaultController: htmx.DefaultController{},
		store:             store,
	}
}

// Prepare ...
func (l *ListUsersController) Prepare() error {
	err := l.BindQuery(&l.Results)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.ListUsers(ctx, &l.Results)
	})
}

// Prepare ...
func (l *ListUsersController) Get() error {
	return l.Render(
		components.Page(
			components.PageProps{
				Title: "Users",
			},
			components.Layout(
				components.LayoutProps{
					Path: l.Ctx().Path(),
				},
				cards.CardBordered(
					cards.CardProps{},
					cards.Body(
						cards.BodyProps{},
						users.UsersTable(
							users.UsersTableProps{
								Users:  l.Results.GetRows(),
								Offset: l.Results.GetOffset(),
								Limit:  l.Results.GetLimit(),
								Total:  l.Results.GetLen(),
							},
						),
					),
				),
			),
		),
	)
}
