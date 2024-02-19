package poller

import (
	"time"

	kadapter "knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig ...
func NewEnvconfig() kadapter.EnvConfigAccessor {
	return &config{}
}

type config struct {
	kadapter.EnvConfig

	EventType         string            `envconfig:"HTTP_EVENT_TYPE" required:"true"`
	EventSource       string            `envconfig:"HTTP_EVENT_SOURCE" required:"true"`
	URL               string            `envconfig:"HTTP_URL" required:"true"`
	Method            string            `envconfig:"HTTP_METHOD" required:"true"`
	SkipVerify        bool              `envconfig:"HTTP_SKIP_VERIFY"`
	CACertificate     string            `envconfig:"HTTP_CA_CERTIFICATE"`
	BasicAuthUsername string            `envconfig:"HTTP_BASICAUTH_USERNAME"`
	BasicAuthPassword string            `envconfig:"HTTP_BASICAUTH_PASSWORD"`
	Headers           map[string]string `envconfig:"HTTP_HEADERS"`
	Interval          time.Duration     `envconfig:"HTTP_INTERVAL" required:"true"`
}
