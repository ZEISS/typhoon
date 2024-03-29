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

// NatsTargetsGetter has a method to return a NatsTargetInterface.
// A group's client should implement this interface.
type NatsTargetsGetter interface {
	NatsTargets(namespace string) NatsTargetInterface
}

// NatsTargetInterface has methods to work with NatsTarget resources.
type NatsTargetInterface interface {
	Create(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.CreateOptions) (*v1alpha1.NatsTarget, error)
	Update(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (*v1alpha1.NatsTarget, error)
	UpdateStatus(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (*v1alpha1.NatsTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NatsTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NatsTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatsTarget, err error)
	NatsTargetExpansion
}

// natsTargets implements NatsTargetInterface
type natsTargets struct {
	client rest.Interface
	ns     string
}

// newNatsTargets returns a NatsTargets
func newNatsTargets(c *TargetsV1alpha1Client, namespace string) *natsTargets {
	return &natsTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the natsTarget, and returns the corresponding natsTarget object, and an error if there is any.
func (c *natsTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NatsTarget, err error) {
	result = &v1alpha1.NatsTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natstargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NatsTargets that match those selectors.
func (c *natsTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NatsTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NatsTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("natstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested natsTargets.
func (c *natsTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("natstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a natsTarget and creates it.  Returns the server's representation of the natsTarget, and an error, if there is any.
func (c *natsTargets) Create(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.CreateOptions) (result *v1alpha1.NatsTarget, err error) {
	result = &v1alpha1.NatsTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("natstargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(natsTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a natsTarget and updates it. Returns the server's representation of the natsTarget, and an error, if there is any.
func (c *natsTargets) Update(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (result *v1alpha1.NatsTarget, err error) {
	result = &v1alpha1.NatsTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natstargets").
		Name(natsTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(natsTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *natsTargets) UpdateStatus(ctx context.Context, natsTarget *v1alpha1.NatsTarget, opts v1.UpdateOptions) (result *v1alpha1.NatsTarget, err error) {
	result = &v1alpha1.NatsTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("natstargets").
		Name(natsTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(natsTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the natsTarget and deletes it. Returns an error if one occurs.
func (c *natsTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natstargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *natsTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("natstargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched natsTarget.
func (c *natsTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatsTarget, err error) {
	result = &v1alpha1.NatsTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("natstargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
