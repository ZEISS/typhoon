// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/flow/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFlowV1alpha1 struct {
	*testing.Fake
}

func (c *FakeFlowV1alpha1) Bridges(namespace string) v1alpha1.BridgeInterface {
	return newFakeBridges(c, namespace)
}

func (c *FakeFlowV1alpha1) JQTransformations(namespace string) v1alpha1.JQTransformationInterface {
	return newFakeJQTransformations(c, namespace)
}

func (c *FakeFlowV1alpha1) Synchronizers(namespace string) v1alpha1.SynchronizerInterface {
	return newFakeSynchronizers(c, namespace)
}

func (c *FakeFlowV1alpha1) Transformations(namespace string) v1alpha1.TransformationInterface {
	return newFakeTransformations(c, namespace)
}

func (c *FakeFlowV1alpha1) WorkerTransformations(namespace string) v1alpha1.WorkerTransformationInterface {
	return newFakeWorkerTransformations(c, namespace)
}

func (c *FakeFlowV1alpha1) XMLToJSONTransformations(namespace string) v1alpha1.XMLToJSONTransformationInterface {
	return newFakeXMLToJSONTransformations(c, namespace)
}

func (c *FakeFlowV1alpha1) XSLTTransformations(namespace string) v1alpha1.XSLTTransformationInterface {
	return newFakeXSLTTransformations(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFlowV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
