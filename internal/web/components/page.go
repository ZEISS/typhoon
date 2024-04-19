package components

import (
	htmx "github.com/zeiss/fiber-htmx"
)

// PageProps is the properties for the Page component.
type PageProps struct {
	Title    string
	Path     string
	Children []htmx.Node
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
			Head: []htmx.Node{
				htmx.Link(
					htmx.Attribute("href", "/static/output.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
				),
				htmx.Script(
					htmx.Attribute("src", "/static/output.js"),
					htmx.Attribute("type", "application/javascript"),
				),
			},
		},
		htmx.Body(
			// htmx.HxBoost(true),
			htmx.Group(children...),
		),
	)
}
