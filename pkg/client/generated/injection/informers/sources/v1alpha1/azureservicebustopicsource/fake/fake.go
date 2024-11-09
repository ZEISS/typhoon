// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/factory/fake"
	azureservicebustopicsource "github.com/zeiss/typhoon/pkg/client/generated/injection/informers/sources/v1alpha1/azureservicebustopicsource"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = azureservicebustopicsource.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Sources().V1alpha1().AzureServiceBusTopicSources()
	return context.WithValue(ctx, azureservicebustopicsource.Key{}, inf), inf.Informer()
}
