package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/typhoon/internal/adapter"
	"github.com/zeiss/typhoon/internal/adapter/handlers"
	"github.com/zeiss/typhoon/internal/config"
	"github.com/zeiss/typhoon/internal/controllers"
	"github.com/zeiss/typhoon/internal/services/api"

	"github.com/katallaxie/pkg/logger"
	"github.com/katallaxie/pkg/server"
	models "github.com/zeiss/typhoon/api"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = config.New()

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

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

	conn.AutoMigrate(&models.Team{})

	db := adapter.NewDB(conn)
	srv, _ := server.WithContext(ctx)

	teamsCtrl := controllers.NewTeamsController(db)

	h := adapter.NewHandlers(
		handlers.NewTeamsHandler(teamsCtrl),
		handlers.NewSystemsHandler(),
		handlers.NewVersionHandler(),
	)

	service := api.New(cfg, h)

	srv.Listen(service, true)
	if err := srv.Wait(); err != nil {
		return err
	}

	return nil
}
