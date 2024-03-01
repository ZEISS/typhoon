package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/typhoon/pkg/brokers/backend/impl/redis"
	"github.com/zeiss/typhoon/pkg/brokers/broker"

	"github.com/katallaxie/pkg/logger"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	Addr string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":3000", "addr")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	// rootCmd.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")

	rootCmd.SilenceUsage = true
}

type srv struct{}

func (b *srv) Start(ctx context.Context, _ server.ReadyFunc, _ server.RunFunc) func() error {
	return func() error {
		l, err := logger.NewLogSink()
		if err != nil {
			return err
		}

		b := redis.New(nil, l.Sugar())

		s, err := broker.NewInstance(nil, b)
		if err != nil {
			return err
		}

		return s.Start(ctx)
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	broker := &srv{}

	srv, _ := server.WithContext(ctx)
	srv.Listen(broker, false)

	err := srv.Wait()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
