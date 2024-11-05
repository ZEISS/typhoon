package jqtransformation

import "knative.dev/eventing/pkg/adapter/v2"

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	// Query represents the jq query to be applied to the incoming event
	Query string `envconfig:"JQ_QUERY" required:"true"`
	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`
	// Sink defines the target sink for the events. If no Sink is defined the
	// events are replied back to the sender.
	Sink string `envconfig:"K_SINK"`

	adapter.EnvConfig
}
