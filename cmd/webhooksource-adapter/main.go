package main

import (
	"github.com/zeiss/typhoon/pkg/sources/adapter/webhooksource"
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/signals"
)

const (
	component = "webhooksource-adapter"
)

func main() {
	ctx := signals.NewContext()
	ctx = adapter.WithInjectorEnabled(ctx)

	adapter.MainWithContext(ctx, component, webhooksource.NewEnvConfig, webhooksource.NewAdapter)
}
