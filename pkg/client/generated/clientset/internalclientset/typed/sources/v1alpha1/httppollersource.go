// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	scheme "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// HTTPPollerSourcesGetter has a method to return a HTTPPollerSourceInterface.
// A group's client should implement this interface.
type HTTPPollerSourcesGetter interface {
	HTTPPollerSources(namespace string) HTTPPollerSourceInterface
}

// HTTPPollerSourceInterface has methods to work with HTTPPollerSource resources.
type HTTPPollerSourceInterface interface {
	Create(ctx context.Context, hTTPPollerSource *sourcesv1alpha1.HTTPPollerSource, opts v1.CreateOptions) (*sourcesv1alpha1.HTTPPollerSource, error)
	Update(ctx context.Context, hTTPPollerSource *sourcesv1alpha1.HTTPPollerSource, opts v1.UpdateOptions) (*sourcesv1alpha1.HTTPPollerSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, hTTPPollerSource *sourcesv1alpha1.HTTPPollerSource, opts v1.UpdateOptions) (*sourcesv1alpha1.HTTPPollerSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*sourcesv1alpha1.HTTPPollerSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*sourcesv1alpha1.HTTPPollerSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *sourcesv1alpha1.HTTPPollerSource, err error)
	HTTPPollerSourceExpansion
}

// hTTPPollerSources implements HTTPPollerSourceInterface
type hTTPPollerSources struct {
	*gentype.ClientWithList[*sourcesv1alpha1.HTTPPollerSource, *sourcesv1alpha1.HTTPPollerSourceList]
}

// newHTTPPollerSources returns a HTTPPollerSources
func newHTTPPollerSources(c *SourcesV1alpha1Client, namespace string) *hTTPPollerSources {
	return &hTTPPollerSources{
		gentype.NewClientWithList[*sourcesv1alpha1.HTTPPollerSource, *sourcesv1alpha1.HTTPPollerSourceList](
			"httppollersources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *sourcesv1alpha1.HTTPPollerSource { return &sourcesv1alpha1.HTTPPollerSource{} },
			func() *sourcesv1alpha1.HTTPPollerSourceList { return &sourcesv1alpha1.HTTPPollerSourceList{} },
		),
	}
}
