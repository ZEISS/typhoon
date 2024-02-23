package main

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/transformation"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	pkgadapter.Main("transformation", transformation.NewEnvConfig, transformation.NewAdapter)
}
