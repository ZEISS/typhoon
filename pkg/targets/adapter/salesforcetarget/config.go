package salesforcetarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessor for configuration parameters
func EnvAccessor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	ClientID   string `envconfig:"SALESFORCE_AUTH_CLIENT_ID" required:"true"`
	AuthServer string `envconfig:"SALESFORCE_AUTH_SERVER" required:"true"`
	User       string `envconfig:"SALESFORCE_AUTH_USER" required:"true"`
	CertKey    string `envconfig:"SALESFORCE_AUTH_CERT_KEY" required:"true"`
	Version    string `envconfig:"SALESFORCE_API_VERSION"`

	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"always"`

	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
}
