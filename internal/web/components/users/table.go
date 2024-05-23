package users

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
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
		htmx.ClassNames{
			"bg-base-100": true,
			"m-4":         true,
		},
		tables.Table(
			tables.TableProps[*models.User]{
				ID: "accounts-tables",
				Columns: []tables.ColumnDef[*models.User]{
					{
						ID:          "id",
						AccessorKey: "id",
						Header: func(p tables.TableProps[*models.User]) htmx.Node {
							return htmx.Th(htmx.Text("ID"))
						},
						Cell: func(p tables.TableProps[*models.User], row *models.User) htmx.Node {
							return htmx.Td(
								htmx.Text(row.ID.String()),
							)
						},
					},
					{
						ID:          "name",
						AccessorKey: "name",
						Header: func(p tables.TableProps[*models.User]) htmx.Node {
							return htmx.Th(htmx.Text("Name"))
						},
						Cell: func(p tables.TableProps[*models.User], row *models.User) htmx.Node {
							return htmx.Td(
								links.Link(
									links.LinkProps{
										Href: "/users/" + row.ID.String(),
									},
									htmx.Text(row.Name),
								),
							)
						},
					},
					{
						Header: func(p tables.TableProps[*models.User]) htmx.Node {
							return nil
						},
						Cell: func(p tables.TableProps[*models.User], row *models.User) htmx.Node {
							return htmx.Td(
								buttons.Button(
									buttons.ButtonProps{
										ClassNames: htmx.ClassNames{
											"btn-square": true,
										},
									},
								),
							)
						},
					},
				},
				Rows: tables.NewRows(props.Users),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps[*models.User]{
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
				// Pagination: ProfileListTablePaginationComponent(
				// 	ProfileListTablePaginationProps{
				// 		Limit:  props.Limit,
				// 		Offset: props.Offset,
				// 		Total:  props.Total,
				// 		Target: "profiles-tables",
				// 		Team:   props.Team,
				// 	},
				// ),
			},
		),
	)
}
