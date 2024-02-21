// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// XSLTTransformationLister helps list XSLTTransformations.
// All objects returned here must be treated as read-only.
type XSLTTransformationLister interface {
	// List lists all XSLTTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.XSLTTransformation, err error)
	// XSLTTransformations returns an object that can list and get XSLTTransformations.
	XSLTTransformations(namespace string) XSLTTransformationNamespaceLister
	XSLTTransformationListerExpansion
}

// xSLTTransformationLister implements the XSLTTransformationLister interface.
type xSLTTransformationLister struct {
	indexer cache.Indexer
}

// NewXSLTTransformationLister returns a new XSLTTransformationLister.
func NewXSLTTransformationLister(indexer cache.Indexer) XSLTTransformationLister {
	return &xSLTTransformationLister{indexer: indexer}
}

// List lists all XSLTTransformations in the indexer.
func (s *xSLTTransformationLister) List(selector labels.Selector) (ret []*v1alpha1.XSLTTransformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XSLTTransformation))
	})
	return ret, err
}

// XSLTTransformations returns an object that can list and get XSLTTransformations.
func (s *xSLTTransformationLister) XSLTTransformations(namespace string) XSLTTransformationNamespaceLister {
	return xSLTTransformationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// XSLTTransformationNamespaceLister helps list and get XSLTTransformations.
// All objects returned here must be treated as read-only.
type XSLTTransformationNamespaceLister interface {
	// List lists all XSLTTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.XSLTTransformation, err error)
	// Get retrieves the XSLTTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.XSLTTransformation, error)
	XSLTTransformationNamespaceListerExpansion
}

// xSLTTransformationNamespaceLister implements the XSLTTransformationNamespaceLister
// interface.
type xSLTTransformationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all XSLTTransformations in the indexer for a given namespace.
func (s xSLTTransformationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.XSLTTransformation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XSLTTransformation))
	})
	return ret, err
}

// Get retrieves the XSLTTransformation from the indexer for a given namespace and name.
func (s xSLTTransformationNamespaceLister) Get(name string) (*v1alpha1.XSLTTransformation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("xslttransformation"), name)
	}
	return obj.(*v1alpha1.XSLTTransformation), nil
}
