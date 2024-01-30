package main

import (
	knative "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/webhook"
)

func main() {
	knative.Main("webhook", webhook.NewEnvconfig, webhook.NewAdapter)
}
