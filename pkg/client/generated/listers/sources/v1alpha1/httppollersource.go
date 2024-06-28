// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HTTPPollerSourceLister helps list HTTPPollerSources.
// All objects returned here must be treated as read-only.
type HTTPPollerSourceLister interface {
	// List lists all HTTPPollerSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HTTPPollerSource, err error)
	// HTTPPollerSources returns an object that can list and get HTTPPollerSources.
	HTTPPollerSources(namespace string) HTTPPollerSourceNamespaceLister
	HTTPPollerSourceListerExpansion
}

// hTTPPollerSourceLister implements the HTTPPollerSourceLister interface.
type hTTPPollerSourceLister struct {
	indexer cache.Indexer
}

// NewHTTPPollerSourceLister returns a new HTTPPollerSourceLister.
func NewHTTPPollerSourceLister(indexer cache.Indexer) HTTPPollerSourceLister {
	return &hTTPPollerSourceLister{indexer: indexer}
}

// List lists all HTTPPollerSources in the indexer.
func (s *hTTPPollerSourceLister) List(selector labels.Selector) (ret []*v1alpha1.HTTPPollerSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HTTPPollerSource))
	})
	return ret, err
}

// HTTPPollerSources returns an object that can list and get HTTPPollerSources.
func (s *hTTPPollerSourceLister) HTTPPollerSources(namespace string) HTTPPollerSourceNamespaceLister {
	return hTTPPollerSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HTTPPollerSourceNamespaceLister helps list and get HTTPPollerSources.
// All objects returned here must be treated as read-only.
type HTTPPollerSourceNamespaceLister interface {
	// List lists all HTTPPollerSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HTTPPollerSource, err error)
	// Get retrieves the HTTPPollerSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HTTPPollerSource, error)
	HTTPPollerSourceNamespaceListerExpansion
}

// hTTPPollerSourceNamespaceLister implements the HTTPPollerSourceNamespaceLister
// interface.
type hTTPPollerSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HTTPPollerSources in the indexer for a given namespace.
func (s hTTPPollerSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HTTPPollerSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HTTPPollerSource))
	})
	return ret, err
}

// Get retrieves the HTTPPollerSource from the indexer for a given namespace and name.
func (s hTTPPollerSourceNamespaceLister) Get(name string) (*v1alpha1.HTTPPollerSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("httppollersource"), name)
	}
	return obj.(*v1alpha1.HTTPPollerSource), nil
}