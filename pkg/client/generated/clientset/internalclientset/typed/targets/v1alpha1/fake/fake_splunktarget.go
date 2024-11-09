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

// FakeSplunkTargets implements SplunkTargetInterface
type FakeSplunkTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var splunktargetsResource = v1alpha1.SchemeGroupVersion.WithResource("splunktargets")

var splunktargetsKind = v1alpha1.SchemeGroupVersion.WithKind("SplunkTarget")

// Get takes name of the splunkTarget, and returns the corresponding splunkTarget object, and an error if there is any.
func (c *FakeSplunkTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SplunkTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(splunktargetsResource, c.ns, name), &v1alpha1.SplunkTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SplunkTarget), err
}

// List takes label and field selectors, and returns the list of SplunkTargets that match those selectors.
func (c *FakeSplunkTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SplunkTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(splunktargetsResource, splunktargetsKind, c.ns, opts), &v1alpha1.SplunkTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SplunkTargetList{ListMeta: obj.(*v1alpha1.SplunkTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.SplunkTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested splunkTargets.
func (c *FakeSplunkTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(splunktargetsResource, c.ns, opts))

}

// Create takes the representation of a splunkTarget and creates it.  Returns the server's representation of the splunkTarget, and an error, if there is any.
func (c *FakeSplunkTargets) Create(ctx context.Context, splunkTarget *v1alpha1.SplunkTarget, opts v1.CreateOptions) (result *v1alpha1.SplunkTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(splunktargetsResource, c.ns, splunkTarget), &v1alpha1.SplunkTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SplunkTarget), err
}

// Update takes the representation of a splunkTarget and updates it. Returns the server's representation of the splunkTarget, and an error, if there is any.
func (c *FakeSplunkTargets) Update(ctx context.Context, splunkTarget *v1alpha1.SplunkTarget, opts v1.UpdateOptions) (result *v1alpha1.SplunkTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(splunktargetsResource, c.ns, splunkTarget), &v1alpha1.SplunkTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SplunkTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSplunkTargets) UpdateStatus(ctx context.Context, splunkTarget *v1alpha1.SplunkTarget, opts v1.UpdateOptions) (*v1alpha1.SplunkTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(splunktargetsResource, "status", c.ns, splunkTarget), &v1alpha1.SplunkTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SplunkTarget), err
}

// Delete takes name of the splunkTarget and deletes it. Returns an error if one occurs.
func (c *FakeSplunkTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(splunktargetsResource, c.ns, name, opts), &v1alpha1.SplunkTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSplunkTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(splunktargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SplunkTargetList{})
	return err
}

// Patch applies the patch and returns the patched splunkTarget.
func (c *FakeSplunkTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SplunkTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(splunktargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SplunkTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SplunkTarget), err
}
