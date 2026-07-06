package jqtransformation

import "knative.dev/eventing/pkg/adapter/v2"

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig
	Query                   string `envconfig:"JQ_QUERY" required:"true"`
	BridgeIdentifier        string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`
	Sink                    string `envconfig:"K_SINK"`
}
