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
	// XSLT document that will be used by default for transformation.
	XSLT string `envconfig:"XSLTTRANSFORMATION_XSLT"`
	// If set to true, enables consuming structured CloudEvents that include
	// fields for the XML and XSLT field.
	AllowXSLTOverride bool `envconfig:"XSLTTRANSFORMATION_ALLOW_XSLT_OVERRIDE" required:"true"`
	// BridgeIdentifier is the name of the bridge workflow this target is part of
	BridgeIdentifier string `envconfig:"EVENTS_BRIDGE_IDENTIFIER"`
	// Sink defines the target sink for the events. If no Sink is defined the
	// events are replied back to the sender.
	Sink string `envconfig:"K_SINK"`
}

func (e *envAccessor) validate() error {
	if !e.AllowXSLTOverride && e.XSLT == "" {
		return errors.New("if XSLT cannot be overridden by CloudEvent payloads, configured XSLT cannot be empty")
	}
	return nil
}
