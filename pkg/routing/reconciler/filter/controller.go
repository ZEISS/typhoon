package filter

import (
	"context"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	"github.com/kelseyhightower/envconfig"
	"github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	informerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/routing/v1alpha1/filter"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/routing/v1alpha1/filter"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// NewController creates a Reconciler and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	typ := (*v1alpha1.Filter)(nil)
	app := common.ComponentName(typ)

	// Calling envconfig.Process() with a prefix appends that prefix
	// (uppercased) to the Go field name, e.g. MYSOURCE_IMAGE.
	adapterCfg := &adapterConfig{
		configs: source.WatchConfigurations(ctx, app, cmw),
	}
	envconfig.MustProcess(app, adapterCfg)

	informer := informerv1alpha1.Get(ctx)

	r := &Reconciler{
		adapterCfg: adapterCfg,
	}
	impl := reconcilerv1alpha1.NewImpl(ctx, r)

	logger := logging.FromContext(ctx)

	r.base = common.NewMTGenericServiceReconciler[*v1alpha1.Filter](
		ctx,
		typ,
		impl.Tracker,
		common.EnqueueObjectsInNamespaceOf(informer.Informer(), impl.FilteredGlobalResync, logger),
		informer.Lister().Filters,
	)

	informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
