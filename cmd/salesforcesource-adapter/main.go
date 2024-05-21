package main

import (
	"github.com/golang-jwt/jwt/v4"

	"knative.dev/eventing/pkg/adapter/v2"

	"github.com/zeiss/typhoon/pkg/sources/adapter/salesforcesource"
)

func main() {
	// JWT package marshals Audience as array even if there is only one element in it. This does not
	// seem to be supported by Salesforce. By setting the following option to false we tell the imported
	// library to marshal single item Audience array as a string.
	jwt.MarshalSingleStringAsArray = false

	adapter.Main("salesforce", salesforcesource.NewEnvConfig, salesforcesource.NewAdapter)
}
