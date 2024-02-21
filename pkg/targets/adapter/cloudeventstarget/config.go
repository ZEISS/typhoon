package cloudeventstarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	URL                   string `envconfig:"CLOUDEVENTS_URL" required:"true"`
	Path                  string `envconfig:"CLOUDEVENTS_PATH"`
	BasicAuthUsername     string `envconfig:"CLOUDEVENTS_BASICAUTH_USERNAME"`
	BasicAuthPasswordPath string `envconfig:"CLOUDEVENTS_BASICAUTH_PASSWORD_PATH"`
}
