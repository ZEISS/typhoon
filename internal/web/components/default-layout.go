package components

import (
	"github.com/zeiss/fiber-goth/adapters"
	htmx "github.com/zeiss/fiber-htmx"
)

// DefaultLayoutProps ...
type DefaultLayoutProps struct {
	ClassNames htmx.ClassNames
	User       adapters.GothUser
	Path       string
	Title      string
}

// DefaultLayout ...
func DefaultLayout(props DefaultLayoutProps, node htmx.ErrBoundaryFunc) htmx.Node {
	return Page(
		PageProps{
			Title: props.Title,
		},
		Layout(
			LayoutProps{
				Path: props.Path,
				User: props.User,
			},
			htmx.Fallback(
				htmx.ErrorBoundary(node),
				func(err error) htmx.Node {
					return htmx.Text(err.Error())
				},
			),
		),
	)
}
