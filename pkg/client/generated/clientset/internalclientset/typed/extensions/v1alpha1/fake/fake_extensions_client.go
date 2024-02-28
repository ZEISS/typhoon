// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/extensions/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeExtensionsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeExtensionsV1alpha1) Functions(namespace string) v1alpha1.FunctionInterface {
	return &FakeFunctions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeExtensionsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}