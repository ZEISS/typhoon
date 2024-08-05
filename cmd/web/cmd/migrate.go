package cmd

import (
	"github.com/zeiss/typhoon/internal/web/adapters/db"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := gorm.Open(postgres.Open(cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: cfg.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		db, err := db.NewDB(conn)
		if err != nil {
			return err
		}

		err = db.Migrate(cmd.Context())
		if err != nil {
			return err
		}

		return nil
	},
}
