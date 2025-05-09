// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// PingSourceLister helps list PingSources.
// All objects returned here must be treated as read-only.
type PingSourceLister interface {
	// List lists all PingSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.PingSource, err error)
	// PingSources returns an object that can list and get PingSources.
	PingSources(namespace string) PingSourceNamespaceLister
	PingSourceListerExpansion
}

// pingSourceLister implements the PingSourceLister interface.
type pingSourceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.PingSource]
}

// NewPingSourceLister returns a new PingSourceLister.
func NewPingSourceLister(indexer cache.Indexer) PingSourceLister {
	return &pingSourceLister{listers.New[*sourcesv1alpha1.PingSource](indexer, sourcesv1alpha1.Resource("pingsource"))}
}

// PingSources returns an object that can list and get PingSources.
func (s *pingSourceLister) PingSources(namespace string) PingSourceNamespaceLister {
	return pingSourceNamespaceLister{listers.NewNamespaced[*sourcesv1alpha1.PingSource](s.ResourceIndexer, namespace)}
}

// PingSourceNamespaceLister helps list and get PingSources.
// All objects returned here must be treated as read-only.
type PingSourceNamespaceLister interface {
	// List lists all PingSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.PingSource, err error)
	// Get retrieves the PingSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*sourcesv1alpha1.PingSource, error)
	PingSourceNamespaceListerExpansion
}

// pingSourceNamespaceLister implements the PingSourceNamespaceLister
// interface.
type pingSourceNamespaceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.PingSource]
}
