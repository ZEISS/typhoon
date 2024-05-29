package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/targets/adapter/jiratarget"
)

func main() {
	pkgadapter.Main("jiratarget", jiratarget.EnvAccessorCtor, jiratarget.NewTarget)
}
