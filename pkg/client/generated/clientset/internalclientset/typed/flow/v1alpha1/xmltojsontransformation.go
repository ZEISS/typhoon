// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// XMLToJSONTransformationsGetter has a method to return a XMLToJSONTransformationInterface.
// A group's client should implement this interface.
type XMLToJSONTransformationsGetter interface {
	XMLToJSONTransformations(namespace string) XMLToJSONTransformationInterface
}

// XMLToJSONTransformationInterface has methods to work with XMLToJSONTransformation resources.
type XMLToJSONTransformationInterface interface {
	Create(ctx context.Context, xMLToJSONTransformation *v1alpha1.XMLToJSONTransformation, opts v1.CreateOptions) (*v1alpha1.XMLToJSONTransformation, error)
	Update(ctx context.Context, xMLToJSONTransformation *v1alpha1.XMLToJSONTransformation, opts v1.UpdateOptions) (*v1alpha1.XMLToJSONTransformation, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, xMLToJSONTransformation *v1alpha1.XMLToJSONTransformation, opts v1.UpdateOptions) (*v1alpha1.XMLToJSONTransformation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.XMLToJSONTransformation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.XMLToJSONTransformationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.XMLToJSONTransformation, err error)
	XMLToJSONTransformationExpansion
}

// xMLToJSONTransformations implements XMLToJSONTransformationInterface
type xMLToJSONTransformations struct {
	*gentype.ClientWithList[*v1alpha1.XMLToJSONTransformation, *v1alpha1.XMLToJSONTransformationList]
}

// newXMLToJSONTransformations returns a XMLToJSONTransformations
func newXMLToJSONTransformations(c *FlowV1alpha1Client, namespace string) *xMLToJSONTransformations {
	return &xMLToJSONTransformations{
		gentype.NewClientWithList[*v1alpha1.XMLToJSONTransformation, *v1alpha1.XMLToJSONTransformationList](
			"xmltojsontransformations",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.XMLToJSONTransformation { return &v1alpha1.XMLToJSONTransformation{} },
			func() *v1alpha1.XMLToJSONTransformationList { return &v1alpha1.XMLToJSONTransformationList{} }),
	}
}
