// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/fake"
	cloudeventstarget "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/targets/v1alpha1/cloudeventstarget"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = cloudeventstarget.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Targets().V1alpha1().CloudEventsTargets()
	return context.WithValue(ctx, cloudeventstarget.Key{}, inf), inf.Informer()
}
