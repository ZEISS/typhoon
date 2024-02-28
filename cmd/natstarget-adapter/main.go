package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/natstarget"
)

func main() {
	pkgadapter.Main("natstarget", natstarget.EnvAccessorCtor, natstarget.NewTarget)
}
