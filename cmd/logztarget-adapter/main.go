package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/logztarget"
)

func main() {
	pkgadapter.Main("logztarget", logztarget.EnvAccessorCtor, logztarget.NewTarget)
}
