// Package apis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package apis

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all managed systems.
	// (GET /systems)
	ListSystems(c *fiber.Ctx) error
	// List all managed systems.
	// (GET /systems/{systemId})
	ShowSystem(c *fiber.Ctx, systemId string) error
	// Gets a team by ID
	// (GET /team/{teamId})
	GetTeamTeamId(c *fiber.Ctx, teamId openapi_types.UUID) error
	// List all teams
	// (GET /teams)
	ListTeam(c *fiber.Ctx, params ListTeamParams) error
	// Creates a new team
	// (POST /teams)
	CreateTeam(c *fiber.Ctx) error
	// Returns the current version of the API.
	// (GET /version)
	Version(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// ListSystems operation middleware
func (siw *ServerInterfaceWrapper) ListSystems(c *fiber.Ctx) error {

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.ListSystems(c)
}

// ShowSystem operation middleware
func (siw *ServerInterfaceWrapper) ShowSystem(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "systemId" -------------
	var systemId string

	err = runtime.BindStyledParameterWithOptions("simple", "systemId", c.Params("systemId"), &systemId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter systemId: %w", err).Error())
	}

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.ShowSystem(c, systemId)
}

// GetTeamTeamId operation middleware
func (siw *ServerInterfaceWrapper) GetTeamTeamId(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "teamId" -------------
	var teamId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "teamId", c.Params("teamId"), &teamId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter teamId: %w", err).Error())
	}

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.GetTeamTeamId(c, teamId)
}

// ListTeam operation middleware
func (siw *ServerInterfaceWrapper) ListTeam(c *fiber.Ctx) error {

	var err error

	c.Context().SetUserValue(BearerAuthScopes, []string{"read:teams"})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListTeamParams

	var query url.Values
	query, err = url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
	}

	return siw.Handler.ListTeam(c, params)
}

// CreateTeam operation middleware
func (siw *ServerInterfaceWrapper) CreateTeam(c *fiber.Ctx) error {

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.CreateTeam(c)
}

// Version operation middleware
func (siw *ServerInterfaceWrapper) Version(c *fiber.Ctx) error {

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.Version(c)
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

	router.Get(options.BaseURL+"/systems", wrapper.ListSystems)

	router.Get(options.BaseURL+"/systems/:systemId", wrapper.ShowSystem)

	router.Get(options.BaseURL+"/team/:teamId", wrapper.GetTeamTeamId)

	router.Get(options.BaseURL+"/teams", wrapper.ListTeam)

	router.Post(options.BaseURL+"/teams", wrapper.CreateTeam)

	router.Get(options.BaseURL+"/version", wrapper.Version)

}

type ListSystemsRequestObject struct {
}

type ListSystemsResponseObject interface {
	VisitListSystemsResponse(ctx *fiber.Ctx) error
}

type ListSystems200JSONResponse Systems

func (response ListSystems200JSONResponse) VisitListSystemsResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type ShowSystemRequestObject struct {
	SystemId string `json:"systemId"`
}

type ShowSystemResponseObject interface {
	VisitShowSystemResponse(ctx *fiber.Ctx) error
}

type ShowSystem200Response struct {
}

func (response ShowSystem200Response) VisitShowSystemResponse(ctx *fiber.Ctx) error {
	ctx.Status(200)
	return nil
}

type GetTeamTeamIdRequestObject struct {
	TeamId openapi_types.UUID `json:"teamId"`
}

type GetTeamTeamIdResponseObject interface {
	VisitGetTeamTeamIdResponse(ctx *fiber.Ctx) error
}

type GetTeamTeamId200JSONResponse Team

func (response GetTeamTeamId200JSONResponse) VisitGetTeamTeamIdResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type ListTeamRequestObject struct {
	Params ListTeamParams
}

type ListTeamResponseObject interface {
	VisitListTeamResponse(ctx *fiber.Ctx) error
}

type ListTeam200JSONResponse struct {
	Limit   *float32 `json:"limit,omitempty"`
	Offset  *float32 `json:"offset,omitempty"`
	Results *[]Team  `json:"results,omitempty"`
	Total   *float32 `json:"total,omitempty"`
}

func (response ListTeam200JSONResponse) VisitListTeamResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type CreateTeamRequestObject struct {
	Body *CreateTeamJSONRequestBody
}

type CreateTeamResponseObject interface {
	VisitCreateTeamResponse(ctx *fiber.Ctx) error
}

type CreateTeam201JSONResponse Team

func (response CreateTeam201JSONResponse) VisitCreateTeamResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(201)

	return ctx.JSON(&response)
}

type VersionRequestObject struct {
}

type VersionResponseObject interface {
	VisitVersionResponse(ctx *fiber.Ctx) error
}

type Version200JSONResponse Version

func (response Version200JSONResponse) VisitVersionResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// List all managed systems.
	// (GET /systems)
	ListSystems(ctx context.Context, request ListSystemsRequestObject) (ListSystemsResponseObject, error)
	// List all managed systems.
	// (GET /systems/{systemId})
	ShowSystem(ctx context.Context, request ShowSystemRequestObject) (ShowSystemResponseObject, error)
	// Gets a team by ID
	// (GET /team/{teamId})
	GetTeamTeamId(ctx context.Context, request GetTeamTeamIdRequestObject) (GetTeamTeamIdResponseObject, error)
	// List all teams
	// (GET /teams)
	ListTeam(ctx context.Context, request ListTeamRequestObject) (ListTeamResponseObject, error)
	// Creates a new team
	// (POST /teams)
	CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error)
	// Returns the current version of the API.
	// (GET /version)
	Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error)
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

// ListSystems operation middleware
func (sh *strictHandler) ListSystems(ctx *fiber.Ctx) error {
	var request ListSystemsRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ListSystems(ctx.UserContext(), request.(ListSystemsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListSystems")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ListSystemsResponseObject); ok {
		if err := validResponse.VisitListSystemsResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ShowSystem operation middleware
func (sh *strictHandler) ShowSystem(ctx *fiber.Ctx, systemId string) error {
	var request ShowSystemRequestObject

	request.SystemId = systemId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ShowSystem(ctx.UserContext(), request.(ShowSystemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ShowSystem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ShowSystemResponseObject); ok {
		if err := validResponse.VisitShowSystemResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTeamTeamId operation middleware
func (sh *strictHandler) GetTeamTeamId(ctx *fiber.Ctx, teamId openapi_types.UUID) error {
	var request GetTeamTeamIdRequestObject

	request.TeamId = teamId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetTeamTeamId(ctx.UserContext(), request.(GetTeamTeamIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTeamTeamId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetTeamTeamIdResponseObject); ok {
		if err := validResponse.VisitGetTeamTeamIdResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ListTeam operation middleware
func (sh *strictHandler) ListTeam(ctx *fiber.Ctx, params ListTeamParams) error {
	var request ListTeamRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ListTeam(ctx.UserContext(), request.(ListTeamRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListTeam")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ListTeamResponseObject); ok {
		if err := validResponse.VisitListTeamResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreateTeam operation middleware
func (sh *strictHandler) CreateTeam(ctx *fiber.Ctx) error {
	var request CreateTeamRequestObject

	var body CreateTeamJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.CreateTeam(ctx.UserContext(), request.(CreateTeamRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateTeam")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(CreateTeamResponseObject); ok {
		if err := validResponse.VisitCreateTeamResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Version operation middleware
func (sh *strictHandler) Version(ctx *fiber.Ctx) error {
	var request VersionRequestObject

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.Version(ctx.UserContext(), request.(VersionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Version")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(VersionResponseObject); ok {
		if err := validResponse.VisitVersionResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}