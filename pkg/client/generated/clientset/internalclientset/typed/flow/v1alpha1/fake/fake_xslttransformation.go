// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXSLTTransformations implements XSLTTransformationInterface
type FakeXSLTTransformations struct {
	Fake *FakeFlowV1alpha1
	ns   string
}

var xslttransformationsResource = v1alpha1.SchemeGroupVersion.WithResource("xslttransformations")

var xslttransformationsKind = v1alpha1.SchemeGroupVersion.WithKind("XSLTTransformation")

// Get takes name of the xSLTTransformation, and returns the corresponding xSLTTransformation object, and an error if there is any.
func (c *FakeXSLTTransformations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.XSLTTransformation, err error) {
	emptyResult := &v1alpha1.XSLTTransformation{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(xslttransformationsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.XSLTTransformation), err
}

// List takes label and field selectors, and returns the list of XSLTTransformations that match those selectors.
func (c *FakeXSLTTransformations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.XSLTTransformationList, err error) {
	emptyResult := &v1alpha1.XSLTTransformationList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(xslttransformationsResource, xslttransformationsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.XSLTTransformationList{ListMeta: obj.(*v1alpha1.XSLTTransformationList).ListMeta}
	for _, item := range obj.(*v1alpha1.XSLTTransformationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xSLTTransformations.
func (c *FakeXSLTTransformations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(xslttransformationsResource, c.ns, opts))

}

// Create takes the representation of a xSLTTransformation and creates it.  Returns the server's representation of the xSLTTransformation, and an error, if there is any.
func (c *FakeXSLTTransformations) Create(ctx context.Context, xSLTTransformation *v1alpha1.XSLTTransformation, opts v1.CreateOptions) (result *v1alpha1.XSLTTransformation, err error) {
	emptyResult := &v1alpha1.XSLTTransformation{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(xslttransformationsResource, c.ns, xSLTTransformation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.XSLTTransformation), err
}

// Update takes the representation of a xSLTTransformation and updates it. Returns the server's representation of the xSLTTransformation, and an error, if there is any.
func (c *FakeXSLTTransformations) Update(ctx context.Context, xSLTTransformation *v1alpha1.XSLTTransformation, opts v1.UpdateOptions) (result *v1alpha1.XSLTTransformation, err error) {
	emptyResult := &v1alpha1.XSLTTransformation{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(xslttransformationsResource, c.ns, xSLTTransformation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.XSLTTransformation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeXSLTTransformations) UpdateStatus(ctx context.Context, xSLTTransformation *v1alpha1.XSLTTransformation, opts v1.UpdateOptions) (result *v1alpha1.XSLTTransformation, err error) {
	emptyResult := &v1alpha1.XSLTTransformation{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(xslttransformationsResource, "status", c.ns, xSLTTransformation, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.XSLTTransformation), err
}

// Delete takes name of the xSLTTransformation and deletes it. Returns an error if one occurs.
func (c *FakeXSLTTransformations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(xslttransformationsResource, c.ns, name, opts), &v1alpha1.XSLTTransformation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXSLTTransformations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(xslttransformationsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.XSLTTransformationList{})
	return err
}

// Patch applies the patch and returns the patched xSLTTransformation.
func (c *FakeXSLTTransformations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.XSLTTransformation, err error) {
	emptyResult := &v1alpha1.XSLTTransformation{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(xslttransformationsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.XSLTTransformation), err
}
