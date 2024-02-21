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

// FakeTransformations implements TransformationInterface
type FakeTransformations struct {
	Fake *FakeFlowV1alpha1
	ns   string
}

var transformationsResource = v1alpha1.SchemeGroupVersion.WithResource("transformations")

var transformationsKind = v1alpha1.SchemeGroupVersion.WithKind("Transformation")

// Get takes name of the transformation, and returns the corresponding transformation object, and an error if there is any.
func (c *FakeTransformations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Transformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transformationsResource, c.ns, name), &v1alpha1.Transformation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Transformation), err
}

// List takes label and field selectors, and returns the list of Transformations that match those selectors.
func (c *FakeTransformations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransformationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transformationsResource, transformationsKind, c.ns, opts), &v1alpha1.TransformationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransformationList{ListMeta: obj.(*v1alpha1.TransformationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransformationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transformations.
func (c *FakeTransformations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transformationsResource, c.ns, opts))

}

// Create takes the representation of a transformation and creates it.  Returns the server's representation of the transformation, and an error, if there is any.
func (c *FakeTransformations) Create(ctx context.Context, transformation *v1alpha1.Transformation, opts v1.CreateOptions) (result *v1alpha1.Transformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transformationsResource, c.ns, transformation), &v1alpha1.Transformation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Transformation), err
}

// Update takes the representation of a transformation and updates it. Returns the server's representation of the transformation, and an error, if there is any.
func (c *FakeTransformations) Update(ctx context.Context, transformation *v1alpha1.Transformation, opts v1.UpdateOptions) (result *v1alpha1.Transformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transformationsResource, c.ns, transformation), &v1alpha1.Transformation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Transformation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransformations) UpdateStatus(ctx context.Context, transformation *v1alpha1.Transformation, opts v1.UpdateOptions) (*v1alpha1.Transformation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transformationsResource, "status", c.ns, transformation), &v1alpha1.Transformation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Transformation), err
}

// Delete takes name of the transformation and deletes it. Returns an error if one occurs.
func (c *FakeTransformations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(transformationsResource, c.ns, name, opts), &v1alpha1.Transformation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransformations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transformationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransformationList{})
	return err
}

// Patch applies the patch and returns the patched transformation.
func (c *FakeTransformations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Transformation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transformationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Transformation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Transformation), err
}
