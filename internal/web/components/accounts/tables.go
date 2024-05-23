package accounts

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
)

// AccountsTableProps ...
type AccountsTableProps struct {
	Accounts []*models.Account
	Offset   int
	Limit    int
	Total    int
}

// AccountsTable ...
func AccountsTable(props AccountsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"bg-base-100": true,
			"m-4":         true,
		},
		tables.Table(
			tables.TableProps[*models.Account]{
				ID: "accounts-tables",
				Columns: []tables.ColumnDef[*models.Account]{
					{
						ID:          "id",
						AccessorKey: "id",
						Header: func(p tables.TableProps[*models.Account]) htmx.Node {
							return htmx.Th(htmx.Text("ID"))
						},
						Cell: func(p tables.TableProps[*models.Account], row *models.Account) htmx.Node {
							return htmx.Td(
								htmx.Text(row.ID.String()),
							)
						},
					},
					{
						ID:          "name",
						AccessorKey: "name",
						Header: func(p tables.TableProps[*models.Account]) htmx.Node {
							return htmx.Th(htmx.Text("Name"))
						},
						Cell: func(p tables.TableProps[*models.Account], row *models.Account) htmx.Node {
							return htmx.Td(
								links.Link(
									links.LinkProps{
										Href: "/accounts/" + row.ID.String(),
									},
									htmx.Text(row.Name),
								),
							)
						},
					},
					{
						Header: func(p tables.TableProps[*models.Account]) htmx.Node {
							return nil
						},
						Cell: func(p tables.TableProps[*models.Account], row *models.Account) htmx.Node {
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
				Rows: tables.NewRows(props.Accounts),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps[*models.Account]{
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
								Placeholder: "Search ...",
							},
						),
					),
					htmx.A(
						htmx.Href("/accounts/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Account"),
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
