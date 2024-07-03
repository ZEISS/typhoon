// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/zeiss/typhoon/pkg/client/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Filters returns a FilterInformer.
	Filters() FilterInformer
	// Splitters returns a SplitterInformer.
	Splitters() SplitterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Filters returns a FilterInformer.
func (v *version) Filters() FilterInformer {
	return &filterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Splitters returns a SplitterInformer.
func (v *version) Splitters() SplitterInformer {
	return &splitterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}