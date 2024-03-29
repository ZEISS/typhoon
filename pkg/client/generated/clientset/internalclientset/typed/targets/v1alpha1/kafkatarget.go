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

// KafkaTargetsGetter has a method to return a KafkaTargetInterface.
// A group's client should implement this interface.
type KafkaTargetsGetter interface {
	KafkaTargets(namespace string) KafkaTargetInterface
}

// KafkaTargetInterface has methods to work with KafkaTarget resources.
type KafkaTargetInterface interface {
	Create(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.CreateOptions) (*v1alpha1.KafkaTarget, error)
	Update(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.UpdateOptions) (*v1alpha1.KafkaTarget, error)
	UpdateStatus(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.UpdateOptions) (*v1alpha1.KafkaTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KafkaTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KafkaTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KafkaTarget, err error)
	KafkaTargetExpansion
}

// kafkaTargets implements KafkaTargetInterface
type kafkaTargets struct {
	client rest.Interface
	ns     string
}

// newKafkaTargets returns a KafkaTargets
func newKafkaTargets(c *TargetsV1alpha1Client, namespace string) *kafkaTargets {
	return &kafkaTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaTarget, and returns the corresponding kafkaTarget object, and an error if there is any.
func (c *kafkaTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KafkaTarget, err error) {
	result = &v1alpha1.KafkaTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkatargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaTargets that match those selectors.
func (c *kafkaTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KafkaTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KafkaTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkatargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaTargets.
func (c *kafkaTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkatargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kafkaTarget and creates it.  Returns the server's representation of the kafkaTarget, and an error, if there is any.
func (c *kafkaTargets) Create(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.CreateOptions) (result *v1alpha1.KafkaTarget, err error) {
	result = &v1alpha1.KafkaTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkatargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kafkaTarget and updates it. Returns the server's representation of the kafkaTarget, and an error, if there is any.
func (c *kafkaTargets) Update(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.UpdateOptions) (result *v1alpha1.KafkaTarget, err error) {
	result = &v1alpha1.KafkaTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkatargets").
		Name(kafkaTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kafkaTargets) UpdateStatus(ctx context.Context, kafkaTarget *v1alpha1.KafkaTarget, opts v1.UpdateOptions) (result *v1alpha1.KafkaTarget, err error) {
	result = &v1alpha1.KafkaTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkatargets").
		Name(kafkaTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kafkaTarget and deletes it. Returns an error if one occurs.
func (c *kafkaTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkatargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkatargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kafkaTarget.
func (c *kafkaTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KafkaTarget, err error) {
	result = &v1alpha1.KafkaTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkatargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
