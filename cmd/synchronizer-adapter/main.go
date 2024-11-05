package main

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/synchronizer"
	"knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	adapter.Main("synchronizer", synchronizer.EnvAccessorCtor, synchronizer.NewAdapter)
}
