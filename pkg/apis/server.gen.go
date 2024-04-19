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
	// List all teams
	// (GET /teams)
	ListTeams(c *fiber.Ctx, params ListTeamsParams) error
	// Creates a new team
	// (POST /teams)
	CreateTeam(c *fiber.Ctx) error
	// Gets a team by ID
	// (GET /teams/{teamId})
	GetTeam(c *fiber.Ctx, teamId TeamId) error
	// List all accounts for a team
	// (GET /teams/{teamId}/accounts)
	ListAccounts(c *fiber.Ctx, teamId TeamId, params ListAccountsParams) error
	// Gets an account by ID
	// (GET /teams/{teamId}/accounts/{accountId})
	GetAccount(c *fiber.Ctx, teamId openapi_types.UUID, accountId openapi_types.UUID) error
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

// ListTeams operation middleware
func (siw *ServerInterfaceWrapper) ListTeams(c *fiber.Ctx) error {

	var err error

	c.Context().SetUserValue(BearerAuthScopes, []string{"read:teams"})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListTeamsParams

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

	return siw.Handler.ListTeams(c, params)
}

// CreateTeam operation middleware
func (siw *ServerInterfaceWrapper) CreateTeam(c *fiber.Ctx) error {

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.CreateTeam(c)
}

// GetTeam operation middleware
func (siw *ServerInterfaceWrapper) GetTeam(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "teamId" -------------
	var teamId TeamId

	err = runtime.BindStyledParameterWithOptions("simple", "teamId", c.Params("teamId"), &teamId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter teamId: %w", err).Error())
	}

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.GetTeam(c, teamId)
}

// ListAccounts operation middleware
func (siw *ServerInterfaceWrapper) ListAccounts(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "teamId" -------------
	var teamId TeamId

	err = runtime.BindStyledParameterWithOptions("simple", "teamId", c.Params("teamId"), &teamId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter teamId: %w", err).Error())
	}

	c.Context().SetUserValue(BearerAuthScopes, []string{"read:accounts"})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListAccountsParams

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

	return siw.Handler.ListAccounts(c, teamId, params)
}

// GetAccount operation middleware
func (siw *ServerInterfaceWrapper) GetAccount(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "teamId" -------------
	var teamId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "teamId", c.Params("teamId"), &teamId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter teamId: %w", err).Error())
	}

	// ------------- Path parameter "accountId" -------------
	var accountId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "accountId", c.Params("accountId"), &accountId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter accountId: %w", err).Error())
	}

	c.Context().SetUserValue(CookieAuthScopes, []string{})

	c.Context().SetUserValue(BearerAuthScopes, []string{})

	c.Context().SetUserValue(Api_keyScopes, []string{})

	return siw.Handler.GetAccount(c, teamId, accountId)
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

	router.Get(options.BaseURL+"/teams", wrapper.ListTeams)

	router.Post(options.BaseURL+"/teams", wrapper.CreateTeam)

	router.Get(options.BaseURL+"/teams/:teamId", wrapper.GetTeam)

	router.Get(options.BaseURL+"/teams/:teamId/accounts", wrapper.ListAccounts)

	router.Get(options.BaseURL+"/teams/:teamId/accounts/:accountId", wrapper.GetAccount)

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

type ListTeamsRequestObject struct {
	Params ListTeamsParams
}

type ListTeamsResponseObject interface {
	VisitListTeamsResponse(ctx *fiber.Ctx) error
}

type ListTeams200JSONResponse struct {
	Limit   *float32 `json:"limit,omitempty"`
	Offset  *float32 `json:"offset,omitempty"`
	Results *[]Team  `json:"results,omitempty"`
	Total   *float32 `json:"total,omitempty"`
}

func (response ListTeams200JSONResponse) VisitListTeamsResponse(ctx *fiber.Ctx) error {
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

type GetTeamRequestObject struct {
	TeamId TeamId `json:"teamId"`
}

type GetTeamResponseObject interface {
	VisitGetTeamResponse(ctx *fiber.Ctx) error
}

type GetTeam200JSONResponse Team

func (response GetTeam200JSONResponse) VisitGetTeamResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type ListAccountsRequestObject struct {
	TeamId TeamId `json:"teamId"`
	Params ListAccountsParams
}

type ListAccountsResponseObject interface {
	VisitListAccountsResponse(ctx *fiber.Ctx) error
}

type ListAccounts200JSONResponse struct {
	Limit   *float32   `json:"limit,omitempty"`
	Offset  *float32   `json:"offset,omitempty"`
	Results *[]Account `json:"results,omitempty"`
	Total   *float32   `json:"total,omitempty"`
}

func (response ListAccounts200JSONResponse) VisitListAccountsResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetAccountRequestObject struct {
	TeamId    openapi_types.UUID `json:"teamId"`
	AccountId openapi_types.UUID `json:"accountId"`
}

type GetAccountResponseObject interface {
	VisitGetAccountResponse(ctx *fiber.Ctx) error
}

type GetAccount200JSONResponse Account

func (response GetAccount200JSONResponse) VisitGetAccountResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

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
	// List all teams
	// (GET /teams)
	ListTeams(ctx context.Context, request ListTeamsRequestObject) (ListTeamsResponseObject, error)
	// Creates a new team
	// (POST /teams)
	CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error)
	// Gets a team by ID
	// (GET /teams/{teamId})
	GetTeam(ctx context.Context, request GetTeamRequestObject) (GetTeamResponseObject, error)
	// List all accounts for a team
	// (GET /teams/{teamId}/accounts)
	ListAccounts(ctx context.Context, request ListAccountsRequestObject) (ListAccountsResponseObject, error)
	// Gets an account by ID
	// (GET /teams/{teamId}/accounts/{accountId})
	GetAccount(ctx context.Context, request GetAccountRequestObject) (GetAccountResponseObject, error)
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

// ListTeams operation middleware
func (sh *strictHandler) ListTeams(ctx *fiber.Ctx, params ListTeamsParams) error {
	var request ListTeamsRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ListTeams(ctx.UserContext(), request.(ListTeamsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListTeams")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ListTeamsResponseObject); ok {
		if err := validResponse.VisitListTeamsResponse(ctx); err != nil {
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

// GetTeam operation middleware
func (sh *strictHandler) GetTeam(ctx *fiber.Ctx, teamId TeamId) error {
	var request GetTeamRequestObject

	request.TeamId = teamId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetTeam(ctx.UserContext(), request.(GetTeamRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTeam")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetTeamResponseObject); ok {
		if err := validResponse.VisitGetTeamResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ListAccounts operation middleware
func (sh *strictHandler) ListAccounts(ctx *fiber.Ctx, teamId TeamId, params ListAccountsParams) error {
	var request ListAccountsRequestObject

	request.TeamId = teamId
	request.Params = params

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.ListAccounts(ctx.UserContext(), request.(ListAccountsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListAccounts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(ListAccountsResponseObject); ok {
		if err := validResponse.VisitListAccountsResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetAccount operation middleware
func (sh *strictHandler) GetAccount(ctx *fiber.Ctx, teamId openapi_types.UUID, accountId openapi_types.UUID) error {
	var request GetAccountRequestObject

	request.TeamId = teamId
	request.AccountId = accountId

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetAccount(ctx.UserContext(), request.(GetAccountRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAccount")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetAccountResponseObject); ok {
		if err := validResponse.VisitGetAccountResponse(ctx); err != nil {
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
