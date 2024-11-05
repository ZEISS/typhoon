package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/natstarget"
)

func main() {
	adapter.Main("natstarget", natstarget.EnvAccessorCtor, natstarget.NewTarget)
}
