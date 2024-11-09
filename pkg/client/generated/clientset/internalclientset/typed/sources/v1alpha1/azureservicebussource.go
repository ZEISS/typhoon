// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureServiceBusSourcesGetter has a method to return a AzureServiceBusSourceInterface.
// A group's client should implement this interface.
type AzureServiceBusSourcesGetter interface {
	AzureServiceBusSources(namespace string) AzureServiceBusSourceInterface
}

// AzureServiceBusSourceInterface has methods to work with AzureServiceBusSource resources.
type AzureServiceBusSourceInterface interface {
	Create(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.CreateOptions) (*v1alpha1.AzureServiceBusSource, error)
	Update(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (*v1alpha1.AzureServiceBusSource, error)
	UpdateStatus(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (*v1alpha1.AzureServiceBusSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzureServiceBusSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzureServiceBusSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServiceBusSource, err error)
	AzureServiceBusSourceExpansion
}

// azureServiceBusSources implements AzureServiceBusSourceInterface
type azureServiceBusSources struct {
	client rest.Interface
	ns     string
}

// newAzureServiceBusSources returns a AzureServiceBusSources
func newAzureServiceBusSources(c *SourcesV1alpha1Client, namespace string) *azureServiceBusSources {
	return &azureServiceBusSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureServiceBusSource, and returns the corresponding azureServiceBusSource object, and an error if there is any.
func (c *azureServiceBusSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	result = &v1alpha1.AzureServiceBusSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureservicebussources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureServiceBusSources that match those selectors.
func (c *azureServiceBusSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureServiceBusSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureServiceBusSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureservicebussources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureServiceBusSources.
func (c *azureServiceBusSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureservicebussources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureServiceBusSource and creates it.  Returns the server's representation of the azureServiceBusSource, and an error, if there is any.
func (c *azureServiceBusSources) Create(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.CreateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	result = &v1alpha1.AzureServiceBusSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureservicebussources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServiceBusSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureServiceBusSource and updates it. Returns the server's representation of the azureServiceBusSource, and an error, if there is any.
func (c *azureServiceBusSources) Update(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	result = &v1alpha1.AzureServiceBusSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureservicebussources").
		Name(azureServiceBusSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServiceBusSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *azureServiceBusSources) UpdateStatus(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	result = &v1alpha1.AzureServiceBusSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureservicebussources").
		Name(azureServiceBusSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServiceBusSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureServiceBusSource and deletes it. Returns an error if one occurs.
func (c *azureServiceBusSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureservicebussources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureServiceBusSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureservicebussources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureServiceBusSource.
func (c *azureServiceBusSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServiceBusSource, err error) {
	result = &v1alpha1.AzureServiceBusSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureservicebussources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
