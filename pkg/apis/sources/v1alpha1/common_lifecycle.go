package v1alpha1

// EventType returns an event type in a format suitable for usage as a
// CloudEvent type attribute.
func EventType(service, eventType string) string {
	return "com.zeiss." + service + "." + eventType
}

// AWSEventType returns an event type in a format suitable for usage as a
// CloudEvent type attribute.
func AWSEventType(awsService, eventType string) string {
	return "com.amazon." + awsService + "." + eventType
}
