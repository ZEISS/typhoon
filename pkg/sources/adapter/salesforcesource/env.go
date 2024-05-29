package salesforcesource

import "knative.dev/eventing/pkg/adapter/v2"

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig

	// Salesforce configuration for the server-to-server integration using
	// the OAuth2 client credentials flow.
	// See: https://help.salesforce.com/s/articleView?id=release-notes.rn_security_client_credentials_flow.htm
	TokenURL     string `envconfig:"SALESFORCE_TOKEN_URL" required:"true"`
	ClientID     string `envconfig:"SALESFORCE_CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"SALESFORCE_CLIENT_SECRET" required:"true"`
	InstanceURL  string `envconfig:"SALESFORCE_URL" required:"true"`
	Version      string `envconfig:"SALESFORCE_API_VERSION" default:"48.0"`

	// We are supporting only one subscription + replayID per source instance
	SubscriptionChannel  string `envconfig:"SALESFORCE_SUBCRIPTION_CHANNEL" required:"true"`
	SubscriptionReplayID int    `envconfig:"SALESFORCE_SUBCRIPTION_REPLAY_ID" default:"-1"`
}
