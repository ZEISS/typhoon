package kafkasource

import (
	"knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig satisfies pkgadapter.EnvConfigConstructor.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	adapter.EnvConfig

	SASLEnable bool `envconfig:"SASL_ENABLE" required:"false"`
	TLSEnable  bool `envconfig:"TLS_ENABLE" required:"false"`

	BootstrapServers []string `envconfig:"BOOTSTRAP_SERVERS" required:"true"`
	Username         string   `envconfig:"USERNAME" required:"false"`
	Password         string   `envconfig:"PASSWORD" required:"false"`
	Topic            string   `envconfig:"TOPIC" required:"true"`
	GroupID          string   `envconfig:"GROUP_ID" required:"false"`

	SecurityMechanisms  string `envconfig:"SECURITY_MECHANISMS" required:"false"`
	KerberosConfigPath  string `envconfig:"KERBEROS_CONFIG_PATH" required:"false" `
	KerberosServiceName string `envconfig:"KERBEROS_SERVICE_NAME" required:"false" `
	KerberosKeytabPath  string `envconfig:"KERBEROS_KEYTAB_PATH" required:"false"`
	KerberosRealm       string `envconfig:"KERBEROS_REALM" required:"false"`
	KerberosUsername    string `envconfig:"KERBEROS_USERNAME" required:"false"`
	KerberosPassword    string `envconfig:"KERBEROS_PASSWORD" required:"false"`

	CA         string `envconfig:"CA" required:"false"`
	ClientCert string `envconfig:"CLIENT_CERT" required:"false"`
	ClientKey  string `envconfig:"CLIENT_KEY" required:"false"`
	SkipVerify bool   `envconfig:"SKIP_VERIFY" required:"false"`
}
