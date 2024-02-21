package datadogtarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig
	// APIKey defines a Datadog API key to be used for authentication
	APIKey string `envconfig:"DD_CLIENT_API_KEY" required:"true"`

	// Site defines the site of the Datadog intake API, defaults to `datadoghq.com`
	Site string `envconfig:"DD_CLIENT_SITE" default:"datadoghq.com"`

	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`

	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
}
