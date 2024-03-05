package bridge

import (
	"context"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	"github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	informerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/flow/v1alpha1/bridge"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/flow/v1alpha1/bridge"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	typ := (*v1alpha1.Bridge)(nil)
	informer := informerv1alpha1.Get(ctx)

	r := &Reconciler{}
	impl := reconcilerv1alpha1.NewImpl(ctx, r)

	r.base = common.NewGenericServiceReconciler[*v1alpha1.Bridge](
		ctx,
		typ.GetGroupVersionKind(),
		impl.Tracker,
		impl.EnqueueControllerOf,
		informer.Lister().Bridges,
	)

	informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
