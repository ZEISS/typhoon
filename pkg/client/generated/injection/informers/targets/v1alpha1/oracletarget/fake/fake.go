// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/fake"
	oracletarget "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/targets/v1alpha1/oracletarget"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = oracletarget.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Targets().V1alpha1().OracleTargets()
	return context.WithValue(ctx, oracletarget.Key{}, inf), inf.Informer()
}