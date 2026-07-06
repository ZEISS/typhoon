package synchronizer

import (
	"time"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig
	CorrelationKey       string        `envconfig:"CORRELATION_KEY"`
	BridgeIdentifier     string        `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	CorrelationKeyLength int           `envconfig:"CORRELATION_KEY_LENGTH"`
	ResponseWaitTimeout  time.Duration `envconfig:"RESPONSE_WAIT_TIMEOUT"`
}
