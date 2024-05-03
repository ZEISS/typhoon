package fake

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

type contextKey int

const fakeAPIKey = "405ad447-1784-45e3-8760-1cf5d1a3c2ae"
const (
	AuthCtx contextKey = iota
)

func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(middleware.GetFiberContext(ctx), input)
	}
}

// Authenticate is a fake implementation of the authenticator.
func Authenticate(ctx *fiber.Ctx, input *openapi3filter.AuthenticationInput) error {
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

	usrCtx := ctx.UserContext()
	authCtx := context.WithValue(usrCtx, AuthCtx, key)

	ctx.SetUserContext(authCtx)

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
