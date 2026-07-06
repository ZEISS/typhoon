//go:build !noclibs

package xslttransformation

import (
	"errors"

	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig
	XSLT              string `envconfig:"XSLTTRANSFORMATION_XSLT"`
	BridgeIdentifier  string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	Sink              string `envconfig:"K_SINK"`
	AllowXSLTOverride bool   `envconfig:"XSLTTRANSFORMATION_ALLOW_XSLT_OVERRIDE" required:"true"`
}

func (e *envAccessor) validate() error {
	if !e.AllowXSLTOverride && e.XSLT == "" {
		return errors.New("if XSLT cannot be overridden by CloudEvent payloads, configured XSLT cannot be empty")
	}
	return nil
}
