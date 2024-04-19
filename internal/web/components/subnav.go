package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// SubNavProps ...
type SubNavProps struct {
	ClassNames htmx.ClassNames
}

// SubNav ...
func SubNav(p SubNavProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar":         true,
				"bg-base-100":    true,
				"w-full":         true,
				"border-neutral": true,
				"border-b":       true,
				"border-t":       true,
				"px-6":           true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SubNavBreadcrumbProps ...
type SubNavBreadcrumbProps struct{}

// SubNavBreadcrumb ...
func SubNavBreadcrumb(p SubNavBreadcrumbProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"navbar-start": true,
		},
		htmx.Group(children...),
	)
}

// SubNavActionsProps ...
type SubNavActionsProps struct{}

// SubNavActions ...
func SubNavActions(p SubNavActionsProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"navbar-end": true,
		},
		htmx.Group(children...),
	)
}
