package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/katallaxie/pkg/server"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func init() {
	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")
	Root.PersistentFlags().StringVar(&cfg.Flags.Nats.Credentials, "credentials", cfg.Flags.Nats.Credentials, "credentials file")
	Root.PersistentFlags().StringVar(&cfg.Flags.Nats.Url, "url", cfg.Flags.Nats.Url, "url")

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

var _ server.Listener = (*AccountServ)(nil)

// AccountServ is the server that implements the Noop interface.
type AccountServ struct {
	cfg *Config
}

// NewWebSrv returns a new instance of AccountServ.
func NewWebSrv(cfg *Config) *AccountServ {
	return &AccountServ{cfg}
}

// Start starts the server.
func (s *AccountServ) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		nc, err := nats.Connect(cfg.Flags.Nats.Url, nats.UserCredentials(cfg.Flags.Nats.Credentials))
		if err != nil {
			return err
		}

		sub, err := nc.Subscribe("$SYS.REQ.ACCOUNT.*.CLAIMS.LOOKUP", func(msg *nats.Msg) {
			accountId := strings.TrimSuffix(strings.TrimPrefix(msg.Subject, "$SYS.REQ.ACCOUNT."), ".CLAIMS.LOOKUP")
			log.Println("account lookup", "accountId", accountId)
		})
		if err != nil {
			return err
		}

		<-ctx.Done()

		return sub.Unsubscribe()
	}
}
