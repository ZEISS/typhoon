package cmd

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/adapters/database"
	"github.com/zeiss/typhoon/internal/accounts/adapters/handlers"
	"github.com/zeiss/typhoon/internal/accounts/controllers"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis/accounts"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	middleware "github.com/oapi-codegen/fiber-middleware"
	"github.com/spf13/cobra"
	authz "github.com/zeiss/fiber-authz"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", cfg.Flags.Addr, "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DatabaseURI, "db-uri", cfg.Flags.DatabaseURI, "Database URI")
	Root.PersistentFlags().StringVar(&cfg.Flags.DatabaseTablePrefix, "db-table-prefix", cfg.Flags.DatabaseTablePrefix, "Database table prefix")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewAccountsSrv(cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*AccountsSrv)(nil)

// AccountsSrv is the server that implements the Noop interface.
type AccountsSrv struct {
	cfg *Config
}

// NewAccountsSrv returns a new instance of NoopSrv.
func NewAccountsSrv(cfg *Config) *AccountsSrv {
	return &AccountsSrv{cfg}
}

// Start starts the server.
func (s *AccountsSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		conn, err := gorm.Open(postgres.Open(s.cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "typhoon_",
			},
		})
		if err != nil {
			return err
		}

		store, err := seed.NewDatabase(conn, database.NewReadTx(), database.NewWriteTx())
		if err != nil {
			return err
		}

		swagger, err := openapi.GetSwagger()
		if err != nil {
			return err
		}
		swagger.Servers = nil

		c := fiber.Config{
			ErrorHandler: utils.DefaultErrorHandler,
		}

		app := fiber.New(c)
		app.Use(requestid.New())
		app.Use(logger.New())

		validatorOptions := &middleware.Options{}
		validatorOptions.Options.AuthenticationFunc = authz.NewOpenAPIAuthenticator(authz.WithAuthzChecker(authz.NewFake(true)))
		validatorOptions.ErrorHandler = authz.NewOpenAPIErrorHandler()

		app.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

		ac := controllers.NewAccountsController(store)
		handlers := handlers.NewAccountsHandler(ac)

		handler := openapi.NewStrictHandler(handlers, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
