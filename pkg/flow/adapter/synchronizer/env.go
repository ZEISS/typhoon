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
	CorrelationKeyLength int           `envconfig:"CORRELATION_KEY_LENGTH"`
	ResponseWaitTimeout  time.Duration `envconfig:"RESPONSE_WAIT_TIMEOUT"`

	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
}
