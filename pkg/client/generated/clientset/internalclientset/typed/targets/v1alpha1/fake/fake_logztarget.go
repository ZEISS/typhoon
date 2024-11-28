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

// FakeLogzTargets implements LogzTargetInterface
type FakeLogzTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var logztargetsResource = v1alpha1.SchemeGroupVersion.WithResource("logztargets")

var logztargetsKind = v1alpha1.SchemeGroupVersion.WithKind("LogzTarget")

// Get takes name of the logzTarget, and returns the corresponding logzTarget object, and an error if there is any.
func (c *FakeLogzTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogzTarget, err error) {
	emptyResult := &v1alpha1.LogzTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(logztargetsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LogzTarget), err
}

// List takes label and field selectors, and returns the list of LogzTargets that match those selectors.
func (c *FakeLogzTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogzTargetList, err error) {
	emptyResult := &v1alpha1.LogzTargetList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(logztargetsResource, logztargetsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogzTargetList{ListMeta: obj.(*v1alpha1.LogzTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogzTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logzTargets.
func (c *FakeLogzTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(logztargetsResource, c.ns, opts))

}

// Create takes the representation of a logzTarget and creates it.  Returns the server's representation of the logzTarget, and an error, if there is any.
func (c *FakeLogzTargets) Create(ctx context.Context, logzTarget *v1alpha1.LogzTarget, opts v1.CreateOptions) (result *v1alpha1.LogzTarget, err error) {
	emptyResult := &v1alpha1.LogzTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(logztargetsResource, c.ns, logzTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LogzTarget), err
}

// Update takes the representation of a logzTarget and updates it. Returns the server's representation of the logzTarget, and an error, if there is any.
func (c *FakeLogzTargets) Update(ctx context.Context, logzTarget *v1alpha1.LogzTarget, opts v1.UpdateOptions) (result *v1alpha1.LogzTarget, err error) {
	emptyResult := &v1alpha1.LogzTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(logztargetsResource, c.ns, logzTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LogzTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogzTargets) UpdateStatus(ctx context.Context, logzTarget *v1alpha1.LogzTarget, opts v1.UpdateOptions) (result *v1alpha1.LogzTarget, err error) {
	emptyResult := &v1alpha1.LogzTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(logztargetsResource, "status", c.ns, logzTarget, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LogzTarget), err
}

// Delete takes name of the logzTarget and deletes it. Returns an error if one occurs.
func (c *FakeLogzTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(logztargetsResource, c.ns, name, opts), &v1alpha1.LogzTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogzTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(logztargetsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogzTargetList{})
	return err
}

// Patch applies the patch and returns the patched logzTarget.
func (c *FakeLogzTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogzTarget, err error) {
	emptyResult := &v1alpha1.LogzTarget{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(logztargetsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LogzTarget), err
}
