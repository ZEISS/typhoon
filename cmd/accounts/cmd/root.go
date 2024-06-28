package cmd

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/adapters/database"
	"github.com/zeiss/typhoon/internal/accounts/adapters/handlers"
	"github.com/zeiss/typhoon/internal/accounts/config"
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
	Root.PersistentFlags().StringVar(&config.Cfg.Flags.Addr, "addr", config.Cfg.Flags.Addr, "addr")
	Root.PersistentFlags().StringVar(&config.Cfg.Flags.DB.Addr, "db-addr", config.Cfg.Flags.DB.Addr, "Database address")
	Root.PersistentFlags().StringVar(&config.Cfg.Flags.DB.Database, "db-database", config.Cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&config.Cfg.Flags.DB.Username, "db-username", config.Cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&config.Cfg.Flags.DB.Password, "db-password", config.Cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&config.Cfg.Flags.DB.Port, "db-port", config.Cfg.Flags.DB.Port, "Database port")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewAccountsSrv(config.Cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*AccountsSrv)(nil)

// AccountsSrv is the server that implements the Noop interface.
type AccountsSrv struct {
	cfg *config.Config
}

// NewAccountsSrv returns a new instance of NoopSrv.
func NewAccountsSrv(cfg *config.Config) *AccountsSrv {
	return &AccountsSrv{cfg}
}

// Start starts the server.
func (s *AccountsSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		conn, err := gorm.Open(postgres.Open(s.cfg.DSN()), &gorm.Config{
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
