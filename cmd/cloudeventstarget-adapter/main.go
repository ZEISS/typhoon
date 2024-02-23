package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/cloudeventstarget"
)

func main() {
	pkgadapter.Main("cloudeventstarget", cloudeventstarget.EnvAccessorCtor, cloudeventstarget.NewTarget)
}
