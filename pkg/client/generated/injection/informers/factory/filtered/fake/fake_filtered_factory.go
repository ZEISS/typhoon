// Code generated by injection-gen. DO NOT EDIT.

package fakeFilteredFactory

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/client/fake"
	filtered "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/filtered"
	externalversions "github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

var Get = filtered.Get

func init() {
	injection.Fake.RegisterInformerFactory(withInformerFactory)
}

func withInformerFactory(ctx context.Context) context.Context {
	c := fake.Get(ctx)
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	for _, selector := range labelSelectors {
		selectorVal := selector
		opts := []externalversions.SharedInformerOption{}
		if injection.HasNamespaceScope(ctx) {
			opts = append(opts, externalversions.WithNamespace(injection.GetNamespaceScope(ctx)))
		}
		opts = append(opts, externalversions.WithTweakListOptions(func(l *v1.ListOptions) {
			l.LabelSelector = selectorVal
		}))
		ctx = context.WithValue(ctx, filtered.Key{Selector: selectorVal},
			externalversions.NewSharedInformerFactoryWithOptions(c, controller.GetResyncPeriod(ctx), opts...))
	}
	return ctx
}
