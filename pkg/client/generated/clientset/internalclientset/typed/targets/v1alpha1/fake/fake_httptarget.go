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

// FakeHTTPTargets implements HTTPTargetInterface
type FakeHTTPTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var httptargetsResource = v1alpha1.SchemeGroupVersion.WithResource("httptargets")

var httptargetsKind = v1alpha1.SchemeGroupVersion.WithKind("HTTPTarget")

// Get takes name of the hTTPTarget, and returns the corresponding hTTPTarget object, and an error if there is any.
func (c *FakeHTTPTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HTTPTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httptargetsResource, c.ns, name), &v1alpha1.HTTPTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPTarget), err
}

// List takes label and field selectors, and returns the list of HTTPTargets that match those selectors.
func (c *FakeHTTPTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HTTPTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httptargetsResource, httptargetsKind, c.ns, opts), &v1alpha1.HTTPTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HTTPTargetList{ListMeta: obj.(*v1alpha1.HTTPTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.HTTPTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPTargets.
func (c *FakeHTTPTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httptargetsResource, c.ns, opts))

}

// Create takes the representation of a hTTPTarget and creates it.  Returns the server's representation of the hTTPTarget, and an error, if there is any.
func (c *FakeHTTPTargets) Create(ctx context.Context, hTTPTarget *v1alpha1.HTTPTarget, opts v1.CreateOptions) (result *v1alpha1.HTTPTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httptargetsResource, c.ns, hTTPTarget), &v1alpha1.HTTPTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPTarget), err
}

// Update takes the representation of a hTTPTarget and updates it. Returns the server's representation of the hTTPTarget, and an error, if there is any.
func (c *FakeHTTPTargets) Update(ctx context.Context, hTTPTarget *v1alpha1.HTTPTarget, opts v1.UpdateOptions) (result *v1alpha1.HTTPTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httptargetsResource, c.ns, hTTPTarget), &v1alpha1.HTTPTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHTTPTargets) UpdateStatus(ctx context.Context, hTTPTarget *v1alpha1.HTTPTarget, opts v1.UpdateOptions) (*v1alpha1.HTTPTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(httptargetsResource, "status", c.ns, hTTPTarget), &v1alpha1.HTTPTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPTarget), err
}

// Delete takes name of the hTTPTarget and deletes it. Returns an error if one occurs.
func (c *FakeHTTPTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(httptargetsResource, c.ns, name, opts), &v1alpha1.HTTPTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httptargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HTTPTargetList{})
	return err
}

// Patch applies the patch and returns the patched hTTPTarget.
func (c *FakeHTTPTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httptargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HTTPTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPTarget), err
}
