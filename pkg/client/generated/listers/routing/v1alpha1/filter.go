// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	routingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// FilterLister helps list Filters.
// All objects returned here must be treated as read-only.
type FilterLister interface {
	// List lists all Filters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*routingv1alpha1.Filter, err error)
	// Filters returns an object that can list and get Filters.
	Filters(namespace string) FilterNamespaceLister
	FilterListerExpansion
}

// filterLister implements the FilterLister interface.
type filterLister struct {
	listers.ResourceIndexer[*routingv1alpha1.Filter]
}

// NewFilterLister returns a new FilterLister.
func NewFilterLister(indexer cache.Indexer) FilterLister {
	return &filterLister{listers.New[*routingv1alpha1.Filter](indexer, routingv1alpha1.Resource("filter"))}
}

// Filters returns an object that can list and get Filters.
func (s *filterLister) Filters(namespace string) FilterNamespaceLister {
	return filterNamespaceLister{listers.NewNamespaced[*routingv1alpha1.Filter](s.ResourceIndexer, namespace)}
}

// FilterNamespaceLister helps list and get Filters.
// All objects returned here must be treated as read-only.
type FilterNamespaceLister interface {
	// List lists all Filters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*routingv1alpha1.Filter, err error)
	// Get retrieves the Filter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*routingv1alpha1.Filter, error)
	FilterNamespaceListerExpansion
}

// filterNamespaceLister implements the FilterNamespaceLister
// interface.
type filterNamespaceLister struct {
	listers.ResourceIndexer[*routingv1alpha1.Filter]
}
