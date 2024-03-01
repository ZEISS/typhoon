// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/zeiss/typhoon/pkg/apis/eventing/v1alpha1"
	"github.com/zeiss/typhoon/pkg/client/generated/clientset/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type EventingV1alpha1Interface interface {
	RESTClient() rest.Interface
	RedisBrokersGetter
	TriggersGetter
}

// EventingV1alpha1Client is used to interact with features provided by the eventing.typhoon.zeiss.com group.
type EventingV1alpha1Client struct {
	restClient rest.Interface
}

func (c *EventingV1alpha1Client) RedisBrokers(namespace string) RedisBrokerInterface {
	return newRedisBrokers(c, namespace)
}

func (c *EventingV1alpha1Client) Triggers(namespace string) TriggerInterface {
	return newTriggers(c, namespace)
}

// NewForConfig creates a new EventingV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*EventingV1alpha1Client, error) {
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

// NewForConfigAndClient creates a new EventingV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*EventingV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &EventingV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new EventingV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EventingV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new EventingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *EventingV1alpha1Client {
	return &EventingV1alpha1Client{c}
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
func (c *EventingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}