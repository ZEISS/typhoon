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

// OCIMetricsSourcesGetter has a method to return a OCIMetricsSourceInterface.
// A group's client should implement this interface.
type OCIMetricsSourcesGetter interface {
	OCIMetricsSources(namespace string) OCIMetricsSourceInterface
}

// OCIMetricsSourceInterface has methods to work with OCIMetricsSource resources.
type OCIMetricsSourceInterface interface {
	Create(ctx context.Context, oCIMetricsSource *sourcesv1alpha1.OCIMetricsSource, opts v1.CreateOptions) (*sourcesv1alpha1.OCIMetricsSource, error)
	Update(ctx context.Context, oCIMetricsSource *sourcesv1alpha1.OCIMetricsSource, opts v1.UpdateOptions) (*sourcesv1alpha1.OCIMetricsSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, oCIMetricsSource *sourcesv1alpha1.OCIMetricsSource, opts v1.UpdateOptions) (*sourcesv1alpha1.OCIMetricsSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*sourcesv1alpha1.OCIMetricsSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*sourcesv1alpha1.OCIMetricsSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *sourcesv1alpha1.OCIMetricsSource, err error)
	OCIMetricsSourceExpansion
}

// oCIMetricsSources implements OCIMetricsSourceInterface
type oCIMetricsSources struct {
	*gentype.ClientWithList[*sourcesv1alpha1.OCIMetricsSource, *sourcesv1alpha1.OCIMetricsSourceList]
}

// newOCIMetricsSources returns a OCIMetricsSources
func newOCIMetricsSources(c *SourcesV1alpha1Client, namespace string) *oCIMetricsSources {
	return &oCIMetricsSources{
		gentype.NewClientWithList[*sourcesv1alpha1.OCIMetricsSource, *sourcesv1alpha1.OCIMetricsSourceList](
			"ocimetricssources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *sourcesv1alpha1.OCIMetricsSource { return &sourcesv1alpha1.OCIMetricsSource{} },
			func() *sourcesv1alpha1.OCIMetricsSourceList { return &sourcesv1alpha1.OCIMetricsSourceList{} },
		),
	}
}
