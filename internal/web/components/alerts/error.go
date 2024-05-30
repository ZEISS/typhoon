package alerts

import (
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/toasts"
)

// ErrorProps ...
type ErrorProps struct {
	ClassNames htmx.ClassNames
	Error      error
	ID         string
}

// Error ...
func Error(props ErrorProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{},
			props.ClassNames,
		),
		htmx.ID(props.ID),
		htmx.HxSwapOob("true"),
		toasts.ToastEnd(
			toasts.ToastProps{},
			toasts.ToastAlertError(
				icons.ExclamationCircleOutline(
					icons.IconProps{},
				),
				htmx.Span(htmx.Text(props.Error.Error())),
			),
		),
	)
}
