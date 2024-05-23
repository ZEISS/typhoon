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
			Title:    "index",
			Language: "en",
			Attributes: []htmx.Node{
				htmx.DataAttribute("theme", "light"),
			},
			Head: []htmx.Node{
				htmx.Link(
					htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui/dist/full.css"),
					htmx.Attribute("rel", "stylesheet"),
					htmx.Attribute("type", "text/css"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://cdn.tailwindcss.com"),
				),
				htmx.Script(
					htmx.Attribute("src", "https://unpkg.com/htmx.org@1.9.12"),
					htmx.Integrity("sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2"),
					htmx.CrossOrigin("anonymous"),
				),
			},
		},
		htmx.Body(
			htmx.HxBoost(true),
			htmx.Group(children...),
		),
	)
}
