// Code generated by injection-gen. DO NOT EDIT.

package servicenowtarget

import (
	context "context"

	factory "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory"
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/targets/v1alpha1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Targets().V1alpha1().ServiceNowTargets()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.ServiceNowTargetInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/targets/v1alpha1.ServiceNowTargetInformer from context.")
	}
	return untyped.(v1alpha1.ServiceNowTargetInformer)
}
