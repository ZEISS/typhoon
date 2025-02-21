package main

import (
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/signals"

	"github.com/zeiss/typhoon/pkg/sources/adapter/webhooksource"
)

const (
	component = "webhooksource-adapter"
)

func main() {
	sctx := signals.NewContext()

	ctx := adapter.WithController(sctx, webhooksource.NewController)
	ctx = adapter.WithHAEnabled(ctx)
	ctx = adapter.WithConfigWatcherEnabled(ctx)

	adapter.MainWithContext(ctx, component, webhooksource.NewEnvConfig, webhooksource.NewAdapter)
}
