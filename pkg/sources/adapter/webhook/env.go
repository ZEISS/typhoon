package webhook

import (
	kadapter "knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig ...
func NewEnvconfig() kadapter.EnvConfigAccessor {
	return &env{}
}

type env struct {
	kadapter.EnvConfig

	BasicAuthPassword            string                   `envconfig:"WEBHOOK_BASICAUTH_PASSWORD"`
	BasicAuthUsername            string                   `envconfig:"WEBHOOK_BASICAUTH_USERNAME"`
	CORSAllowOrigin              string                   `envconfig:"WEBHOOK_CORS_ALLOW_ORIGIN"`
	EventExtensionAttributesFrom *ExtensionAttributesFrom `envconfig:"WEBHOOK_EVENT_EXTENSION_ATTRIBUTES_FROM"`
	EventSource                  string                   `envconfig:"WEBHOOK_EVENT_SOURCE" required:"true"`
	EventType                    string                   `envconfig:"WEBHOOK_EVENT_TYPE" required:"true"`
}

type ExtensionAttributesFrom struct {
	method  bool
	path    bool
	host    bool
	queries bool
	headers bool
}
