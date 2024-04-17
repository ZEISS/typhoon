// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	internalclientset "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/zeiss/typhoon/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SalesforceTargetInformer provides access to a shared informer and lister for
// SalesforceTargets.
type SalesforceTargetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SalesforceTargetLister
}

type salesforceTargetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSalesforceTargetInformer constructs a new informer for SalesforceTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSalesforceTargetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSalesforceTargetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSalesforceTargetInformer constructs a new informer for SalesforceTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSalesforceTargetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TargetsV1alpha1().SalesforceTargets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TargetsV1alpha1().SalesforceTargets(namespace).Watch(context.TODO(), options)
			},
		},
		&targetsv1alpha1.SalesforceTarget{},
		resyncPeriod,
		indexers,
	)
}

func (f *salesforceTargetInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSalesforceTargetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *salesforceTargetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&targetsv1alpha1.SalesforceTarget{}, f.defaultInformer)
}

func (f *salesforceTargetInformer) Lister() v1alpha1.SalesforceTargetLister {
	return v1alpha1.NewSalesforceTargetLister(f.Informer().GetIndexer())
}
