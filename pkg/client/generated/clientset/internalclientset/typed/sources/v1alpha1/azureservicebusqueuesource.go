// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AzureServiceBusQueueSourcesGetter has a method to return a AzureServiceBusQueueSourceInterface.
// A group's client should implement this interface.
type AzureServiceBusQueueSourcesGetter interface {
	AzureServiceBusQueueSources(namespace string) AzureServiceBusQueueSourceInterface
}

// AzureServiceBusQueueSourceInterface has methods to work with AzureServiceBusQueueSource resources.
type AzureServiceBusQueueSourceInterface interface {
	Create(ctx context.Context, azureServiceBusQueueSource *v1alpha1.AzureServiceBusQueueSource, opts v1.CreateOptions) (*v1alpha1.AzureServiceBusQueueSource, error)
	Update(ctx context.Context, azureServiceBusQueueSource *v1alpha1.AzureServiceBusQueueSource, opts v1.UpdateOptions) (*v1alpha1.AzureServiceBusQueueSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, azureServiceBusQueueSource *v1alpha1.AzureServiceBusQueueSource, opts v1.UpdateOptions) (*v1alpha1.AzureServiceBusQueueSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzureServiceBusQueueSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzureServiceBusQueueSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServiceBusQueueSource, err error)
	AzureServiceBusQueueSourceExpansion
}

// azureServiceBusQueueSources implements AzureServiceBusQueueSourceInterface
type azureServiceBusQueueSources struct {
	*gentype.ClientWithList[*v1alpha1.AzureServiceBusQueueSource, *v1alpha1.AzureServiceBusQueueSourceList]
}

// newAzureServiceBusQueueSources returns a AzureServiceBusQueueSources
func newAzureServiceBusQueueSources(c *SourcesV1alpha1Client, namespace string) *azureServiceBusQueueSources {
	return &azureServiceBusQueueSources{
		gentype.NewClientWithList[*v1alpha1.AzureServiceBusQueueSource, *v1alpha1.AzureServiceBusQueueSourceList](
			"azureservicebusqueuesources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.AzureServiceBusQueueSource { return &v1alpha1.AzureServiceBusQueueSource{} },
			func() *v1alpha1.AzureServiceBusQueueSourceList { return &v1alpha1.AzureServiceBusQueueSourceList{} }),
	}
}
