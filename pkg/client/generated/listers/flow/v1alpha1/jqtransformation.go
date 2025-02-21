// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// JQTransformationLister helps list JQTransformations.
// All objects returned here must be treated as read-only.
type JQTransformationLister interface {
	// List lists all JQTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*flowv1alpha1.JQTransformation, err error)
	// JQTransformations returns an object that can list and get JQTransformations.
	JQTransformations(namespace string) JQTransformationNamespaceLister
	JQTransformationListerExpansion
}

// jQTransformationLister implements the JQTransformationLister interface.
type jQTransformationLister struct {
	listers.ResourceIndexer[*flowv1alpha1.JQTransformation]
}

// NewJQTransformationLister returns a new JQTransformationLister.
func NewJQTransformationLister(indexer cache.Indexer) JQTransformationLister {
	return &jQTransformationLister{listers.New[*flowv1alpha1.JQTransformation](indexer, flowv1alpha1.Resource("jqtransformation"))}
}

// JQTransformations returns an object that can list and get JQTransformations.
func (s *jQTransformationLister) JQTransformations(namespace string) JQTransformationNamespaceLister {
	return jQTransformationNamespaceLister{listers.NewNamespaced[*flowv1alpha1.JQTransformation](s.ResourceIndexer, namespace)}
}

// JQTransformationNamespaceLister helps list and get JQTransformations.
// All objects returned here must be treated as read-only.
type JQTransformationNamespaceLister interface {
	// List lists all JQTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*flowv1alpha1.JQTransformation, err error)
	// Get retrieves the JQTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*flowv1alpha1.JQTransformation, error)
	JQTransformationNamespaceListerExpansion
}

// jQTransformationNamespaceLister implements the JQTransformationNamespaceLister
// interface.
type jQTransformationNamespaceLister struct {
	listers.ResourceIndexer[*flowv1alpha1.JQTransformation]
}
