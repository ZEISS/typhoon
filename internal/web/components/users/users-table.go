package users

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/pkg/stringx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
)

const (
	userShowURLFormat = "/users/%s"
	accountShowURL    = "/accounts/%s"
)

// UsersTableProps ...
type UsersTableProps struct {
	URL    string
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
				ID: "users-tables",
				Pagination: tables.TablePagination(
					tables.TablePaginationProps{},
					tables.Pagination(
						tables.PaginationProps{},
						tables.Prev(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),

						tables.Select(
							tables.SelectProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								Limits: tables.DefaultLimits,
								URL:    props.URL,
							},
						),
						tables.Next(
							tables.PaginationProps{
								Total:  props.Total,
								Offset: props.Offset,
								Limit:  props.Limit,
								URL:    props.URL,
							},
						),
					),
				),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"inline-flex":  true,
							"items-center": true,
							"gap-3":        true,
						},
						tables.Search(
							tables.SearchProps{
								Placeholder: "Search users...",
								Name:        "search",
								URL:         props.URL,
							},
						),
					),
					htmx.A(
						htmx.Href("/users/new"),
						buttons.Button(
							buttons.ButtonProps{},
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
						return htmx.Td(htmx.Text(stringx.FirstN(row.Key.ID, 8)))
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
							links.Link(
								links.LinkProps{
									ClassNames: htmx.ClassNames{
										"truncate": true,
									},
									Href: fmt.Sprintf(accountShowURL, row.Account.ID),
								},
								htmx.Text(row.Account.Name),
							),
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
							htmx.ClassNames{},
							htmx.A(
								htmx.ClassNames{
									"btn":  true,
									"mx-2": true,
								},
								htmx.Href(fmt.Sprintf(utils.DownloadCredentialsUserUrlFormat, row.ID)),
								icons.ArrowDownOnSquareOutline(
									icons.IconProps{},
								),
							),
							buttons.Button(
								buttons.ButtonProps{},
								htmx.HxDelete(fmt.Sprintf(utils.DeleteUserUrlFormat, row.ID)),
								htmx.HxConfirm("Are you sure you want to delete this user?"),
								htmx.HxTarget("closest tr"),
								htmx.HxSwap("outerHTML swap:1s"),
								icons.TrashOutline(
									icons.IconProps{},
								),
							),
						)

						// return htmx.Td(
						// 	dropdowns.Dropdown(
						// 		dropdowns.DropdownProps{},
						// 		dropdowns.DropdownButton(
						// 			dropdowns.DropdownButtonProps{},
						// 			icons.BoltOutline(
						// 				icons.IconProps{},
						// 			),
						// 		),
						// 		dropdowns.DropdownMenuItems(
						// 			dropdowns.DropdownMenuItemsProps{},
						// 			dropdowns.DropdownMenuItem(
						// 				dropdowns.DropdownMenuItemProps{},
						// 				htmx.A(
						// 					htmx.Href(fmt.Sprintf("/users/%s/credentials", row.ID)),
						// 					htmx.Text("Get Credentials"),
						// 				),
						// 			),
						// 			dropdowns.DropdownMenuItem(
						// 				dropdowns.DropdownMenuItemProps{},
						// 				buttons.Error(
						// 					buttons.ButtonProps{
						// 						ClassNames: htmx.ClassNames{
						// 							"btn-sm": true,
						// 						},
						// 					},
						// 					htmx.HxDelete(fmt.Sprintf("/users/%s", row.ID)),
						// 					htmx.HxConfirm("Are you sure you want to delete this user?"),
						// 					htmx.Text("Delete"),
						// 				),
						// 			),
						// 		),
						// 	),
						// )
					},
				},
			},
			props.Users,
		),
	)
}
