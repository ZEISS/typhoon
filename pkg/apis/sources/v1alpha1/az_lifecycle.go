package v1alpha1

// AzureEventType returns an event type in a format suitable for usage as a
// CloudEvent type attribute.
func AzureEventType(service, eventType string) string {
	return "com.microsoft.azure." + service + "." + eventType
}

// Reasons for status conditions
const (
	// AzureReasonNoClient is set on a status condition when an Azure API client cannot be obtained.
	AzureReasonNoClient = "NoClient"
	// AzureReasonAPIError is set on a status condition when an Azure API returns an error.
	AzureReasonAPIError = "APIError"
)
