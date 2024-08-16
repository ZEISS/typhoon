package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/pkg/cast"
)

// ProfileMenuProps ...
type ProfileMenuProps struct {
	ClassNames htmx.ClassNames
	User       adapters.GothUser
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
					htmx.Attribute("src", cast.Value(p.User.Image)),
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
				htmx.A(
					htmx.Attribute("href", "/logout"),
					htmx.Text("Logout"),
				),
			),
		),
	)
}
