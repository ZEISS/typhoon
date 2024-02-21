package filter

import (
	"context"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	pkgcontroller "knative.dev/pkg/controller"

	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/routing/v1alpha1/filter"
	"github.com/zeiss/typhoon/pkg/routing/adapter/common/controller"
)

// NewController returns a constructor for the Router's Reconciler.
//
// NOTE(antoineco): although the returned controller doesn't do anything, it is
// necessary to return a valid implementation in order to trigger the Informer
// injection in Knative's sharedmain.Main.
func NewController(component string) pkgadapter.ControllerConstructor {
	return func(ctx context.Context, _ pkgadapter.Adapter) *pkgcontroller.Impl {
		r := (*Reconciler)(nil)
		impl := reconcilerv1alpha1.NewImpl(ctx, r, controller.Opts(component))

		return impl
	}
}
