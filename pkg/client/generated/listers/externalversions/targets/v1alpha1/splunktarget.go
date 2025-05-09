// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apistargetsv1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	internalclientset "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/zeiss/typhoon/pkg/client/generated/listers/externalversions/internalinterfaces"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/listers/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SplunkTargetInformer provides access to a shared informer and lister for
// SplunkTargets.
type SplunkTargetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() targetsv1alpha1.SplunkTargetLister
}

type splunkTargetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSplunkTargetInformer constructs a new informer for SplunkTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSplunkTargetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSplunkTargetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSplunkTargetInformer constructs a new informer for SplunkTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSplunkTargetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TargetsV1alpha1().SplunkTargets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TargetsV1alpha1().SplunkTargets(namespace).Watch(context.TODO(), options)
			},
		},
		&apistargetsv1alpha1.SplunkTarget{},
		resyncPeriod,
		indexers,
	)
}

func (f *splunkTargetInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSplunkTargetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *splunkTargetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apistargetsv1alpha1.SplunkTarget{}, f.defaultInformer)
}

func (f *splunkTargetInformer) Lister() targetsv1alpha1.SplunkTargetLister {
	return targetsv1alpha1.NewSplunkTargetLister(f.Informer().GetIndexer())
}
