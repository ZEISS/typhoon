// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureServiceBusTopicSourceLister helps list AzureServiceBusTopicSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusTopicSourceLister interface {
	// List lists all AzureServiceBusTopicSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureServiceBusTopicSource, err error)
	// AzureServiceBusTopicSources returns an object that can list and get AzureServiceBusTopicSources.
	AzureServiceBusTopicSources(namespace string) AzureServiceBusTopicSourceNamespaceLister
	AzureServiceBusTopicSourceListerExpansion
}

// azureServiceBusTopicSourceLister implements the AzureServiceBusTopicSourceLister interface.
type azureServiceBusTopicSourceLister struct {
	indexer cache.Indexer
}

// NewAzureServiceBusTopicSourceLister returns a new AzureServiceBusTopicSourceLister.
func NewAzureServiceBusTopicSourceLister(indexer cache.Indexer) AzureServiceBusTopicSourceLister {
	return &azureServiceBusTopicSourceLister{indexer: indexer}
}

// List lists all AzureServiceBusTopicSources in the indexer.
func (s *azureServiceBusTopicSourceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureServiceBusTopicSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureServiceBusTopicSource))
	})
	return ret, err
}

// AzureServiceBusTopicSources returns an object that can list and get AzureServiceBusTopicSources.
func (s *azureServiceBusTopicSourceLister) AzureServiceBusTopicSources(namespace string) AzureServiceBusTopicSourceNamespaceLister {
	return azureServiceBusTopicSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureServiceBusTopicSourceNamespaceLister helps list and get AzureServiceBusTopicSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusTopicSourceNamespaceLister interface {
	// List lists all AzureServiceBusTopicSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureServiceBusTopicSource, err error)
	// Get retrieves the AzureServiceBusTopicSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AzureServiceBusTopicSource, error)
	AzureServiceBusTopicSourceNamespaceListerExpansion
}

// azureServiceBusTopicSourceNamespaceLister implements the AzureServiceBusTopicSourceNamespaceLister
// interface.
type azureServiceBusTopicSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureServiceBusTopicSources in the indexer for a given namespace.
func (s azureServiceBusTopicSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureServiceBusTopicSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureServiceBusTopicSource))
	})
	return ret, err
}

// Get retrieves the AzureServiceBusTopicSource from the indexer for a given namespace and name.
func (s azureServiceBusTopicSourceNamespaceLister) Get(name string) (*v1alpha1.AzureServiceBusTopicSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azureservicebustopicsource"), name)
	}
	return obj.(*v1alpha1.AzureServiceBusTopicSource), nil
}