package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
	"github.com/zeiss/fiber-htmx/components/toasts"
	"github.com/zeiss/typhoon/internal/web/adapters/db"
	"github.com/zeiss/typhoon/internal/web/adapters/handlers"
	"github.com/zeiss/typhoon/static"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	openfga "github.com/openfga/go-sdk/client"
	"github.com/spf13/cobra"
	goth "github.com/zeiss/fiber-goth"
	adapter "github.com/zeiss/fiber-goth/adapters/gorm"
	"github.com/zeiss/pkg/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := envconfig.Process("", cfg.Flags)
	if err != nil {
		log.Fatal(err)
	}

	Root.AddCommand(Migrate)

	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", cfg.Flags.Addr, "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DatabaseURI, "db-uri", cfg.Flags.DatabaseURI, "Database URI")
	Root.PersistentFlags().StringVar(&cfg.Flags.DatabaseTablePrefix, "db-table-prefix", cfg.Flags.DatabaseTablePrefix, "Database table prefix")
	Root.PersistentFlags().StringVar(&cfg.Flags.FGAApiUrl, "fga-api-url", cfg.Flags.FGAApiUrl, "FGA API URL")
	Root.PersistentFlags().StringVar(&cfg.Flags.FGAStoreID, "fga-store-id", cfg.Flags.FGAStoreID, "FGA Store ID")
	Root.PersistentFlags().StringVar(&cfg.Flags.FGAAuthorizationModelID, "fga-authorization-model-id", cfg.Flags.FGAAuthorizationModelID, "FGA Authorization Model ID")
	Root.PersistentFlags().StringVar(&cfg.Flags.GothGitbubKey, "github-key", cfg.Flags.GothGitbubKey, "GitHub Key")
	Root.PersistentFlags().StringVar(&cfg.Flags.GothGithubSecret, "github-secret", cfg.Flags.GothGithubSecret, "GitHub Secret")
	Root.PersistentFlags().StringVar(&cfg.Flags.GothGithubCallback, "github-callback", cfg.Flags.GothGithubCallback, "GitHub Callback")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := envconfig.Process("", cfg.Flags)
		if err != nil {
			return err
		}

		return nil
	},
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
		providers.RegisterProvider(github.New(cfg.Flags.GothGitbubKey, cfg.Flags.GothGithubSecret, cfg.Flags.GothGithubCallback))

		conn, err := gorm.Open(postgres.Open(cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: cfg.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		fga, err := openfga.NewSdkClient(
			&openfga.ClientConfiguration{
				ApiUrl:               cfg.Flags.FGAApiUrl,
				StoreId:              cfg.Flags.FGAStoreID,
				AuthorizationModelId: cfg.Flags.FGAAuthorizationModelID,
			},
		)
		if err != nil {
			return err
		}

		auth := authz.NewFGA(fga)

		store, err := db.NewDB(conn)
		if err != nil {
			return err
		}

		ga := adapter.New(conn)

		gothConfig := goth.Config{
			Adapter:        ga,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
		}

		handlers := handlers.NewHandlers(store, auth)

		app := fiber.New(
			fiber.Config{
				ErrorHandler: toasts.DefaultErrorHandler,
			},
		)
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use("/static", filesystem.New(filesystem.Config{
			Root: http.FS(static.Assets),
		}))

		app.Use(goth.NewProtectMiddleware(gothConfig))

		app.Get("/login", handlers.Login())
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))

		// Root handler
		dashboard := app.Group("/")
		dashboard.Get("/", handlers.Dashboard())
		dashboard.Get("/stats/accounts", handlers.GetTotalNumberOfAccounts())
		dashboard.Get("/stats/operators", handlers.GetTotalNumberOfOperators())
		dashboard.Get("/stats/systems", handlers.GetTotalNumberOfSystems())
		dashboard.Get("/stats/users", handlers.GetTotalNumberOfUsers())
		dashboard.Get("/stats/teams", handlers.GetTotalNumberOfTeams())

		// Site handler
		teams := app.Group("/teams")
		teams.Get("/", handlers.ListTeams())
		teams.Get("/new", handlers.NewTeam())
		teams.Post("/new", handlers.CreateTeam())
		teams.Get("/:id", handlers.ShowTeam())
		teams.Delete("/:id", handlers.DeleteTeam())
		teams.Get("/:id/edit", handlers.EditTeam())
		teams.Post("/:id/edit", handlers.UpdateTeam())

		// Me handler
		app.Get("/me", handlers.Me())

		// Operators handler
		app.Get("/operators", handlers.ListOperators())
		app.Get("/operators/new", handlers.NewOperator())
		app.Post("/operators/new", handlers.CreateOperator())
		app.Get("/operators/:id", handlers.ShowOperator())
		app.Delete("/operators/:id", handlers.DeleteOperator())
		app.Get("/operators/:id/token", handlers.TokenOperator())
		app.Get("/operators/:id/skgs/new", handlers.NewOperatorSkg())
		app.Post("/operators/:id/skgs/create", handlers.CreateOperatorSkg())

		// Users handler
		app.Get("/users", handlers.ListUsers())
		app.Get("/users/new", handlers.NewUser())
		app.Post("/users/create", handlers.CreateUser())
		app.Get("/users/:id", handlers.ShowUser())
		app.Delete("/users/:id", handlers.DeleteUser())
		app.Get("/users/partials/account-skgs", handlers.AccountSksOptions())
		app.Get("/users/:id/credentials", handlers.UserCredentials())

		// Systems handler
		app.Get("/systems", handlers.ListSystems())
		app.Post("/systems", handlers.CreateSystem())
		app.Get("/systems/new", handlers.NewSystem())
		app.Get("/systems/:id", handlers.ShowSystem())
		app.Delete("/systems/:id", handlers.DeleteSystem())

		// Accounts handler
		app.Get("/accounts", handlers.ListAccounts())
		app.Get("/accounts/new", handlers.NewAccount())
		app.Post("/accounts/create", handlers.CreateAccount())
		app.Get("/accounts/search/teams", handlers.AccountTeamSearch())
		app.Get("/accounts/search/operators", handlers.OperatorSearch())
		app.Get("/accounts/:id", handlers.ShowAccount())
		app.Delete("/accounts/:id", handlers.DeleteAccount())
		app.Get("/accounts/:id/token", handlers.GetAccountToken())
		app.Get("/accounts/partials/operator-skgs", handlers.OperatorSkgsOptions())

		// Teams handler
		// team := app.Group("/teams/:team_id")

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
