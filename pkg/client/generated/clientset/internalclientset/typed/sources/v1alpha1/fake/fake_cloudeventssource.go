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

// FakeCloudEventsSources implements CloudEventsSourceInterface
type FakeCloudEventsSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var cloudeventssourcesResource = v1alpha1.SchemeGroupVersion.WithResource("cloudeventssources")

var cloudeventssourcesKind = v1alpha1.SchemeGroupVersion.WithKind("CloudEventsSource")

// Get takes name of the cloudEventsSource, and returns the corresponding cloudEventsSource object, and an error if there is any.
func (c *FakeCloudEventsSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudEventsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudeventssourcesResource, c.ns, name), &v1alpha1.CloudEventsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudEventsSource), err
}

// List takes label and field selectors, and returns the list of CloudEventsSources that match those selectors.
func (c *FakeCloudEventsSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudEventsSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudeventssourcesResource, cloudeventssourcesKind, c.ns, opts), &v1alpha1.CloudEventsSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudEventsSourceList{ListMeta: obj.(*v1alpha1.CloudEventsSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudEventsSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudEventsSources.
func (c *FakeCloudEventsSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudeventssourcesResource, c.ns, opts))

}

// Create takes the representation of a cloudEventsSource and creates it.  Returns the server's representation of the cloudEventsSource, and an error, if there is any.
func (c *FakeCloudEventsSources) Create(ctx context.Context, cloudEventsSource *v1alpha1.CloudEventsSource, opts v1.CreateOptions) (result *v1alpha1.CloudEventsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudeventssourcesResource, c.ns, cloudEventsSource), &v1alpha1.CloudEventsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudEventsSource), err
}

// Update takes the representation of a cloudEventsSource and updates it. Returns the server's representation of the cloudEventsSource, and an error, if there is any.
func (c *FakeCloudEventsSources) Update(ctx context.Context, cloudEventsSource *v1alpha1.CloudEventsSource, opts v1.UpdateOptions) (result *v1alpha1.CloudEventsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudeventssourcesResource, c.ns, cloudEventsSource), &v1alpha1.CloudEventsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudEventsSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudEventsSources) UpdateStatus(ctx context.Context, cloudEventsSource *v1alpha1.CloudEventsSource, opts v1.UpdateOptions) (*v1alpha1.CloudEventsSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudeventssourcesResource, "status", c.ns, cloudEventsSource), &v1alpha1.CloudEventsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudEventsSource), err
}

// Delete takes name of the cloudEventsSource and deletes it. Returns an error if one occurs.
func (c *FakeCloudEventsSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudeventssourcesResource, c.ns, name, opts), &v1alpha1.CloudEventsSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudEventsSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudeventssourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudEventsSourceList{})
	return err
}

// Patch applies the patch and returns the patched cloudEventsSource.
func (c *FakeCloudEventsSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudEventsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudeventssourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudEventsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudEventsSource), err
}
