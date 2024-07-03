// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatadogTargets implements DatadogTargetInterface
type FakeDatadogTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var datadogtargetsResource = v1alpha1.SchemeGroupVersion.WithResource("datadogtargets")

var datadogtargetsKind = v1alpha1.SchemeGroupVersion.WithKind("DatadogTarget")

// Get takes name of the datadogTarget, and returns the corresponding datadogTarget object, and an error if there is any.
func (c *FakeDatadogTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatadogTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datadogtargetsResource, c.ns, name), &v1alpha1.DatadogTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatadogTarget), err
}

// List takes label and field selectors, and returns the list of DatadogTargets that match those selectors.
func (c *FakeDatadogTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatadogTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datadogtargetsResource, datadogtargetsKind, c.ns, opts), &v1alpha1.DatadogTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatadogTargetList{ListMeta: obj.(*v1alpha1.DatadogTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatadogTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datadogTargets.
func (c *FakeDatadogTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datadogtargetsResource, c.ns, opts))

}

// Create takes the representation of a datadogTarget and creates it.  Returns the server's representation of the datadogTarget, and an error, if there is any.
func (c *FakeDatadogTargets) Create(ctx context.Context, datadogTarget *v1alpha1.DatadogTarget, opts v1.CreateOptions) (result *v1alpha1.DatadogTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datadogtargetsResource, c.ns, datadogTarget), &v1alpha1.DatadogTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatadogTarget), err
}

// Update takes the representation of a datadogTarget and updates it. Returns the server's representation of the datadogTarget, and an error, if there is any.
func (c *FakeDatadogTargets) Update(ctx context.Context, datadogTarget *v1alpha1.DatadogTarget, opts v1.UpdateOptions) (result *v1alpha1.DatadogTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datadogtargetsResource, c.ns, datadogTarget), &v1alpha1.DatadogTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatadogTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatadogTargets) UpdateStatus(ctx context.Context, datadogTarget *v1alpha1.DatadogTarget, opts v1.UpdateOptions) (*v1alpha1.DatadogTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datadogtargetsResource, "status", c.ns, datadogTarget), &v1alpha1.DatadogTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatadogTarget), err
}

// Delete takes name of the datadogTarget and deletes it. Returns an error if one occurs.
func (c *FakeDatadogTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(datadogtargetsResource, c.ns, name, opts), &v1alpha1.DatadogTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatadogTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datadogtargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatadogTargetList{})
	return err
}

// Patch applies the patch and returns the patched datadogTarget.
func (c *FakeDatadogTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatadogTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datadogtargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatadogTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatadogTarget), err
}