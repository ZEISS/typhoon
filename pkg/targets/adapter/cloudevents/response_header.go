package cloudevents

import (
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// ResponseHeaderValue is a function that given an
// incoming CloudEvent returns a string to be used
// as a header value on the outgoing event.
type ResponseHeaderValue func(event *cloudevents.Event) (string, error)

// StaticResponse always returns the same fixed value.
// Should be used when returned types or sources do not vary
// depending on the incoming event data.
func StaticResponse(value string) ResponseHeaderValue {
	return func(event *cloudevents.Event) (string, error) {
		return value, nil
	}
}

// MappedResponseType returns an static eventType based on the incoming eventType.
// When no type is mapped a default value is returned.
func MappedResponseType(eventTypes map[string]string) ResponseHeaderValue {
	return func(event *cloudevents.Event) (string, error) {
		v, ok := eventTypes[event.Type()]
		if ok {
			return v, nil
		}
		return "", fmt.Errorf("incoming type %q cannot be mapped to an outgoing event type", event.Type())
	}
}

// SuffixResponseType appends a string to the the incoming eventType.
func SuffixResponseType(suffix string) ResponseHeaderValue {
	return func(event *cloudevents.Event) (string, error) {
		return event.Type() + suffix, nil
	}
}
