package users

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
)

const (
	userShowURLFormat = "/users/%s"
)

// UsersTableProps ...
type UsersTableProps struct {
	Users  []*models.User
	Offset int
	Limit  int
	Total  int
}

// UsersTable ...
func UsersTable(props UsersTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "accounts-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{
						Pagination: tables.Pagination(
							tables.PaginationProps{
								Offset: props.Offset,
								Limit:  props.Limit,
								Total:  props.Total,
							},
							tables.Prev(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/users",
								},
							),

							tables.Select(
								tables.SelectProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									Limits: tables.DefaultLimits,
									URL:    "/users",
								},
							),
							tables.Next(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/users",
								},
							),
						),
					},
				),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
							"px-5":            true,
							"pt-5":            true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"inline-flex":  true,
							"items-center": true,
							"gap-3":        true,
						},
						forms.TextInputBordered(
							forms.TextInputProps{
								ClassNames: htmx.ClassNames{
									"input-sm": true,
								},
								Placeholder: "Search ...",
							},
						),
					),
					htmx.A(
						htmx.Href("/users/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create User"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.User]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.User) htmx.Node {
						return htmx.Td(
							htmx.Class("truncate"),
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "pubKey",
					AccessorKey: "pubKey",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Public Key"))
					},
					Cell: func(p tables.TableProps, row *models.User) htmx.Node {
						return htmx.Td(htmx.Text(utils.ShortPubKey(row.Key.ID)))
					},
				},
				{
					ID:          "accountName",
					AccessorKey: "accountName",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Account"))
					},
					Cell: func(p tables.TableProps, row *models.User) htmx.Node {
						return htmx.Td(
							htmx.Class("truncate"),
							htmx.Text(row.Account.Name),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.User) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{
									Href: fmt.Sprintf(userShowURLFormat, row.ID),
								},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.User) htmx.Node {
						return htmx.Td(
							dropdowns.Dropdown(
								dropdowns.DropdownProps{},
								dropdowns.DropdownButton(
									dropdowns.DropdownButtonProps{},
									icons.BoltOutline(
										icons.IconProps{},
									),
								),
								dropdowns.DropdownMenuItems(
									dropdowns.DropdownMenuItemsProps{},
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										htmx.A(
											htmx.Href(fmt.Sprintf("/users/%s/credentials", row.ID)),
											htmx.Text("Get Credentials"),
										),
									),
									dropdowns.DropdownMenuItem(
										dropdowns.DropdownMenuItemProps{},
										buttons.Error(
											buttons.ButtonProps{
												ClassNames: htmx.ClassNames{
													"btn-sm": true,
												},
											},
											htmx.HxDelete(fmt.Sprintf("/users/%s", row.ID)),
											htmx.Text("Delete User"),
										),
									),
								),
							),
						)
					},
				},
			},
			props.Users,
		),
	)
}
