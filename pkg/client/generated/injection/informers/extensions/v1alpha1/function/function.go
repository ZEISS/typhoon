// Code generated by injection-gen. DO NOT EDIT.

package function

import (
	context "context"

	factory "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory"
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/extensions/v1alpha1"
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
	inf := f.Extensions().V1alpha1().Functions()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.FunctionInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/extensions/v1alpha1.FunctionInformer from context.")
	}
	return untyped.(v1alpha1.FunctionInformer)
}
