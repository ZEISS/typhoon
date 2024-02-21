package sharedmain

import (
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/signals"

	"github.com/zeiss/typhoon/pkg/sources/adapter/common/env"
)

type (
	namedControllerConstructor func(component string) adapter.ControllerConstructor
	namedAdapterConstructor    func(component string) adapter.AdapterConstructor
)

// MainWithController is a shared main tailored to multi-tenant receive-adapters.
// It performs the following initializations:
//   - process environment variables
//   - enable leader election / HA
//   - set the scope to a single namespace
//   - inject the given controller constructor
func MainWithController(envCtor env.ConfigConstructor,
	cCtor namedControllerConstructor, aCtor namedAdapterConstructor,
) {
	envAcc := env.MustProcessConfig(envCtor)
	ns := envAcc.GetNamespace()
	component := envAcc.GetComponent()

	ctx := signals.NewContext()
	ctx = adapter.WithHAEnabled(ctx)
	ctx = injection.WithNamespaceScope(ctx, ns)
	ctx = adapter.WithController(ctx, cCtor(component))

	adapter.MainWithEnv(ctx, component, envAcc, aCtor(component))
}
