package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/typhoon/internal/api/adapter"
	"github.com/zeiss/typhoon/internal/api/config"
	"github.com/zeiss/typhoon/internal/api/services"

	"github.com/katallaxie/pkg/logger"
	"github.com/katallaxie/pkg/server"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = config.New()

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	rootCmd.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")

	rootCmd.SilenceUsage = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	dsn := "host=host.docker.internal user=example password=example dbname=example port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = config.RunMigrations(conn)
	if err != nil {
		return err
	}

	// db := adapter.NewDB(conn)
	srv, _ := server.WithContext(ctx)

	h := adapter.NewHandlers()
	service := services.NewWebSrv(cfg, h)

	srv.Listen(service, true)
	if err := srv.Wait(); err != nil {
		return err
	}

	return nil
}
