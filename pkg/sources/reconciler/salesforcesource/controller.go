package salesforcesource

import (
	"context"

	"github.com/kelseyhightower/envconfig"

	"knative.dev/eventing/pkg/reconciler/source"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	"github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	informerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/sources/v1alpha1/salesforcesource"
	reconcilerv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/salesforcesource"
	common "github.com/zeiss/typhoon/pkg/reconciler"
)

// NewController creates a Reconciler for the event source and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	typ := (*v1alpha1.SalesforceSource)(nil)
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

	r.base = common.NewGenericDeploymentReconciler[*v1alpha1.SalesforceSource](
		ctx,
		typ.GetGroupVersionKind(),
		impl.Tracker,
		impl.EnqueueControllerOf,
		informer.Lister().SalesforceSources,
	)

	informer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
