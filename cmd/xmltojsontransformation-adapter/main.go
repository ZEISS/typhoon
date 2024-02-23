package main

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/flow/adapter/xmltojsontransformation"
)

func main() {
	pkgadapter.Main("xmltojsontransformation", xmltojsontransformation.EnvAccessorCtor, xmltojsontransformation.NewAdapter)
}
