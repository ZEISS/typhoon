package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/cloudeventstarget"
)

func main() {
	adapter.Main("cloudeventstarget", cloudeventstarget.EnvAccessorCtor, cloudeventstarget.NewTarget)
}
