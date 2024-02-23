package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/httptarget"
)

func main() {
	pkgadapter.Main("httptarget", httptarget.EnvAccessorCtor, httptarget.NewTarget)
}
