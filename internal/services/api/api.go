package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/controllers"

	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

var (
	_ server.Listener               = (*ApiSrv)(nil)
	_ openapi.StrictServerInterface = (*ApiSrv)(nil)
)

// ApiSrv is the server that implements the Noop interface.
type ApiSrv struct {
	ctrl *controllers.Systems
}

// New returns a new instance of NoopSrv.
func New(ctrl *controllers.Systems) *ApiSrv {
	return &ApiSrv{
		ctrl: ctrl,
	}
}

// Start starts the server.
func (a *ApiSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		swagger, err := openapi.GetSwagger()
		if err != nil {
			return err
		}
		swagger.Servers = nil

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())
		app.Use(middleware.OapiRequestValidator(swagger))

		handler := openapi.NewStrictHandler(a, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(":3000")
		if err != nil {
			return err
		}

		return nil
	}
}

// ListSystems ...
func (a *ApiSrv) ListSystems(ctx context.Context, request openapi.ListSystemsRequestObject) (openapi.ListSystemsResponseObject, error) {
	return nil, nil
}

// GetSystem ...
func (a *ApiSrv) ShowSystem(ctx context.Context, request openapi.ShowSystemRequestObject) (openapi.ShowSystemResponseObject, error) {
	return nil, nil
}

// Version ...
func (a *ApiSrv) Version(ctx context.Context, request openapi.VersionRequestObject) (openapi.VersionResponseObject, error) {
	return nil, nil
}
