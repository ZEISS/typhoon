package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Boost    bool
	Children []htmx.Node
	Head     []htmx.Node
}

// Page is a whole document to output.
func Page(props PageProps, children ...htmx.Node) htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    props.Title,
			Language: "en",
			Attributes: []htmx.Node{
				htmx.DataAttribute("theme", "light"),
			},
			Head: append([]htmx.Node{
				htmx.Link(
					htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui/dist/full.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://cdn.tailwindcss.com"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/htmx.org@2.0.0"),
					htmx.CrossOrigin("anonymous"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/hyperscript.org@0.9.12"),
					htmx.Attribute("type", "application/javascript"),
				),
				htmx.Script(
					htmx.Attribute("src", "//unpkg.com/alpinejs"),
					htmx.Defer(),
				),
			}, props.Head...),
		},
		htmx.Body(
			htmx.HxBoost(props.Boost),
			htmx.Group(children...),
		),
	)
}
