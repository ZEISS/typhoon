package webhooksource

import (
	"context"

	"github.com/kelseyhightower/envconfig"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	informerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/sources/v1alpha1/webhooksource"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/webhooksource"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	typ := (*v1alpha1.WebhookSource)(nil)
	app := common.ComponentName(typ)

	adapterCfg := &adapterConfig{
		configs: source.WatchConfigurations(ctx, app, cmw),
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
		informer.Lister().WebhookSources,
	)

	informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
