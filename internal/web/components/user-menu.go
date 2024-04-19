package components

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/menus"
)

// UserMenuProps ...
type UserMenuProps struct {
	ClassNames htmx.ClassNames
	Path       string
}

// UserMenu ...
func UserMenu(p UserMenuProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
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
						Href:   "/me",
						Active: p.Path == "/me",
					},
					htmx.Text("Profile"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href:   "/settings",
						Active: p.Path == "/settings",
					},
					htmx.Text("Settings"),
				),
			),
			menus.MenuItem(
				menus.MenuItemProps{},
				menus.MenuLink(
					menus.MenuLinkProps{
						Href: "/logout",
					},
					htmx.Text("Logout"),
				),
			),
		),
	)
}
