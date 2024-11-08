package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/flow/adapter/jqtransformation"
)

func main() {
	adapter.Main("jqtransformation", jqtransformation.EnvAccessorCtor, jqtransformation.NewAdapter)
}
