package webhooksource

import (
	"context"

	"k8s.io/client-go/tools/cache"

	sourcesalphav1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	webhookinformer "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/sources/v1alpha1/webhooksource"
	webhookreconciler "github.com/zeiss/typhoon/pkg/client/generated/injection/reconciler/sources/v1alpha1/webhooksource"
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

// MTAdapter is the interface the multi-tenant webhook adapter must implement.
type MTAdapter interface {
	// Update is called when the source is ready and when the specification and/or status has changed.
	Update(ctx context.Context, source *sourcesalphav1.WebhookSource)

	// Remove is called when the source has been deleted.
	Remove(source *sourcesalphav1.WebhookSource)

	// RemoveAll is called when the adapter stopped leading
	RemoveAll(ctx context.Context)
}

// NewController ...
func NewController(ctx context.Context, adapter adapter.Adapter) *controller.Impl {
	mtadapter, ok := adapter.(MTAdapter)
	if !ok {
		logging.FromContext(ctx).Fatal("Multi-tenant adapters must implement the MTAdapter interface")
	}

	r := &Reconciler{mtadapter}

	impl := webhookreconciler.NewImpl(ctx, r, func(impl *controller.Impl) controller.Options {
		return controller.Options{
			SkipStatusUpdates: true,
			DemoteFunc: func(b reconciler.Bucket) {
				mtadapter.RemoveAll(ctx)
			},
		}
	})

	webhookinformer.Get(ctx).Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    impl.Enqueue,
			UpdateFunc: controller.PassNew(impl.Enqueue),
			DeleteFunc: r.deleteFunc,
		})
	return impl
}
