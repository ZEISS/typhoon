package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/servicenowtarget"
)

func main() {
	pkgadapter.Main("servicenowtarget", servicenowtarget.EnvAccessorCtor, servicenowtarget.NewTarget)
}
