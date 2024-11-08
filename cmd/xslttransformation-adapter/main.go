//go:build !noclibs

package main

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/xslttransformation"
	"knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	adapter.Main("xslttransformation", xslttransformation.EnvAccessorCtor, xslttransformation.NewTarget)
}
