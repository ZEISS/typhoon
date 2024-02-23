package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/webhooksource"
)

func main() {
	adapter.Main("webhook", webhooksource.NewEnvConfig, webhooksource.NewAdapter)
}
