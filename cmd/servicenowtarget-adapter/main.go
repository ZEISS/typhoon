package main

import (
	"github.com/zeiss/typhoon/pkg/targets/adapter/servicenowtarget"
	"knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	adapter.Main("servicenowtarget", servicenowtarget.EnvAccessorCtor, servicenowtarget.NewTarget)
}
