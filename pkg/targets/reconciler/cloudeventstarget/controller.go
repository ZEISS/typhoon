package cloudeventstarget

import (
	"context"

	"github.com/kelseyhightower/envconfig"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	"github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	informerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/targets/v1alpha1/cloudeventstarget"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/targets/v1alpha1/cloudeventstarget"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	typ := (*v1alpha1.CloudEventsTarget)(nil)
	app := common.ComponentName(typ)

	// Calling envconfig.Process() with a prefix appends that prefix
	// (uppercased) to the Go field name, e.g. MYTARGET_IMAGE.
	adapterCfg := &adapterConfig{
		obsConfig: source.WatchConfigurations(ctx, app, cmw),
	}
	envconfig.MustProcess(app, adapterCfg)

	informer := informerv1alpha1.Get(ctx)

	r := &Reconciler{
		adapterCfg: adapterCfg,
	}
	impl := reconcilerv1alpha1.NewImpl(ctx, r)

	r.base = common.NewGenericServiceReconciler(
		ctx,
		typ.GetGroupVersionKind(),
		impl.Tracker,
		impl.EnqueueControllerOf,
		informer.Lister().CloudEventsTargets,
	)

	informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
