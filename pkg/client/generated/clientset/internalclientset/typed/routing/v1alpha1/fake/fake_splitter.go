// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSplitters implements SplitterInterface
type FakeSplitters struct {
	Fake *FakeRoutingV1alpha1
	ns   string
}

var splittersResource = v1alpha1.SchemeGroupVersion.WithResource("splitters")

var splittersKind = v1alpha1.SchemeGroupVersion.WithKind("Splitter")

// Get takes name of the splitter, and returns the corresponding splitter object, and an error if there is any.
func (c *FakeSplitters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Splitter, err error) {
	emptyResult := &v1alpha1.Splitter{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(splittersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Splitter), err
}

// List takes label and field selectors, and returns the list of Splitters that match those selectors.
func (c *FakeSplitters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SplitterList, err error) {
	emptyResult := &v1alpha1.SplitterList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(splittersResource, splittersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SplitterList{ListMeta: obj.(*v1alpha1.SplitterList).ListMeta}
	for _, item := range obj.(*v1alpha1.SplitterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested splitters.
func (c *FakeSplitters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(splittersResource, c.ns, opts))

}

// Create takes the representation of a splitter and creates it.  Returns the server's representation of the splitter, and an error, if there is any.
func (c *FakeSplitters) Create(ctx context.Context, splitter *v1alpha1.Splitter, opts v1.CreateOptions) (result *v1alpha1.Splitter, err error) {
	emptyResult := &v1alpha1.Splitter{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(splittersResource, c.ns, splitter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Splitter), err
}

// Update takes the representation of a splitter and updates it. Returns the server's representation of the splitter, and an error, if there is any.
func (c *FakeSplitters) Update(ctx context.Context, splitter *v1alpha1.Splitter, opts v1.UpdateOptions) (result *v1alpha1.Splitter, err error) {
	emptyResult := &v1alpha1.Splitter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(splittersResource, c.ns, splitter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Splitter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSplitters) UpdateStatus(ctx context.Context, splitter *v1alpha1.Splitter, opts v1.UpdateOptions) (result *v1alpha1.Splitter, err error) {
	emptyResult := &v1alpha1.Splitter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(splittersResource, "status", c.ns, splitter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Splitter), err
}

// Delete takes name of the splitter and deletes it. Returns an error if one occurs.
func (c *FakeSplitters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(splittersResource, c.ns, name, opts), &v1alpha1.Splitter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSplitters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(splittersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SplitterList{})
	return err
}

// Patch applies the patch and returns the patched splitter.
func (c *FakeSplitters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Splitter, err error) {
	emptyResult := &v1alpha1.Splitter{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(splittersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Splitter), err
}
