package services

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/adapter"
	"github.com/zeiss/typhoon/internal/api/config"
	openapi "github.com/zeiss/typhoon/pkg/apis"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

var (
	_ server.Listener               = (*WebSrv)(nil)
	_ openapi.StrictServerInterface = (*WebSrv)(nil)
)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *config.Config
	*adapter.Handlers
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *config.Config, handlers *adapter.Handlers) *WebSrv {
	return &WebSrv{cfg, handlers}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
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

		handler := openapi.NewStrictHandler(s, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
