// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HTTPTargetLister helps list HTTPTargets.
// All objects returned here must be treated as read-only.
type HTTPTargetLister interface {
	// List lists all HTTPTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HTTPTarget, err error)
	// HTTPTargets returns an object that can list and get HTTPTargets.
	HTTPTargets(namespace string) HTTPTargetNamespaceLister
	HTTPTargetListerExpansion
}

// hTTPTargetLister implements the HTTPTargetLister interface.
type hTTPTargetLister struct {
	indexer cache.Indexer
}

// NewHTTPTargetLister returns a new HTTPTargetLister.
func NewHTTPTargetLister(indexer cache.Indexer) HTTPTargetLister {
	return &hTTPTargetLister{indexer: indexer}
}

// List lists all HTTPTargets in the indexer.
func (s *hTTPTargetLister) List(selector labels.Selector) (ret []*v1alpha1.HTTPTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HTTPTarget))
	})
	return ret, err
}

// HTTPTargets returns an object that can list and get HTTPTargets.
func (s *hTTPTargetLister) HTTPTargets(namespace string) HTTPTargetNamespaceLister {
	return hTTPTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HTTPTargetNamespaceLister helps list and get HTTPTargets.
// All objects returned here must be treated as read-only.
type HTTPTargetNamespaceLister interface {
	// List lists all HTTPTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HTTPTarget, err error)
	// Get retrieves the HTTPTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HTTPTarget, error)
	HTTPTargetNamespaceListerExpansion
}

// hTTPTargetNamespaceLister implements the HTTPTargetNamespaceLister
// interface.
type hTTPTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HTTPTargets in the indexer for a given namespace.
func (s hTTPTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HTTPTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HTTPTarget))
	})
	return ret, err
}

// Get retrieves the HTTPTarget from the indexer for a given namespace and name.
func (s hTTPTargetNamespaceLister) Get(name string) (*v1alpha1.HTTPTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("httptarget"), name)
	}
	return obj.(*v1alpha1.HTTPTarget), nil
}