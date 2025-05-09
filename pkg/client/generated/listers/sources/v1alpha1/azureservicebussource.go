// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/apis/sources/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AzureServiceBusSourceLister helps list AzureServiceBusSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusSourceLister interface {
	// List lists all AzureServiceBusSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusSource, err error)
	// AzureServiceBusSources returns an object that can list and get AzureServiceBusSources.
	AzureServiceBusSources(namespace string) AzureServiceBusSourceNamespaceLister
	AzureServiceBusSourceListerExpansion
}

// azureServiceBusSourceLister implements the AzureServiceBusSourceLister interface.
type azureServiceBusSourceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusSource]
}

// NewAzureServiceBusSourceLister returns a new AzureServiceBusSourceLister.
func NewAzureServiceBusSourceLister(indexer cache.Indexer) AzureServiceBusSourceLister {
	return &azureServiceBusSourceLister{listers.New[*sourcesv1alpha1.AzureServiceBusSource](indexer, sourcesv1alpha1.Resource("azureservicebussource"))}
}

// AzureServiceBusSources returns an object that can list and get AzureServiceBusSources.
func (s *azureServiceBusSourceLister) AzureServiceBusSources(namespace string) AzureServiceBusSourceNamespaceLister {
	return azureServiceBusSourceNamespaceLister{listers.NewNamespaced[*sourcesv1alpha1.AzureServiceBusSource](s.ResourceIndexer, namespace)}
}

// AzureServiceBusSourceNamespaceLister helps list and get AzureServiceBusSources.
// All objects returned here must be treated as read-only.
type AzureServiceBusSourceNamespaceLister interface {
	// List lists all AzureServiceBusSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*sourcesv1alpha1.AzureServiceBusSource, err error)
	// Get retrieves the AzureServiceBusSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*sourcesv1alpha1.AzureServiceBusSource, error)
	AzureServiceBusSourceNamespaceListerExpansion
}

// azureServiceBusSourceNamespaceLister implements the AzureServiceBusSourceNamespaceLister
// interface.
type azureServiceBusSourceNamespaceLister struct {
	listers.ResourceIndexer[*sourcesv1alpha1.AzureServiceBusSource]
}
