package main

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/synchronizer"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	pkgadapter.Main("synchronizer", synchronizer.EnvAccessorCtor, synchronizer.NewAdapter)
}
