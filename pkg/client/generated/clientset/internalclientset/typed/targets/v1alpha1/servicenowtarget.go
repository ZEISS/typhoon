// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceNowTargetsGetter has a method to return a ServiceNowTargetInterface.
// A group's client should implement this interface.
type ServiceNowTargetsGetter interface {
	ServiceNowTargets(namespace string) ServiceNowTargetInterface
}

// ServiceNowTargetInterface has methods to work with ServiceNowTarget resources.
type ServiceNowTargetInterface interface {
	Create(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.CreateOptions) (*v1alpha1.ServiceNowTarget, error)
	Update(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.UpdateOptions) (*v1alpha1.ServiceNowTarget, error)
	UpdateStatus(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.UpdateOptions) (*v1alpha1.ServiceNowTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceNowTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceNowTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNowTarget, err error)
	ServiceNowTargetExpansion
}

// serviceNowTargets implements ServiceNowTargetInterface
type serviceNowTargets struct {
	client rest.Interface
	ns     string
}

// newServiceNowTargets returns a ServiceNowTargets
func newServiceNowTargets(c *TargetsV1alpha1Client, namespace string) *serviceNowTargets {
	return &serviceNowTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceNowTarget, and returns the corresponding serviceNowTarget object, and an error if there is any.
func (c *serviceNowTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceNowTarget, err error) {
	result = &v1alpha1.ServiceNowTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenowtargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceNowTargets that match those selectors.
func (c *serviceNowTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceNowTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceNowTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenowtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceNowTargets.
func (c *serviceNowTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicenowtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceNowTarget and creates it.  Returns the server's representation of the serviceNowTarget, and an error, if there is any.
func (c *serviceNowTargets) Create(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.CreateOptions) (result *v1alpha1.ServiceNowTarget, err error) {
	result = &v1alpha1.ServiceNowTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicenowtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNowTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceNowTarget and updates it. Returns the server's representation of the serviceNowTarget, and an error, if there is any.
func (c *serviceNowTargets) Update(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.UpdateOptions) (result *v1alpha1.ServiceNowTarget, err error) {
	result = &v1alpha1.ServiceNowTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenowtargets").
		Name(serviceNowTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNowTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceNowTargets) UpdateStatus(ctx context.Context, serviceNowTarget *v1alpha1.ServiceNowTarget, opts v1.UpdateOptions) (result *v1alpha1.ServiceNowTarget, err error) {
	result = &v1alpha1.ServiceNowTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenowtargets").
		Name(serviceNowTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNowTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceNowTarget and deletes it. Returns an error if one occurs.
func (c *serviceNowTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenowtargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceNowTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenowtargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceNowTarget.
func (c *serviceNowTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNowTarget, err error) {
	result = &v1alpha1.ServiceNowTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicenowtargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}