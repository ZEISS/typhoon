package logztarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	// ShippingToken defines the API token.
	ShippingToken string `envconfig:"LOGZ_SHIPPING_TOKEN" required:"true"`

	// LogsListenerURL Defines the Log listener URL
	LogsListenerURL string `envconfig:"LOGZ_LISTENER_URL" required:"true"`

	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`

	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`
}
