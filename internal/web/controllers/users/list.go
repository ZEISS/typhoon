package users

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/components/users"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&ListUsersController{})

// ListUsersController ...
type ListUsersController struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	Search string `json:"search" form:"search"`
	Sort   string `json:"sort" form:"sort"`

	ports.Users
	htmx.DefaultController
}

// NewListUsersController ...
func NewListUsersController(db ports.Users) *ListUsersController {
	return &ListUsersController{
		Users:             db,
		DefaultController: htmx.DefaultController{},
	}
}

// Prepare ...
func (l *ListUsersController) Prepare() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	return nil
}

// Prepare ...
func (l *ListUsersController) Get() error {
	pagination := models.Pagination[models.User]{
		Offset: l.Offset,
		Limit:  l.Limit,
		Sort:   l.Sort,
		Search: l.Search,
	}

	err := l.ListUsers(l.Context(), &pagination)
	if err != nil {
		return err
	}

	accs := make([]*models.User, 0, len(pagination.Rows))
	for _, row := range pagination.Rows {
		accs = append(accs, &row)
	}

	return htmx.RenderComp(
		l.Ctx(),
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
								Users:  accs,
								Offset: pagination.Offset,
								Limit:  pagination.Limit,
								Total:  pagination.TotalRows,
							},
						),
					),
				),
			),
		),
	)
}
