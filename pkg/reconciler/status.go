package reconciler

import duckv1 "knative.dev/pkg/apis/duck/v1"

// CreateCloudEventAttributes returns CloudEvent attributes for the event types
// supported by the target.
func CreateCloudEventAttributes(source string, eventTypes []string) []duckv1.CloudEventAttributes {
	ceAttributes := make([]duckv1.CloudEventAttributes, len(eventTypes))

	for i, typ := range eventTypes {
		ceAttributes[i] = duckv1.CloudEventAttributes{
			Type:   typ,
			Source: source,
		}
	}

	return ceAttributes
}
