// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureServiceBusSources implements AzureServiceBusSourceInterface
type FakeAzureServiceBusSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var azureservicebussourcesResource = v1alpha1.SchemeGroupVersion.WithResource("azureservicebussources")

var azureservicebussourcesKind = v1alpha1.SchemeGroupVersion.WithKind("AzureServiceBusSource")

// Get takes name of the azureServiceBusSource, and returns the corresponding azureServiceBusSource object, and an error if there is any.
func (c *FakeAzureServiceBusSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSource{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(azureservicebussourcesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AzureServiceBusSource), err
}

// List takes label and field selectors, and returns the list of AzureServiceBusSources that match those selectors.
func (c *FakeAzureServiceBusSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureServiceBusSourceList, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSourceList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(azureservicebussourcesResource, azureservicebussourcesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureServiceBusSourceList{ListMeta: obj.(*v1alpha1.AzureServiceBusSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureServiceBusSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureServiceBusSources.
func (c *FakeAzureServiceBusSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(azureservicebussourcesResource, c.ns, opts))

}

// Create takes the representation of a azureServiceBusSource and creates it.  Returns the server's representation of the azureServiceBusSource, and an error, if there is any.
func (c *FakeAzureServiceBusSources) Create(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.CreateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSource{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(azureservicebussourcesResource, c.ns, azureServiceBusSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AzureServiceBusSource), err
}

// Update takes the representation of a azureServiceBusSource and updates it. Returns the server's representation of the azureServiceBusSource, and an error, if there is any.
func (c *FakeAzureServiceBusSources) Update(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(azureservicebussourcesResource, c.ns, azureServiceBusSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AzureServiceBusSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureServiceBusSources) UpdateStatus(ctx context.Context, azureServiceBusSource *v1alpha1.AzureServiceBusSource, opts v1.UpdateOptions) (result *v1alpha1.AzureServiceBusSource, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSource{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(azureservicebussourcesResource, "status", c.ns, azureServiceBusSource, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AzureServiceBusSource), err
}

// Delete takes name of the azureServiceBusSource and deletes it. Returns an error if one occurs.
func (c *FakeAzureServiceBusSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(azureservicebussourcesResource, c.ns, name, opts), &v1alpha1.AzureServiceBusSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureServiceBusSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(azureservicebussourcesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureServiceBusSourceList{})
	return err
}

// Patch applies the patch and returns the patched azureServiceBusSource.
func (c *FakeAzureServiceBusSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServiceBusSource, err error) {
	emptyResult := &v1alpha1.AzureServiceBusSource{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(azureservicebussourcesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.AzureServiceBusSource), err
}
