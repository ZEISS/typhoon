// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	filtered "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/extensions/v1alpha1/function/filtered"
	factoryfiltered "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/filtered"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

var Get = filtered.Get

func init() {
	injection.Fake.RegisterFilteredInformers(withInformer)
}

func withInformer(ctx context.Context) (context.Context, []controller.Informer) {
	untyped := ctx.Value(factoryfiltered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	infs := []controller.Informer{}
	for _, selector := range labelSelectors {
		f := factoryfiltered.Get(ctx, selector)
		inf := f.Extensions().V1alpha1().Functions()
		ctx = context.WithValue(ctx, filtered.Key{Selector: selector}, inf)
		infs = append(infs, inf.Informer())
	}
	return ctx, infs
}
