package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/forms"
)

// AccountSwitchProps is the properties for the AccountSwitch component.
type AccountSwitchProps struct {
	// ClassNames is the class names for the component.
	ClassNames htmx.ClassNames
	// User is the user to user adapters.
	User adapters.GothUser
}

// AccountSwitch ...
func AccountSwitch(p AccountSwitchProps) htmx.Node {
	var teams []*adapters.GothTeam

	if p.User.Teams != nil {
		for _, team := range *p.User.Teams {
			teams = append(teams, &team)
		}
	}

	return forms.SelectBordered(
		forms.SelectProps{
			ClassNames: htmx.ClassNames{
				"w-full": false,
			},
		},
		htmx.Group(
			htmx.ForEach(teams, func(team *adapters.GothTeam, idx int) htmx.Node {
				return forms.Option(
					forms.OptionProps{
						Value: team.ID.String(),
					},
					htmx.Text(team.Name),
				)
			})...,
		),
	)
}
