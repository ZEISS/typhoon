package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/jiratarget"
)

func main() {
	adapter.Main("jiratarget", jiratarget.EnvAccessorCtor, jiratarget.NewTarget)
}
