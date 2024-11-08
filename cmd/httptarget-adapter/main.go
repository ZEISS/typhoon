package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/httptarget"
)

func main() {
	adapter.Main("httptarget", httptarget.EnvAccessorCtor, httptarget.NewTarget)
}
