package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zeiss/typhoon/internal/api/adapters/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Migrate = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: cfg.Prefix(),
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

		return nil
	},
}
