package main

import (
	"github.com/zeiss/typhoon/pkg/flow/adapter/xmltojsontransformation"
	"knative.dev/eventing/pkg/adapter/v2"
)

func main() {
	adapter.Main("xmltojsontransformation", xmltojsontransformation.EnvAccessorCtor, xmltojsontransformation.NewAdapter)
}
