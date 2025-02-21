// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AzureServiceBusQueueSourceLister helps list AzureServiceBusQueueSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusQueueSourceLister interface {
	// List lists all AzureServiceBusQueueSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusQueueSource, err error)
	// AzureServiceBusQueueSources returns an object that can list and get AzureServiceBusQueueSources.
	AzureServiceBusQueueSources(namespace string) AzureServiceBusQueueSourceNamespaceLister
	AzureServiceBusQueueSourceListerExpansion
}

// azureServiceBusQueueSourceLister implements the AzureServiceBusQueueSourceLister interface.
type azureServiceBusQueueSourceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusQueueSource]
}

// NewAzureServiceBusQueueSourceLister returns a new AzureServiceBusQueueSourceLister.
func NewAzureServiceBusQueueSourceLister(indexer cache.Indexer) AzureServiceBusQueueSourceLister {
	return &azureServiceBusQueueSourceLister{listers.New[*sourcesv1alpha1.AzureServiceBusQueueSource](indexer, sourcesv1alpha1.Resource("azureservicebusqueuesource"))}
}

// AzureServiceBusQueueSources returns an object that can list and get AzureServiceBusQueueSources.
func (s *azureServiceBusQueueSourceLister) AzureServiceBusQueueSources(namespace string) AzureServiceBusQueueSourceNamespaceLister {
	return azureServiceBusQueueSourceNamespaceLister{listers.NewNamespaced[*sourcesv1alpha1.AzureServiceBusQueueSource](s.ResourceIndexer, namespace)}
}

// AzureServiceBusQueueSourceNamespaceLister helps list and get AzureServiceBusQueueSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusQueueSourceNamespaceLister interface {
	// List lists all AzureServiceBusQueueSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusQueueSource, err error)
	// Get retrieves the AzureServiceBusQueueSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*sourcesv1alpha1.AzureServiceBusQueueSource, error)
	AzureServiceBusQueueSourceNamespaceListerExpansion
}

// azureServiceBusQueueSourceNamespaceLister implements the AzureServiceBusQueueSourceNamespaceLister
// interface.
type azureServiceBusQueueSourceNamespaceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusQueueSource]
}
