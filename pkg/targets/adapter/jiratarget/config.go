package jiratarget

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	JiraBasicAuthUser  string `envconfig:"JIRA_BASIC_AUTH_USERNAME" required:"true"`
	JiraBasicAuthToken string `envconfig:"JIRA_BASIC_AUTH_PASSWORD" required:"true"`
	JiraURL            string `envconfig:"JIRA_URL" required:"true"`
}
