// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AzureServiceBusTopicSourceLister helps list AzureServiceBusTopicSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusTopicSourceLister interface {
	// List lists all AzureServiceBusTopicSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusTopicSource, err error)
	// AzureServiceBusTopicSources returns an object that can list and get AzureServiceBusTopicSources.
	AzureServiceBusTopicSources(namespace string) AzureServiceBusTopicSourceNamespaceLister
	AzureServiceBusTopicSourceListerExpansion
}

// azureServiceBusTopicSourceLister implements the AzureServiceBusTopicSourceLister interface.
type azureServiceBusTopicSourceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusTopicSource]
}

// NewAzureServiceBusTopicSourceLister returns a new AzureServiceBusTopicSourceLister.
func NewAzureServiceBusTopicSourceLister(indexer cache.Indexer) AzureServiceBusTopicSourceLister {
	return &azureServiceBusTopicSourceLister{listers.New[*sourcesv1alpha1.AzureServiceBusTopicSource](indexer, sourcesv1alpha1.Resource("azureservicebustopicsource"))}
}

// AzureServiceBusTopicSources returns an object that can list and get AzureServiceBusTopicSources.
func (s *azureServiceBusTopicSourceLister) AzureServiceBusTopicSources(namespace string) AzureServiceBusTopicSourceNamespaceLister {
	return azureServiceBusTopicSourceNamespaceLister{listers.NewNamespaced[*sourcesv1alpha1.AzureServiceBusTopicSource](s.ResourceIndexer, namespace)}
}

// AzureServiceBusTopicSourceNamespaceLister helps list and get AzureServiceBusTopicSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusTopicSourceNamespaceLister interface {
	// List lists all AzureServiceBusTopicSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusTopicSource, err error)
	// Get retrieves the AzureServiceBusTopicSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*sourcesv1alpha1.AzureServiceBusTopicSource, error)
	AzureServiceBusTopicSourceNamespaceListerExpansion
}

// azureServiceBusTopicSourceNamespaceLister implements the AzureServiceBusTopicSourceNamespaceLister
// interface.
type azureServiceBusTopicSourceNamespaceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusTopicSource]
}
