package fake

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

type contextKey int

const fakeAPIKey = "405ad447-1784-45e3-8760-1cf5d1a3c2ae"
const (
	AuthCtx contextKey = iota
)

// NewAuthenticator returns a new authenticator.
func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return Authenticate
}

// Authenticate is a fake implementation of the authenticator.
func Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	c := middleware.GetFiberContext(ctx)

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

	usrCtx := c.UserContext()
	authCtx := context.WithValue(usrCtx, AuthCtx, key)

	// nolint:contextcheck
	c.SetUserContext(authCtx)

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
