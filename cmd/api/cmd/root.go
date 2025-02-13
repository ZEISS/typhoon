package cmd

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/adapters/db"
	"github.com/zeiss/typhoon/internal/api/adapters/handlers"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/utils"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	middleware "github.com/oapi-codegen/fiber-middleware"
	"github.com/spf13/cobra"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/pkg/server"
	openapi "github.com/zeiss/typhoon/pkg/apis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	Root.AddCommand(Migrate)

	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Addr, "db-addr", cfg.Flags.DB.Addr, "Database address")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Prefix, "db-prefix", cfg.Flags.DB.Prefix, "Database prefix")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewWebSrv(cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *Config
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *Config) *WebSrv {
	return &WebSrv{cfg}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		conn, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: cfg.Prefix(),
			},
		})
		if err != nil {
			return err
		}

		db := db.NewDB(conn)

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

		tc := controllers.NewTeamsController(db)
		sc := controllers.NewSystemsController(db)
		oc := controllers.NewOperatorsController(db)
		ac := controllers.NewAccountsController(db)
		uc := controllers.NewUsersController(db)

		handlers := handlers.NewApiHandlers(sc, tc, oc, ac, uc)

		handler := openapi.NewStrictHandler(handlers, nil)
		openapi.RegisterHandlers(app, handler)

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
