package webhooksource

import (
	"fmt"
	"strings"

	"knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig

	EventType                    string                   `envconfig:"WEBHOOK_EVENT_TYPE" required:"true"`
	EventSource                  string                   `envconfig:"WEBHOOK_EVENT_SOURCE" required:"true"`
	EventExtensionAttributesFrom *ExtensionAttributesFrom `envconfig:"WEBHOOK_EVENT_EXTENSION_ATTRIBUTES_FROM"`
	BasicAuthUsername            string                   `envconfig:"WEBHOOK_BASICAUTH_USERNAME"`
	BasicAuthPassword            string                   `envconfig:"WEBHOOK_BASICAUTH_PASSWORD"`
	CORSAllowOrigin              string                   `envconfig:"WEBHOOK_CORS_ALLOW_ORIGIN"`
}

// ExtensionAttributesFrom holds the configuration for the extension attributes to be extracted from the HTTP request.
type ExtensionAttributesFrom struct {
	headers bool
	host    bool
	method  bool
	path    bool
	queries bool
}

// Decode an array of KeyMountedValues
func (ea *ExtensionAttributesFrom) Decode(value string) error {
	for _, o := range strings.Split(value, ",") {
		switch o {
		case "method":
			ea.method = true
		case "path":
			ea.path = true
		case "host":
			ea.host = true
		case "queries":
			ea.queries = true
		case "headers":
			ea.headers = true
		default:
			return fmt.Errorf("CloudEvent extension from HTTP element not supported: %s", o)
		}
	}

	return nil
}
