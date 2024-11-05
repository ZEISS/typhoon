package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/logztarget"
)

func main() {
	adapter.Main("logztarget", logztarget.EnvAccessorCtor, logztarget.NewTarget)
}
