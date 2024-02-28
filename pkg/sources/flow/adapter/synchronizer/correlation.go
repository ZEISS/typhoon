package synchronizer

import (
	"fmt"
	"math/rand"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2/event"
)

// Correlation Key charset.
const correlationKeycharset = "abcdefghijklmnopqrstuvwxyz0123456789"

var (
	// CloudEvent attributes cannot be used as a correltaion key.
	restrictedKeys = []string{
		"id",
		"type",
		"time",
		"subject",
		"schemaurl",
		"dataschema",
		"specversion",
		"datamediatype",
		"datacontenttype",
		"datacontentencoding",
	}

	// #nosec G404
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// correlationKey is the correlation attribute for the CloudEvents.
type correlationKey struct {
	attribute string
	length    int
}

// NewCorrelationKey returns an instance of the CloudEvent Correlation key.
func newCorrelationKey(attribute string, length int) (*correlationKey, error) {
	for _, rk := range restrictedKeys {
		if attribute == rk {
			return nil, fmt.Errorf("%q cannot be used as a correlation key", attribute)
		}
	}

	return &correlationKey{
		attribute: attribute,
		length:    length,
	}, nil
}

// Get returns the value of Correlation Key.
func (k *correlationKey) get(event cloudevents.Event) (string, bool) {
	if val, exists := event.Extensions()[k.attribute]; exists {
		return val.(string), true
	}
	return "", false
}

// Set updates the CloudEvent's context with the random Correlation Key value.
func (k *correlationKey) set(event *cloudevents.Event) string {
	correlationID := randString(k.length)
	event.SetExtension(k.attribute, correlationID)
	return correlationID
}

// randString generates the random string with fixed length.
func randString(length int) string {
	k := make([]byte, length)
	l := len(correlationKeycharset) - 1
	for i := range k {
		k[i] = correlationKeycharset[seededRand.Intn(l)]
	}
	return string(k)
}
