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
	adapter.MainWithContext(signals.NewContext(), component, webhooksource.NewEnvConfig, webhooksource.NewAdapter)
}
