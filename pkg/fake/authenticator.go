package fake

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
)

const fakeAPIKey = "405ad447-1784-45e3-8760-1cf5d1a3c2ae"

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(ctx, input)
	}
}

// Authenticate is a fake implementation of the authenticator.
func Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	fmt.Println(input.SecuritySchemeName)

	if input.SecuritySchemeName != "apiKey" {
		return fmt.Errorf("security scheme %s != 'apiKey'", input.SecuritySchemeName)
	}

	key, err := GetAPIKeyFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return err
	}

	if key != fakeAPIKey {
		return fmt.Errorf("invalid API key")
	}

	return nil
}

// GetAPIKeyFromRequest is a fake implementation of the API key extractor.
func GetAPIKeyFromRequest(req *http.Request) (string, error) {
	key := req.Header.Get("X-API-Key")

	if key != fakeAPIKey {
		return "", fmt.Errorf("invalid API key")
	}

	return fakeAPIKey, nil
}
