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

// FakeFilters implements FilterInterface
type FakeFilters struct {
	Fake *FakeRoutingV1alpha1
	ns   string
}

var filtersResource = v1alpha1.SchemeGroupVersion.WithResource("filters")

var filtersKind = v1alpha1.SchemeGroupVersion.WithKind("Filter")

// Get takes name of the filter, and returns the corresponding filter object, and an error if there is any.
func (c *FakeFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Filter, err error) {
	emptyResult := &v1alpha1.Filter{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(filtersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Filter), err
}

// List takes label and field selectors, and returns the list of Filters that match those selectors.
func (c *FakeFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FilterList, err error) {
	emptyResult := &v1alpha1.FilterList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(filtersResource, filtersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FilterList{ListMeta: obj.(*v1alpha1.FilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.FilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested filters.
func (c *FakeFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(filtersResource, c.ns, opts))

}

// Create takes the representation of a filter and creates it.  Returns the server's representation of the filter, and an error, if there is any.
func (c *FakeFilters) Create(ctx context.Context, filter *v1alpha1.Filter, opts v1.CreateOptions) (result *v1alpha1.Filter, err error) {
	emptyResult := &v1alpha1.Filter{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(filtersResource, c.ns, filter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Filter), err
}

// Update takes the representation of a filter and updates it. Returns the server's representation of the filter, and an error, if there is any.
func (c *FakeFilters) Update(ctx context.Context, filter *v1alpha1.Filter, opts v1.UpdateOptions) (result *v1alpha1.Filter, err error) {
	emptyResult := &v1alpha1.Filter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(filtersResource, c.ns, filter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Filter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFilters) UpdateStatus(ctx context.Context, filter *v1alpha1.Filter, opts v1.UpdateOptions) (result *v1alpha1.Filter, err error) {
	emptyResult := &v1alpha1.Filter{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(filtersResource, "status", c.ns, filter, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Filter), err
}

// Delete takes name of the filter and deletes it. Returns an error if one occurs.
func (c *FakeFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(filtersResource, c.ns, name, opts), &v1alpha1.Filter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(filtersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FilterList{})
	return err
}

// Patch applies the patch and returns the patched filter.
func (c *FakeFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Filter, err error) {
	emptyResult := &v1alpha1.Filter{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(filtersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Filter), err
}
