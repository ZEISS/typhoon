// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JQTransformationLister helps list JQTransformations.
// All objects returned here must be treated as read-only.
type JQTransformationLister interface {
	// List lists all JQTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JQTransformation, err error)
	// JQTransformations returns an object that can list and get JQTransformations.
	JQTransformations(namespace string) JQTransformationNamespaceLister
	JQTransformationListerExpansion
}

// jQTransformationLister implements the JQTransformationLister interface.
type jQTransformationLister struct {
	indexer cache.Indexer
}

// NewJQTransformationLister returns a new JQTransformationLister.
func NewJQTransformationLister(indexer cache.Indexer) JQTransformationLister {
	return &jQTransformationLister{indexer: indexer}
}

// List lists all JQTransformations in the indexer.
func (s *jQTransformationLister) List(selector labels.Selector) (ret []*v1alpha1.JQTransformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JQTransformation))
	})
	return ret, err
}

// JQTransformations returns an object that can list and get JQTransformations.
func (s *jQTransformationLister) JQTransformations(namespace string) JQTransformationNamespaceLister {
	return jQTransformationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JQTransformationNamespaceLister helps list and get JQTransformations.
// All objects returned here must be treated as read-only.
type JQTransformationNamespaceLister interface {
	// List lists all JQTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JQTransformation, err error)
	// Get retrieves the JQTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JQTransformation, error)
	JQTransformationNamespaceListerExpansion
}

// jQTransformationNamespaceLister implements the JQTransformationNamespaceLister
// interface.
type jQTransformationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JQTransformations in the indexer for a given namespace.
func (s jQTransformationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.JQTransformation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JQTransformation))
	})
	return ret, err
}

// Get retrieves the JQTransformation from the indexer for a given namespace and name.
func (s jQTransformationNamespaceLister) Get(name string) (*v1alpha1.JQTransformation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jqtransformation"), name)
	}
	return obj.(*v1alpha1.JQTransformation), nil
}
