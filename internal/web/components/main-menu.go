package components

import (
	"strings"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// MainMenuProps ...
type MainMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
}

// MainMenu ...
func MainMenu(p MainMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.Merge(
			htmx.ClassNames{},
		),
		menus.Menu(
			menus.MenuProps{
				ClassNames: htmx.ClassNames{
					"w-full":      true,
					"bg-base-200": false,
				},
			},
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/",
						Active: p.Path == "/",
					},
					htmx.Text("Dashboard"),
				),
			),
			menus.MenuTitle(
				menus.MenuTitleProps{},
				htmx.Text("Identity & Access"),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/teams",
						Active: strings.HasPrefix(p.Path, "/teams"),
					},
					htmx.Text("Teams"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/accounts",
						Active: strings.HasPrefix(p.Path, "/accounts"),
					},
					htmx.Text("Accounts"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/users",
						Active: strings.HasPrefix(p.Path, "/users"),
					},
					htmx.Text("Users"),
				),
			),
			menus.MenuTitle(
				menus.MenuTitleProps{},
				htmx.Text("System & Operators"),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/systems",
						Active: strings.HasPrefix(p.Path, "/systems"),
					},
					htmx.Text("Systems"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/operators",
						Active: strings.HasPrefix(p.Path, "/operators"),
					},
					htmx.Text("Operators"),
				),
			),
		),
	)
}
