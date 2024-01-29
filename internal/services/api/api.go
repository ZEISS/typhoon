package api

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/adapter"
	"github.com/zeiss/typhoon/internal/config"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
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
	cfg *config.Config
	*adapter.Handlers
}

// New returns a new instance of NoopSrv.
func New(cfg *config.Config, handlers *adapter.Handlers) *ApiSrv {
	return &ApiSrv{cfg, handlers}
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

		validatorOptions := &middleware.Options{}
		validatorOptions.Options.AuthenticationFunc = func(ctx context.Context, filter *openapi3filter.AuthenticationInput) error {
			return nil
		}

		app.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

		handler := openapi.NewStrictHandler(a, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(a.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
