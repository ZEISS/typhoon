// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/routing/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// SplitterLister helps list Splitters.
// All objects returned here must be treated as read-only.
type SplitterLister interface {
	// List lists all Splitters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error)
	// Splitters returns an object that can list and get Splitters.
	Splitters(namespace string) SplitterNamespaceLister
	SplitterListerExpansion
}

// splitterLister implements the SplitterLister interface.
type splitterLister struct {
	listers.ResourceIndexer[*v1alpha1.Splitter]
}

// NewSplitterLister returns a new SplitterLister.
func NewSplitterLister(indexer cache.Indexer) SplitterLister {
	return &splitterLister{listers.New[*v1alpha1.Splitter](indexer, v1alpha1.Resource("splitter"))}
}

// Splitters returns an object that can list and get Splitters.
func (s *splitterLister) Splitters(namespace string) SplitterNamespaceLister {
	return splitterNamespaceLister{listers.NewNamespaced[*v1alpha1.Splitter](s.ResourceIndexer, namespace)}
}

// SplitterNamespaceLister helps list and get Splitters.
// All objects returned here must be treated as read-only.
type SplitterNamespaceLister interface {
	// List lists all Splitters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error)
	// Get retrieves the Splitter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Splitter, error)
	SplitterNamespaceListerExpansion
}

// splitterNamespaceLister implements the SplitterNamespaceLister
// interface.
type splitterNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.Splitter]
}
