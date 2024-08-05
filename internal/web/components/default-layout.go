package components

import htmx "github.com/zeiss/fiber-htmx"

// DefaultLayoutProps ...
type DefaultLayoutProps struct {
	ClassNames htmx.ClassNames
	Path       string
	Title      string
}

// DefaultLayout ...
func DefaultLayout(props DefaultLayoutProps, children ...htmx.Node) htmx.Node {
	return Page(
		PageProps{
			Title: props.Title,
		},
		Layout(
			LayoutProps{
				Path: props.Path,
			},
			htmx.Fallback(
				htmx.ErrorBoundary(
					func() htmx.Node {
						return htmx.Fragment(children...)
					},
				),
				func(err error) htmx.Node {
					return htmx.Text(err.Error())
				},
			),
		),
	)
}
