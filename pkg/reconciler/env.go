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

	// Common AWS attributes
	EnvARN             = "ARN"
	EnvAccessKeyID     = "AWS_ACCESS_KEY_ID"
	EnvSecretAccessKey = "AWS_SECRET_ACCESS_KEY" //nolint:gosec
	EnvSessionToken    = "AWS_SESSION_TOKEN"
	EnvEndpointURL     = "AWS_ENDPOINT_URL"
	EnvAssumeIamRole   = "AWS_ASSUME_ROLE_ARN"

	// Common Azure attributes
	EnvAADTenantID     = "AZURE_TENANT_ID"
	EnvAADClientID     = "AZURE_CLIENT_ID"
	EnvAADClientSecret = "AZURE_CLIENT_SECRET"

	// Azure Event Hub attributes
	// https://pkg.go.dev/github.com/Azure/azure-event-hubs-go/v3#readme-environment-variables
	EnvHubNamespace        = "EVENTHUB_NAMESPACE"
	EnvHubName             = "EVENTHUB_NAME"
	EnvHubKeyName          = "EVENTHUB_KEY_NAME"
	EnvHubKeyValue         = "EVENTHUB_KEY_VALUE"
	EnvHubConnStr          = "EVENTHUB_CONNECTION_STRING"
	EnvHubResourceID       = "EVENTHUB_RESOURCE_ID"
	EnvHubConsumerGroup    = "EVENTHUB_CONSUMER_GROUP"
	EnvHubMessageTimeout   = "EVENTHUB_MESSAGE_TIMEOUT"
	EnvHubMessageCountSize = "EVENTHUB_MESSAGE_COUNT_SIZE"

	// Azure Service Bus attributes
	EnvServiceBusKeyName          = "SERVICEBUS_KEY_NAME"
	EnvServiceBusKeyValue         = "SERVICEBUS_KEY_VALUE"
	EnvServiceBusConnStr          = "SERVICEBUS_CONNECTION_STRING"
	EnvServiceBusEntityResourceID = "SERVICEBUS_ENTITY_RESOURCE_ID"
	EnvServiceBusWebSocketsEnable = "SERVICEBUS_WEBSOCKETS_ENABLE"
	EnvServiceBusMaxConcurrent    = "SERVICEBUS_MAX_CONCURRENT"

	// Common Google Cloud attributes
	EnvGCloudSAKey = "GCLOUD_SERVICEACCOUNT_KEY"

	// Google Cloud Pub/Sub attributes
	EnvGCloudPubSubSubscription = "GCLOUD_PUBSUB_SUBSCRIPTION"
)
