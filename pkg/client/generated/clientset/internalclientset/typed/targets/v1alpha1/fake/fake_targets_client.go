// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/typed/targets/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTargetsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTargetsV1alpha1) CloudEventsTargets(namespace string) v1alpha1.CloudEventsTargetInterface {
	return &FakeCloudEventsTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) DatadogTargets(namespace string) v1alpha1.DatadogTargetInterface {
	return &FakeDatadogTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) HTTPTargets(namespace string) v1alpha1.HTTPTargetInterface {
	return &FakeHTTPTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) KafkaTargets(namespace string) v1alpha1.KafkaTargetInterface {
	return &FakeKafkaTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) LogzMetricsTargets(namespace string) v1alpha1.LogzMetricsTargetInterface {
	return &FakeLogzMetricsTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) LogzTargets(namespace string) v1alpha1.LogzTargetInterface {
	return &FakeLogzTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) NatsTargets(namespace string) v1alpha1.NatsTargetInterface {
	return &FakeNatsTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) SalesforceTargets(namespace string) v1alpha1.SalesforceTargetInterface {
	return &FakeSalesforceTargets{c, namespace}
}

func (c *FakeTargetsV1alpha1) SplunkTargets(namespace string) v1alpha1.SplunkTargetInterface {
	return &FakeSplunkTargets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTargetsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
