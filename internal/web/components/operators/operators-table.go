package operators

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
		},
		tables.Table(
			tables.TableProps{
				ID: "operators-tables",
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
									URL:    "/operators",
								},
							),

							tables.Select(
								tables.SelectProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									Limits: tables.DefaultLimits,
									URL:    "/operators",
								},
							),
							tables.Next(
								tables.PaginationProps{
									Total:  props.Total,
									Offset: props.Offset,
									Limit:  props.Limit,
									URL:    "/operators",
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
			},
			[]tables.ColumnDef[*models.Operator]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.Operator) htmx.Node {
						return htmx.Td(
							htmx.Text(row.ID.String()),
						)
					},
				},
				{
					ID:          "name",
					AccessorKey: "name",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Name"))
					},
					Cell: func(p tables.TableProps, row *models.Operator) htmx.Node {
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
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.Operator) htmx.Node {
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
										buttons.Error(
											buttons.ButtonProps{
												ClassNames: htmx.ClassNames{
													"btn-sm": true,
												},
											},
											htmx.HxDelete(fmt.Sprintf("/operators/%s", row.ID)),
											htmx.HxConfirm("Are you sure you want to delete this operator?"),
											htmx.Text("Delete"),
										),
									),
								),
							),
						)
					},
				},
			},
			props.Operators,
		),
	)
}
