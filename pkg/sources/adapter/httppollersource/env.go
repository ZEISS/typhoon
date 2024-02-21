package httppollersource

import (
	"time"

	"knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig

	EventType         string            `envconfig:"HTTPPOLLER_EVENT_TYPE" required:"true"`
	EventSource       string            `envconfig:"HTTPPOLLER_EVENT_SOURCE" required:"true"`
	Endpoint          string            `envconfig:"HTTPPOLLER_ENDPOINT" required:"true"`
	Method            string            `envconfig:"HTTPPOLLER_METHOD" required:"true"`
	SkipVerify        bool              `envconfig:"HTTPPOLLER_SKIP_VERIFY"`
	CACertificate     string            `envconfig:"HTTPPOLLER_CA_CERTIFICATE"`
	BasicAuthUsername string            `envconfig:"HTTPPOLLER_BASICAUTH_USERNAME"`
	BasicAuthPassword string            `envconfig:"HTTPPOLLER_BASICAUTH_PASSWORD"`
	Headers           map[string]string `envconfig:"HTTPPOLLER_HEADERS"`
	Interval          time.Duration     `envconfig:"HTTPPOLLER_INTERVAL" required:"true"`
}
