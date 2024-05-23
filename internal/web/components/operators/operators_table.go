package operators

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
)

// OperatorsTableProps ...
type OperatorsTableProps struct {
	Operators []*models.Operator
	Offset    int
	Limit     int
	Total     int
}

// OperatorsTable ...
func OperatorsTable(props OperatorsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"bg-base-100": true,
			"m-4":         true,
		},

		tables.Table(
			tables.TableProps[*models.Operator]{
				ID: "operators-tables",
				Columns: []tables.ColumnDef[*models.Operator]{
					{
						ID:          "id",
						AccessorKey: "id",
						Header: func(p tables.TableProps[*models.Operator]) htmx.Node {
							return htmx.Th(htmx.Text("ID"))
						},
						Cell: func(p tables.TableProps[*models.Operator], row *models.Operator) htmx.Node {
							return htmx.Td(
								htmx.Text(row.ID.String()),
							)
						},
					},
					{
						ID:          "name",
						AccessorKey: "name",
						Header: func(p tables.TableProps[*models.Operator]) htmx.Node {
							return htmx.Th(htmx.Text("Name"))
						},
						Cell: func(p tables.TableProps[*models.Operator], row *models.Operator) htmx.Node {
							return htmx.Td(
								links.Link(
									links.LinkProps{
										Href: "/operators/" + row.ID.String(),
									},
									htmx.Text(row.Name),
								),
							)
						},
					},
					{
						Header: func(p tables.TableProps[*models.Operator]) htmx.Node {
							return nil
						},
						Cell: func(p tables.TableProps[*models.Operator], row *models.Operator) htmx.Node {
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
				Rows: tables.NewRows(props.Operators),
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps[*models.Operator]{
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
						htmx.Href("/operators/new"),
						buttons.Outline(
							buttons.ButtonProps{
								ClassNames: htmx.ClassNames{
									"btn-sm": true,
								},
							},
							htmx.Text("Create Operator"),
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
