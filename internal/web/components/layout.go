package components

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/icons"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node
	Team     *authz.Team
	User     *authz.User
	Path     string
}

// WrapProps ...
type WrapProps struct {
	ClassNames htmx.ClassNames
}

// Wrap ...
func Wrap(p WrapProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Layout is a whole document to output.
func Layout(p LayoutProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{},
		drawers.Drawer(
			drawers.DrawerProps{
				ID: "global-drawer",
				ClassNames: htmx.ClassNames{
					"lg:drawer-open": true,
				},
			},
			drawers.DrawerContent(
				drawers.DrawerContentProps{
					ID: "drawer",
				},
				SubNav(
					SubNavProps{
						ClassNames: htmx.ClassNames{
							"lg:hidden": true,
						},
					},
					drawers.DrawerOpenButton(
						drawers.DrawerOpenProps{
							ID: "global-drawer",
							ClassNames: htmx.ClassNames{
								"lg:hidden":   true,
								"btn-md":      true,
								"btn-square":  true,
								"btn-outline": true,
								"btn-primary": false,
							},
						},
						icons.Bars3Outline(
							icons.IconProps{},
						),
					),
				),
				Wrap(
					WrapProps{
						ClassNames: htmx.ClassNames{
							"m-6": true,
						},
					},
					htmx.Group(children...),
				),
			),
			drawers.DrawerSide(
				drawers.DrawerSideProps{
					ID: "drawer",
				},
				drawers.DrawerSideMenu(
					drawers.DrawerSideMenuProps{},
					dividers.Divider(
						dividers.DividerProps{
							ClassNames: htmx.ClassNames{
								"my-0": true,
							},
						},
					),
					UserMenu(
						UserMenuProps{
							Path: p.Path,
						},
					),
				),
			),
		),
	)
}
