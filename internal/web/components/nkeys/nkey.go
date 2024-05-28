package nkeys

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/tooltips"
	"github.com/zeiss/typhoon/internal/utils"
)

// NKeyProps are the properties of the NKey display component.
type NKeyProps struct {
	ClassNames htmx.ClassNames
	PublicKey  string
	Title      string
}

// NKey renders the NKey display component.
func NKey(props NKeyProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"flex":     true,
				"flex-col": true,
				"py-2":     true,
			},
			props.ClassNames,
		),
		htmx.H4(
			htmx.ClassNames{
				"text-neutral-content": true,
				"flex":                 true,
				"items-center":         true,
			},
			htmx.Text(props.Title),
			tooltips.Tooltip(
				tooltips.TooltipProps{
					ClassNames: htmx.ClassNames{},
					DataTip:    "Public NKey is the public key of the operator",
				},
				icons.InformationCircleOutline(
					icons.IconProps{},
				),
			),
		),
		htmx.H3(
			tooltips.Tooltip(
				tooltips.TooltipProps{
					ClassNames: htmx.ClassNames{},
					DataTip:    props.PublicKey,
				},
				htmx.Text(
					utils.FirstN(props.PublicKey, 8),
				),
			),
		),
	)
}
