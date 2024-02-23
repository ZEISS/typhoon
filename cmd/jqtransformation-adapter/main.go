package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/flow/adapter/jqtransformation"
)

func main() {
	pkgadapter.Main("jqtransformation", jqtransformation.EnvAccessorCtor, jqtransformation.NewAdapter)
}
