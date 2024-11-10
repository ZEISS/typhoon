package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/azservicebus"
)

func main() {
	adapter.Main("azureservicebussource", azservicebus.NewEnvConfig, azservicebus.NewAdapter)
}
