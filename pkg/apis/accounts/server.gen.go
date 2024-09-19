// Package accounts provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package accounts

import (
	"context"
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// This is a test endpoint
	// (GET /accounts/)
	GetHelp(c *fiber.Ctx) error
	// Get account information
	// (GET /accounts/{pubKey})
	GetAccountToken(c *fiber.Ctx, pubKey PubKey) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// GetHelp operation middleware
func (siw *ServerInterfaceWrapper) GetHelp(c *fiber.Ctx) error {

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(ApiKeyScopes, []string{})

	return siw.Handler.GetHelp(c)
}

// GetAccountToken operation middleware
func (siw *ServerInterfaceWrapper) GetAccountToken(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "pubKey" -------------
	var pubKey PubKey

	err = runtime.BindStyledParameterWithOptions("simple", "pubKey", c.Params("pubKey"), &pubKey, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter pubKey: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(ApiKeyScopes, []string{})

	return siw.Handler.GetAccountToken(c, pubKey)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Get(options.BaseURL+"/accounts/", wrapper.GetHelp)

	router.Get(options.BaseURL+"/accounts/:pubKey", wrapper.GetAccountToken)

}

type GetHelpRequestObject struct {
}

type GetHelpResponseObject interface {
	VisitGetHelpResponse(ctx *fiber.Ctx) error
}

type GetHelp200Response struct {
}

func (response GetHelp200Response) VisitGetHelpResponse(ctx *fiber.Ctx) error {
	ctx.Status(200)
	return nil
}

type GetHelpdefaultResponse struct {
	StatusCode int
}

func (response GetHelpdefaultResponse) VisitGetHelpResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

type GetAccountTokenRequestObject struct {
	PubKey PubKey `json:"pubKey"`
}

type GetAccountTokenResponseObject interface {
	VisitGetAccountTokenResponse(ctx *fiber.Ctx) error
}

type GetAccountToken200ResponseHeaders struct {
	CacheControl string
	ETag         string
}

type GetAccountToken200ApplicationjwtResponse struct {
	Body          io.Reader
	Headers       GetAccountToken200ResponseHeaders
	ContentLength int64
}

func (response GetAccountToken200ApplicationjwtResponse) VisitGetAccountTokenResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Cache-Control", fmt.Sprint(response.Headers.CacheControl))
	ctx.Response().Header.Set("ETag", fmt.Sprint(response.Headers.ETag))
	ctx.Response().Header.Set("Content-Type", "application/jwt")
	if response.ContentLength != 0 {
		ctx.Response().Header.Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	ctx.Status(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(ctx.Response().BodyWriter(), response.Body)
	return err
}

type GetAccountToken304Response struct {
}

func (response GetAccountToken304Response) VisitGetAccountTokenResponse(ctx *fiber.Ctx) error {
	ctx.Status(304)
	return nil
}

type GetAccountToken404Response struct {
}

func (response GetAccountToken404Response) VisitGetAccountTokenResponse(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return nil
}

type GetAccountTokendefaultResponse struct {
	StatusCode int
}

func (response GetAccountTokendefaultResponse) VisitGetAccountTokenResponse(ctx *fiber.Ctx) error {
	ctx.Status(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// This is a test endpoint
	// (GET /accounts/)
	GetHelp(ctx context.Context, request GetHelpRequestObject) (GetHelpResponseObject, error)
	// Get account information
	// (GET /accounts/{pubKey})
	GetAccountToken(ctx context.Context, request GetAccountTokenRequestObject) (GetAccountTokenResponseObject, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetHelp operation middleware
func (sh *strictHandler) GetHelp(ctx *fiber.Ctx) error {
	var request GetHelpRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetHelp(ctx.UserContext(), request.(GetHelpRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHelp")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetHelpResponseObject); ok {
		if err := validResponse.VisitGetHelpResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetAccountToken operation middleware
func (sh *strictHandler) GetAccountToken(ctx *fiber.Ctx, pubKey PubKey) error {
	var request GetAccountTokenRequestObject

	request.PubKey = pubKey

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetAccountToken(ctx.UserContext(), request.(GetAccountTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAccountToken")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetAccountTokenResponseObject); ok {
		if err := validResponse.VisitGetAccountTokenResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}