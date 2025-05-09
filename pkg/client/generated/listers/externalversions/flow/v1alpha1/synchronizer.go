// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apisflowv1alpha1 "github.com/zeiss/typhoon/pkg/apis/flow/v1alpha1"
	internalclientset "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/internalinterfaces"
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/flow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SynchronizerInformer provides access to a shared informer and lister for
// Synchronizers.
type SynchronizerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() flowv1alpha1.SynchronizerLister
}

type synchronizerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSynchronizerInformer constructs a new informer for Synchronizer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSynchronizerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSynchronizerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSynchronizerInformer constructs a new informer for Synchronizer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSynchronizerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().Synchronizers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().Synchronizers(namespace).Watch(context.TODO(), options)
			},
		},
		&apisflowv1alpha1.Synchronizer{},
		resyncPeriod,
		indexers,
	)
}

func (f *synchronizerInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSynchronizerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *synchronizerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisflowv1alpha1.Synchronizer{}, f.defaultInformer)
}

func (f *synchronizerInformer) Lister() flowv1alpha1.SynchronizerLister {
	return flowv1alpha1.NewSynchronizerLister(f.Informer().GetIndexer())
}
