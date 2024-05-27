package accounts

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/links"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
)

// SigningKeyGroupsTableProps ...
type SigningKeyGroupsTableProps struct {
	ClassNames       htmx.ClassNames
	SigningKeyGroups []*models.SigningKeyGroup
	Offset           int
	Limit            int
	Total            int
}

// SigningKeyGroupsTable ...
func SigningKeyGroupsTable(props SigningKeyGroupsTableProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			props.ClassNames,
		),
		tables.Table(
			tables.TableProps{
				ID: "signing-key-groups-table",
				Toolbar: tables.TableToolbar(
					tables.TableToolbarProps{
						ClassNames: htmx.ClassNames{
							"flex":            true,
							"items-center":    true,
							"justify-between": true,
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
			[]tables.ColumnDef[*models.SigningKeyGroup]{
				{
					ID:          "id",
					AccessorKey: "id",
					Header: func(p tables.TableProps) htmx.Node {
						return htmx.Th(htmx.Text("ID"))
					},
					Cell: func(p tables.TableProps, row *models.SigningKeyGroup) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.SigningKeyGroup) htmx.Node {
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
					Cell: func(p tables.TableProps, row *models.SigningKeyGroup) htmx.Node {
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
			props.SigningKeyGroups,
		),
	)
}
