package cmd

import (
	"context"
	"strings"

	"github.com/zeiss/typhoon/internal/accounts/adapters"
	"github.com/zeiss/typhoon/internal/accounts/adapters/handlers"
	"github.com/zeiss/typhoon/internal/accounts/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/katallaxie/pkg/server"
	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var cfg = New()

func init() {
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")
	Root.PersistentFlags().StringVar(&cfg.Flags.Nats.Credentials, "nats-credentials", cfg.Flags.Nats.Credentials, "credentials file")
	Root.PersistentFlags().StringVar(&cfg.Flags.Nats.Url, "nats-url", cfg.Flags.Nats.Url, "url")

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
		srv := NewAccountSrv(cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*AccountSrv)(nil)

// AccountSrv is the server that implements the Noop interface.
type AccountSrv struct {
	cfg *Config
}

// NewAccountSrv returns a new instance of AccountServ.
func NewAccountSrv(cfg *Config) *AccountSrv {
	return &AccountSrv{cfg}
}

// Start starts the server.
func (s *AccountSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		dsn := "host=localhost user=example password=example dbname=example port=5432 sslmode=disable"
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

		db := adapters.NewDB(conn)
		ac := controllers.NewAccountsController(db)
		lh := handlers.NewAccountLookupRequestHandler(ac)

		nc, err := nats.Connect(cfg.Flags.Nats.Url, nats.UserCredentials(cfg.Flags.Nats.Credentials))
		if err != nil {
			return err
		}

		sub, err := nc.SubscribeSync("$SYS.REQ.ACCOUNT.*.CLAIMS.LOOKUP")
		if err != nil {
			return err
		}

		for {
			msg, err := sub.NextMsgWithContext(ctx)
			if err != nil {
				return err
			}

			apk := strings.TrimSuffix(strings.TrimPrefix(msg.Subject, "$SYS.REQ.ACCOUNT."), ".CLAIMS.LOOKUP")
			jwt, err := lh.HandleLookupRequest(ctx, apk)
			if err != nil {
				return err
			}

			err = msg.Respond([]byte(jwt))
			if err != nil {
				return err
			}
		}
	}
}
