// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/fake"
	jqtransformation "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/flow/v1alpha1/jqtransformation"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = jqtransformation.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Flow().V1alpha1().JQTransformations()
	return context.WithValue(ctx, jqtransformation.Key{}, inf), inf.Informer()
}
