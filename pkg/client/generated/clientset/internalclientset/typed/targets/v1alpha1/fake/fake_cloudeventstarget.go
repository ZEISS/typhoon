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

// FakeCloudEventsTargets implements CloudEventsTargetInterface
type FakeCloudEventsTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var cloudeventstargetsResource = v1alpha1.SchemeGroupVersion.WithResource("cloudeventstargets")

var cloudeventstargetsKind = v1alpha1.SchemeGroupVersion.WithKind("CloudEventsTarget")

// Get takes name of the cloudEventsTarget, and returns the corresponding cloudEventsTarget object, and an error if there is any.
func (c *FakeCloudEventsTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudEventsTarget, err error) {
	emptyResult := &v1alpha1.CloudEventsTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(cloudeventstargetsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CloudEventsTarget), err
}

// List takes label and field selectors, and returns the list of CloudEventsTargets that match those selectors.
func (c *FakeCloudEventsTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudEventsTargetList, err error) {
	emptyResult := &v1alpha1.CloudEventsTargetList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(cloudeventstargetsResource, cloudeventstargetsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudEventsTargetList{ListMeta: obj.(*v1alpha1.CloudEventsTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudEventsTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudEventsTargets.
func (c *FakeCloudEventsTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(cloudeventstargetsResource, c.ns, opts))

}

// Create takes the representation of a cloudEventsTarget and creates it.  Returns the server's representation of the cloudEventsTarget, and an error, if there is any.
func (c *FakeCloudEventsTargets) Create(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.CreateOptions) (result *v1alpha1.CloudEventsTarget, err error) {
	emptyResult := &v1alpha1.CloudEventsTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(cloudeventstargetsResource, c.ns, cloudEventsTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CloudEventsTarget), err
}

// Update takes the representation of a cloudEventsTarget and updates it. Returns the server's representation of the cloudEventsTarget, and an error, if there is any.
func (c *FakeCloudEventsTargets) Update(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.UpdateOptions) (result *v1alpha1.CloudEventsTarget, err error) {
	emptyResult := &v1alpha1.CloudEventsTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(cloudeventstargetsResource, c.ns, cloudEventsTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CloudEventsTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudEventsTargets) UpdateStatus(ctx context.Context, cloudEventsTarget *v1alpha1.CloudEventsTarget, opts v1.UpdateOptions) (result *v1alpha1.CloudEventsTarget, err error) {
	emptyResult := &v1alpha1.CloudEventsTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(cloudeventstargetsResource, "status", c.ns, cloudEventsTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CloudEventsTarget), err
}

// Delete takes name of the cloudEventsTarget and deletes it. Returns an error if one occurs.
func (c *FakeCloudEventsTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudeventstargetsResource, c.ns, name, opts), &v1alpha1.CloudEventsTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudEventsTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(cloudeventstargetsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudEventsTargetList{})
	return err
}

// Patch applies the patch and returns the patched cloudEventsTarget.
func (c *FakeCloudEventsTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudEventsTarget, err error) {
	emptyResult := &v1alpha1.CloudEventsTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cloudeventstargetsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.CloudEventsTarget), err
}
