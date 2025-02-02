// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/fake"
	azureservicebussource "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/sources/v1alpha1/azureservicebussource"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = azureservicebussource.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Sources().V1alpha1().AzureServiceBusSources()
	return context.WithValue(ctx, azureservicebussource.Key{}, inf), inf.Informer()
}
