package cmd

import (
	"context"
	"net/http"
	"os"

	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
	"github.com/zeiss/typhoon/internal/web/adapters/db"
	"github.com/zeiss/typhoon/internal/web/adapters/handlers"
	"github.com/zeiss/typhoon/static"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"
	authz "github.com/zeiss/fiber-authz"
	goth "github.com/zeiss/fiber-goth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":3000", "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Addr, "db-host", cfg.Flags.DB.Addr, "Database host")

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
		providers.RegisterProvider(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"))

		conn, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "typhoon_",
			},
		})
		if err != nil {
			return err
		}

		db := db.NewDB(conn)
		err = db.RunMigrations()
		if err != nil {
			return err
		}

		tbac := authz.NewTBAC(conn)

		gothConfig := goth.Config{
			Adapter:        tbac,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
			ResponseFilter: func(c *fiber.Ctx) error {
				return c.Redirect("/")
			},
		}

		handlers := handlers.NewHandlers(db)

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use("/static", filesystem.New(filesystem.Config{
			Root: http.FS(static.Assets),
		}))

		app.Use(goth.NewProtectMiddleware(gothConfig))
		app.Use(authz.SetAuthzHandler(authz.NewNoopObjectResolver(), authz.NewNoopActionResolver(), authz.NewGothAuthzPrincipalResolver()))

		app.Get("/login", handlers.Login())
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))

		// Root handler
		app.Get("/", handlers.Dashboard())

		// Me handler
		app.Get("/me", handlers.Me())

		// Operators handler
		app.Get("/operators", handlers.ListOperators())
		app.Get("/operators/new", handlers.NewOperator())
		app.Post("/operators/new", handlers.CreateOperator())
		app.Get("/operators/:id", handlers.ShowOperator())
		app.Delete("/operators/:id", handlers.DeleteOperator())
		app.Get("/operators/:id/skgs/new", handlers.NewOperatorSkg())
		app.Post("/operators/:id/skgs/create", handlers.CreateOperatorSkg())

		// Accounts handler
		app.Get("/accounts", handlers.ListAccounts())
		app.Get("/accounts/new", handlers.NewAccount())
		app.Post("/accounts/create", handlers.CreateAccount())
		app.Get("/accounts/:id", handlers.ShowAccount())
		app.Get("/accounts/partials/operator-skgs", handlers.OperatorSkgsOptions())

		// Users handler
		app.Get("/users", handlers.ListUsers())
		app.Get("/users/new", handlers.NewUser())
		app.Post("/users/create", handlers.CreateUser())
		app.Get("/users/partials/account-skgs", handlers.AccountSksOptions())

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
