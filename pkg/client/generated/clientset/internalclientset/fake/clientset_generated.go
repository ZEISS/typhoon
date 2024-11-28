// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset"
	extensionsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/extensions/v1alpha1"
	fakeextensionsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/extensions/v1alpha1/fake"
	flowv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1"
	fakeflowv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1/fake"
	routingv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/routing/v1alpha1"
	fakeroutingv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/routing/v1alpha1/fake"
	sourcesv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1"
	fakesourcesv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1/fake"
	targetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	faketargetsv1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any field management, validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
//
// DEPRECATED: NewClientset replaces this with support for field management, which significantly improves
// server side apply testing. NewClientset is only available when apply configurations are generated (e.g.
// via --with-applyconfig).
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ExtensionsV1alpha1 retrieves the ExtensionsV1alpha1Client
func (c *Clientset) ExtensionsV1alpha1() extensionsv1alpha1.ExtensionsV1alpha1Interface {
	return &fakeextensionsv1alpha1.FakeExtensionsV1alpha1{Fake: &c.Fake}
}

// FlowV1alpha1 retrieves the FlowV1alpha1Client
func (c *Clientset) FlowV1alpha1() flowv1alpha1.FlowV1alpha1Interface {
	return &fakeflowv1alpha1.FakeFlowV1alpha1{Fake: &c.Fake}
}

// RoutingV1alpha1 retrieves the RoutingV1alpha1Client
func (c *Clientset) RoutingV1alpha1() routingv1alpha1.RoutingV1alpha1Interface {
	return &fakeroutingv1alpha1.FakeRoutingV1alpha1{Fake: &c.Fake}
}

// SourcesV1alpha1 retrieves the SourcesV1alpha1Client
func (c *Clientset) SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface {
	return &fakesourcesv1alpha1.FakeSourcesV1alpha1{Fake: &c.Fake}
}

// TargetsV1alpha1 retrieves the TargetsV1alpha1Client
func (c *Clientset) TargetsV1alpha1() targetsv1alpha1.TargetsV1alpha1Interface {
	return &faketargetsv1alpha1.FakeTargetsV1alpha1{Fake: &c.Fake}
}
