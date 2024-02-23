package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/httppollersource"
)

func main() {
	adapter.Main("httppoller", httppollersource.NewEnvConfig, httppollersource.NewAdapter)
}
