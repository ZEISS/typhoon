package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation"
)

func main() {
	adapter.Main("transformation", transformation.NewEnvConfig, transformation.NewAdapter)
}
