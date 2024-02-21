package azureservicebustarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	// Resource ID of the Service Bus entity (Queue or Topic subscription).
	EntityResourceID string `envconfig:"SERVICEBUS_ENTITY_RESOURCE_ID" required:"true"`

	// WebSocketsEnable.
	WebSocketsEnable bool `envconfig:"SERVICEBUS_WEBSOCKETS_ENABLE" default:"false"`

	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`

	// CloudEvents responses parametrization
	CloudEventPayloadPolicy string `envconfig:"EVENTS_PAYLOAD_POLICY" default:"error"`

	// DiscardCEContext chooses to keep or discard the incoming cloudevent context
	DiscardCEContext bool `envconfig:"DISCARD_CE_CONTEXT"`

	// The environment variables below aren't read from the envConfig struct
	// by the Service Bus SDK, but rather directly using os.Getenv().
	// They are nevertheless listed here for documentation purposes.
	_ string `envconfig:"AZURE_TENANT_ID"`
	_ string `envconfig:"AZURE_CLIENT_ID"`
	_ string `envconfig:"AZURE_CLIENT_SECRET"`
	_ string `envconfig:"SERVICEBUS_KEY_NAME"`
	_ string `envconfig:"SERVICEBUS_KEY_VALUE"`
	_ string `envconfig:"SERVICEBUS_CONNECTION_STRING"`
}
