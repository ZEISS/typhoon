// Package controller contains helpers shared between controllers embedded in
// source adapters.
package controller

import "knative.dev/pkg/controller"

// Opts returns a callback function that sets the controller's agent name and
// configures the reconciler to skip status updates.
func Opts(component string) controller.OptionsFn {
	return func(impl *controller.Impl) controller.Options {
		return controller.Options{
			AgentName:         component,
			SkipStatusUpdates: true,
		}
	}
}
