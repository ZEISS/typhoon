package awss3target

import (
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
)

// NewEnvConfig for configuration parameters
func NewEnvConfig() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	pkgadapter.EnvConfig

	AwsTargetArn string `envconfig:"ARN" required:"true"`

	DiscardCEContext bool `envconfig:"AWS_DISCARD_CE_CONTEXT"`

	// Assume this IAM Role when access keys provided.
	AssumeIamRole string `envconfig:"AWS_ASSUME_ROLE_ARN"`

	// The environment variables below aren't read from the envConfig struct
	// by the AWS SDK, but rather directly using os.Getenv().
	// They are nevertheless listed here for documentation purposes.
	_ string `envconfig:"AWS_ACCESS_KEY_ID"`
	_ string `envconfig:"AWS_SECRET_ACCESS_KEY"`
}
