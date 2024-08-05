package components

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
)

// ProfileMenuProps ...
type ProfileMenuProps struct {
	ClassNames htmx.ClassNames
}

// ProfileMenu ...
func ProfileMenu(p ProfileMenuProps, children ...htmx.Node) htmx.Node {
	return dropdowns.Dropdown(
		dropdowns.DropdownProps{
			ClassNames: htmx.Merge(
				htmx.ClassNames{
					"dropdown-end": true,
				},
				p.ClassNames,
			),
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
					htmx.Attribute("href", "/me"),
					htmx.Text("Profile"),
				),
			),
		),
	)
}
