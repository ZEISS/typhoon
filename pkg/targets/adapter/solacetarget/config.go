

package solacetarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	SASLEnable bool `envconfig:"SASL_ENABLE" required:"false"`
	TLSEnable  bool `envconfig:"TLS_ENABLE" required:"false"`

	URL       string `envconfig:"URL" required:"true"`
	QueueName string `envconfig:"QUEUE_NAME" required:"true"`
	Username  string `envconfig:"USERNAME" required:"false"`
	Password  string `envconfig:"PASSWORD" required:"false"`

	CA         string `envconfig:"CA" required:"false"`
	ClientCert string `envconfig:"CLIENT_CERT" required:"false"`
	ClientKey  string `envconfig:"CLIENT_KEY" required:"false"`
	SkipVerify bool   `envconfig:"SKIP_VERIFY" required:"false"`

	DiscardCEContext bool `envconfig:"DISCARD_CE_CONTEXT"`
}
