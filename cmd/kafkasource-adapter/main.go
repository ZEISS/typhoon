package main

import (
	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/kafkasource"
)

func main() {
	adapter.Main("kafkasource", kafkasource.NewEnvConfig, kafkasource.NewAdapter)
}
