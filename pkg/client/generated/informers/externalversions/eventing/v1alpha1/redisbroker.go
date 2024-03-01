// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	eventingv1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
	internalclientset "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/zeiss/typhoon/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/eventing/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RedisBrokerInformer provides access to a shared informer and lister for
// RedisBrokers.
type RedisBrokerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RedisBrokerLister
}

type redisBrokerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRedisBrokerInformer constructs a new informer for RedisBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRedisBrokerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRedisBrokerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRedisBrokerInformer constructs a new informer for RedisBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRedisBrokerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1alpha1().RedisBrokers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1alpha1().RedisBrokers(namespace).Watch(context.TODO(), options)
			},
		},
		&eventingv1alpha1.RedisBroker{},
		resyncPeriod,
		indexers,
	)
}

func (f *redisBrokerInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRedisBrokerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *redisBrokerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventingv1alpha1.RedisBroker{}, f.defaultInformer)
}

func (f *redisBrokerInformer) Lister() v1alpha1.RedisBrokerLister {
	return v1alpha1.NewRedisBrokerLister(f.Informer().GetIndexer())
}
