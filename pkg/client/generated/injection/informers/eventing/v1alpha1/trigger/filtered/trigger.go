// Code generated by injection-gen. DO NOT EDIT.

package filtered

import (
	context "context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/informers/externalversions/eventing/v1alpha1"
	filtered "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/filtered"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterFilteredInformers(withInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct {
	Selector string
}

func withInformer(ctx context.Context) (context.Context, []controller.Informer) {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	infs := []controller.Informer{}
	for _, selector := range labelSelectors {
		f := filtered.Get(ctx, selector)
		inf := f.Eventing().V1alpha1().Triggers()
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
		infs = append(infs, inf.Informer())
	}
	return ctx, infs
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context, selector string) v1alpha1.TriggerInformer {
	untyped := ctx.Value(Key{Selector: selector})
	if untyped == nil {
		logging.FromContext(ctx).Panicf(
			"Unable to fetch github.com/zeiss/typhoon/pkg/client/generated/informers/externalversions/eventing/v1alpha1.TriggerInformer with selector %s from context.", selector)
	}
	return untyped.(v1alpha1.TriggerInformer)
}
