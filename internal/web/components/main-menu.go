package components

import (
	"fmt"
	"strings"

	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// MainMenuProps ...
type MainMenuProps struct {
	ClassNames htmx.ClassNames
	Team       *authz.Team
	Path       string
}

// MainMenu ...
func MainMenu(p MainMenuProps, children ...htmx.Node) htmx.Node {
	if p.Team == nil {
		p.Team = &authz.Team{}
	}

	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		menus.Menu(
			menus.MenuProps{
				ClassNames: htmx.ClassNames{
					"w-full": true,
				},
			},
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   fmt.Sprintf("/%s", p.Team.Slug),
						Active: p.Path == fmt.Sprintf("/%s", p.Team.Slug),
					},
					htmx.Text("Dashboard"),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/workloads", p.Team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Workloads"),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/workloads/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/workloads/new", p.Team.Slug),
								},
								htmx.Text("New workload"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/workloads/list", p.Team.Slug),
								},
								htmx.Text("List workload"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/lenses", p.Team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Lenses"),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/lenses/new", p.Team.Slug),
								},
								htmx.Text("New Lens"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/lenses/list", p.Team.Slug),
								},
								htmx.Text("List Lens"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/profiles", p.Team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Profiles"),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/profiles/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/profiles/new", p.Team.Slug),
								},
								htmx.Text("New Profile"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/profiles/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/profiles/list", p.Team.Slug),
								},
								htmx.Text("List Profile"),
							),
						),
					),
				),
			),
			htmx.If(
				p.Team.Slug != "",
				menus.MenuItem(
					menus.MenuItemProps{
						ClassNames: htmx.ClassNames{
							"hover:bg-base-300": false,
						},
					},
					menus.MenuCollapsible(
						menus.MenuCollapsibleProps{
							Open: strings.HasPrefix(p.Path, fmt.Sprintf("/teams/%s/environments", p.Team.Slug)),
						},
						menus.MenuCollapsibleSummary(
							menus.MenuCollapsibleSummaryProps{},
							htmx.Text("Environments"),
						),
						menus.MenuItem(
							menus.MenuItemProps{
								ClassNames: htmx.ClassNames{
									"hover:bg-base-300": false,
								},
							},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/environments/new", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/environments/new", p.Team.Slug),
								},
								htmx.Text("New Environment"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   fmt.Sprintf("/teams/%s/environments/list", p.Team.Slug),
									Active: p.Path == fmt.Sprintf("/teams/%s/environments/list", p.Team.Slug),
								},
								htmx.Text("List Environment"),
							),
						),
					),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuCollapsible(
					menus.MenuCollapsibleProps{
						Open: strings.HasPrefix(p.Path, "/site"),
					},
					menus.MenuCollapsibleSummary(
						menus.MenuCollapsibleSummaryProps{},
						htmx.Text("Administration"),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/teams/new",
								Active: p.Path == "/site/teams/new",
							},
							htmx.Text("New Team"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/teams",
								Active: p.Path == "/site/teams",
							},
							htmx.Text("List Teams"),
						),
					),
					menus.MenuItem(
						menus.MenuItemProps{},
						menus.MenuLink(
							menus.MenuLinkProps{
								Href:   "/site/settings",
								Active: p.Path == "/site/settings",
							},
							htmx.Text("Settings"),
						),
					),
				),
			),
		),
	)
}
