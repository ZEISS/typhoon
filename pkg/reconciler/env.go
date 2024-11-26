package reconciler

// Common environment variables propagated to adapters.
const (
	EnvName      = "NAME"
	EnvNamespace = "NAMESPACE"

	envSink                  = "K_SINK"
	envComponent             = "K_COMPONENT"
	envSinkTimeout           = "K_SINK_TIMEOUT"
	envMetricsPrometheusPort = "METRICS_PROMETHEUS_PORT"

	// Overrides for CloudEvents context attributes (only supported by a subset of components)
	EnvCESource = "CE_SOURCE"
	EnvCEType   = "CE_TYPE"

	// Common Azure attributes
	EnvAADTenantID     = "AZURE_TENANT_ID"
	EnvAADClientID     = "AZURE_CLIENT_ID"
	EnvAADClientSecret = "AZURE_CLIENT_SECRET"

	// Azure Service Bus attributes
	EnvServiceBusKeyName          = "SERVICEBUS_KEY_NAME"
	EnvServiceBusKeyValue         = "SERVICEBUS_KEY_VALUE"
	EnvServiceBusConnStr          = "SERVICEBUS_CONNECTION_STRING"
	EnvServiceBusEntityResourceID = "SERVICEBUS_ENTITY_RESOURCE_ID"
	EnvServiceBusWebSocketsEnable = "SERVICEBUS_WEBSOCKETS_ENABLE"
	EnvServiceBusMaxConcurrent    = "SERVICEBUS_MAX_CONCURRENT"
)
