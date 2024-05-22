package components

import (
	authz "github.com/zeiss/fiber-authz"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/fiber-htmx/components/navbars"
	"github.com/zeiss/fiber-htmx/components/swap"
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
		htmx.Div(
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
					htmx.Div(
						htmx.ClassNames{
							"overflow-auto": true,
							"w-full":        true,
							"h-screen":      true,
							"max-w-full":    true,
						},
						htmx.Div(
							htmx.ClassNames{
								"flex":        true,
								"h-full":      true,
								"flex-col":    true,
								"bg-base-200": true,
							},
							navbars.Navbar(
								navbars.NavbarProps{
									ClassNames: htmx.ClassNames{
										"navbar":      true,
										"z-10":        true,
										"border-b":    true,
										"px-3":        true,
										"bg-base-100": true,
									},
								},
								navbars.NavbarStart(
									navbars.NavbarStartProps{
										ClassNames: htmx.ClassNames{
											"gap-3": true,
										},
									},
									drawers.DrawerOpenButton(
										drawers.DrawerOpenProps{
											ID: "global-drawer",
											ClassNames: htmx.ClassNames{
												"btn-sm":      true,
												"btn-square":  true,
												"btn-primary": false,
											},
										},
										icons.Bars3Outline(
											icons.IconProps{},
										),
									),
								),
								navbars.NavbarEnd(
									navbars.NavbarEndProps{},
									swap.Swap(
										swap.SwapProps{
											ClassNames: htmx.ClassNames{
												"swap-rotate": true,
											},
										},
										htmx.Input(
											htmx.Class("theme-controller"),
											htmx.Value("dark"),
											htmx.Attribute("type", "checkbox"),
										),
										swap.SwapOn(
											swap.SwapProps{},
											icons.MoonOutlineSmall(
												icons.IconProps{},
											),
										),
										swap.SwapOff(
											swap.SwapProps{},
											icons.SunOutlineSmall(
												icons.IconProps{},
											),
										),
									),
									buttons.CircleSmall(
										buttons.ButtonProps{},
										icons.BellAlertOutlineSmall(
											icons.IconProps{},
										),
									),
									dropdowns.Dropdown(
										dropdowns.DropdownProps{
											ClassNames: htmx.ClassNames{
												"dropdown-end": true,
											},
										},
										dropdowns.DropdownButton(
											dropdowns.DropdownButtonProps{
												ClassNames: htmx.ClassNames{
													"btn-sm":     true,
													"btn-circle": true,
													"btn-ghost":  true,
												},
											},
											avatars.AvatarRoundSmall(
												avatars.AvatarProps{},
												htmx.Img(
													htmx.Attribute("src", "https://avatars.githubusercontent.com/u/570959?v=4"),
												),
											),
										),
										dropdowns.DropdownMenuItems(
											dropdowns.DropdownMenuItemsProps{},
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.A(
													htmx.Text("Profile"),
												),
											),
										),
									),
								),
							),
							htmx.Div(
								htmx.ClassNames{
									"p-8":            true,
									"grid":           true,
									"gap-6":          true,
									"xl:grid-cols-2": true,
								},
								cards.CardBordered(
									cards.CardProps{
										ClassNames: htmx.ClassNames{
											"shadow-xl": false,
											"border":    true,
											"rounded":   true,
										},
									},
									cards.Body(
										cards.BodyProps{},
										cards.Title(
											cards.TitleProps{},
											htmx.Text("Hello, World!"),
										),
										htmx.Text("This is a card body."),
									),
								),
								cards.Card(
									cards.CardProps{
										ClassNames: htmx.ClassNames{
											"shadow-xl": false,
											"border":    true,
											"rounded":   true,
										},
									},
									cards.Body(
										cards.BodyProps{},
										cards.Title(
											cards.TitleProps{},
											htmx.Text("Hello, World!"),
										),
										htmx.Text("This is a card body."),
									),
								),
							),
						),
					),
				),
				drawers.DrawerSide(
					drawers.DrawerSideProps{
						ID: "drawer",
					},
					drawers.DrawerSideMenu(
						drawers.DrawerSideMenuProps{
							ClassNames: htmx.ClassNames{
								"border-r":    true,
								"bg-base-100": true,
								"bg-base-200": false,
							},
						},
						htmx.Nav(
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
										menus.MenuLinkProps{},
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
										menus.MenuLinkProps{},
										htmx.Text("Operators"),
									),
								),
								menus.MenuItem(
									menus.MenuItemProps{},
									menus.MenuLink(
										menus.MenuLinkProps{},
										htmx.Text("Accounts"),
									),
								),
								menus.MenuItem(
									menus.MenuItemProps{},
									menus.MenuLink(
										menus.MenuLinkProps{},
										htmx.Text("Users"),
									),
								),
								dividers.Divider(
									dividers.DividerProps{},
								),
								menus.MenuItem(
									menus.MenuItemProps{},
									menus.MenuLink(
										menus.MenuLinkProps{
											Href: "/me",
										},
										htmx.Text("Profile"),
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
								// menus.MenuItem(
								// 	menus.MenuItemProps{},
								// 	menus.MenuCollapsible(
								// 		menus.MenuCollapsibleProps{},
								// 		menus.MenuCollapsibleSummary(
								// 			menus.MenuCollapsibleSummaryProps{},
								// 			htmx.Text("Operators"),
								// 		),
								// 		menus.MenuItem(
								// 			menus.MenuItemProps{},
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 1"),
								// 			),
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 2"),
								// 			),
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 3"),
								// 			),
								// 		),
								// 	),
								// 	menus.MenuCollapsible(
								// 		menus.MenuCollapsibleProps{},
								// 		menus.MenuCollapsibleSummary(
								// 			menus.MenuCollapsibleSummaryProps{},
								// 			htmx.Text("Forms"),
								// 		),
								// 		menus.MenuItem(
								// 			menus.MenuItemProps{},
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 1"),
								// 			),
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 2"),
								// 			),
								// 			htmx.A(
								// 				htmx.Attribute("href", "#"),
								// 				htmx.Text("Item 3"),
								// 			),
								// 		),
								// 	),
								// ),
							),
						),
					),
				),
			),
		),
	)
}
