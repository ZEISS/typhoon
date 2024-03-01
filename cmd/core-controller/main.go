package main

import (
	"os"

	injection "knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"

	"github.com/zeiss/typhoon/pkg/reconciler/redisbroker"
	"github.com/zeiss/typhoon/pkg/reconciler/trigger"
)

func main() {
	ctx := signals.NewContext()

	ns := os.Getenv("WORKING_NAMESPACE")
	if len(ns) != 0 {
		ctx = injection.WithNamespaceScope(ctx, ns)
	}

	sharedmain.MainWithContext(ctx, "core-controller",
		redisbroker.NewController,
		trigger.NewController,
	)
}
