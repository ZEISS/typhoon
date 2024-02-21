// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/targets/v1alpha1"
	"github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type TargetsV1alpha1Interface interface {
	RESTClient() rest.Interface
	CloudEventsTargetsGetter
	DatadogTargetsGetter
	HTTPTargetsGetter
	KafkaTargetsGetter
	LogzMetricsTargetsGetter
	LogzTargetsGetter
	OracleTargetsGetter
	SplunkTargetsGetter
}

// TargetsV1alpha1Client is used to interact with features provided by the targets.typhoon.zeiss.com group.
type TargetsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TargetsV1alpha1Client) CloudEventsTargets(namespace string) CloudEventsTargetInterface {
	return newCloudEventsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) DatadogTargets(namespace string) DatadogTargetInterface {
	return newDatadogTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) HTTPTargets(namespace string) HTTPTargetInterface {
	return newHTTPTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) KafkaTargets(namespace string) KafkaTargetInterface {
	return newKafkaTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) LogzMetricsTargets(namespace string) LogzMetricsTargetInterface {
	return newLogzMetricsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) LogzTargets(namespace string) LogzTargetInterface {
	return newLogzTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) OracleTargets(namespace string) OracleTargetInterface {
	return newOracleTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SplunkTargets(namespace string) SplunkTargetInterface {
	return newSplunkTargets(c, namespace)
}

// NewForConfig creates a new TargetsV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*TargetsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new TargetsV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*TargetsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &TargetsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TargetsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TargetsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TargetsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TargetsV1alpha1Client {
	return &TargetsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TargetsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
