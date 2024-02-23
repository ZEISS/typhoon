package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/cloudeventssource"
)

func main() {
	adapter.Main("cloudevents", cloudeventssource.NewEnvConfig, cloudeventssource.NewAdapter)
}
