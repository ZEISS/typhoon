package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/httppollersource"
)

const component = "httppollersource-adapter"

func main() {
	adapter.Main(component, httppollersource.NewEnvConfig, httppollersource.NewAdapter)
}
