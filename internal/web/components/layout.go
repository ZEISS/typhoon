package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/dividers"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/navbars"
	"github.com/zeiss/fiber-htmx/components/swap"
)

// LayoutProps is the properties for the Layout component.
type LayoutProps struct {
	Children []htmx.Node
	Team     adapters.GothTeam
	User     adapters.GothUser
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
						htmx.ID("alerts"),
						// toasts.ToastEnd(
						// 	toasts.ToastProps{},
						// 	toasts.ToastAlertInfo(
						// 		htmx.Text("Info"),
						// 	),
						// 	toasts.ToastAlertError(
						// 		htmx.Text("Error"),
						// 	),
						// 	toasts.ToastAlertSuccess(
						// 		htmx.Text("Success"),
						// 	),
						// ),
					),
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
							htmx.Group(
								children...,
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
						MainMenu(
							MainMenuProps{
								Path: p.Path,
							},
						),
						dividers.Divider(
							dividers.DividerProps{},
						),
						UserMenu(
							UserMenuProps{
								Path: p.Path,
							},
						),
					),
				),
			),
		),
	)
}
