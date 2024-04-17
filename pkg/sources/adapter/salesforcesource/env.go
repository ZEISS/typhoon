package salesforcesource

import "knative.dev/eventing/pkg/adapter/v2"

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig

	ClientID   string `envconfig:"SALESFORCE_AUTH_CLIENT_ID" required:"true"`
	AuthServer string `envconfig:"SALESFORCE_AUTH_SERVER" required:"true"`
	User       string `envconfig:"SALESFORCE_AUTH_USER" required:"true"`
	CertKey    string `envconfig:"SALESFORCE_AUTH_CERT_KEY" required:"true"`
	Version    string `envconfig:"SALESFORCE_API_VERSION" default:"48.0"`

	// We are supporting only one subscription + replayID per source instance
	SubscriptionChannel  string `envconfig:"SALESFORCE_SUBCRIPTION_CHANNEL" required:"true"`
	SubscriptionReplayID int    `envconfig:"SALESFORCE_SUBCRIPTION_REPLAY_ID" default:"-1"`
}
