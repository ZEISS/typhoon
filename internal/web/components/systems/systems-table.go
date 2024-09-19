package systems

import (
	"fmt"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/utils"
)

// SystemsTableProps ...
type SystemsTableProps struct {
	URL     string
	Systems []*models.System
	Offset  int
	Limit   int
	Total   int
}

// SystemsTable ...
func SystemsTable(props SystemsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		tables.Table(
			tables.TableProps{
				ID: "systems-tables",
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
								Name:        "search",
								Placeholder: "Search systems ...",
								URL:         props.URL,
							},
						),
					),
					htmx.A(
						htmx.Href("/systems/new"),
						buttons.Button(
							buttons.ButtonProps{},
							htmx.Text("Create System"),
						),
					),
				),
			},
			[]tables.ColumnDef[*models.System]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.System) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.System) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{Href: fmt.Sprintf(utils.ShowSystemUrlFormat, row.ID)},
								htmx.Text(row.Name),
							),
						)
					},
				},
				{
					ID:          "operator",
					AccessorKey: "operator",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("Operator"))
					},
					Cell: func(p tables.TableProps, row *models.System) htmx.Node {
						return htmx.Td(
							links.Link(
								links.LinkProps{Href: fmt.Sprintf(utils.ShowOperatorUrlFormat, row.Operator.ID)},
								htmx.Text(row.Operator.Name),
							),
						)
					},
				},
				{
					Header: func(p tables.TableProps) htmx.Node {
						return nil
					},
					Cell: func(p tables.TableProps, row *models.System) htmx.Node {
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
											buttons.ButtonProps{},
											htmx.HxDelete(fmt.Sprintf("/systems/%s", row.ID)),
											htmx.HxConfirm("Are you sure you want to delete this system?"),
											htmx.Text("Delete"),
										),
									),
								),
							),
						)
					},
				},
			},
			props.Systems,
		),
	)
}
