package cmd

import (
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/pkg/dbx"
	"github.com/zeiss/typhoon/internal/accounts/adapters/database"
	"github.com/zeiss/typhoon/internal/models"

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

		store, err := dbx.NewDatabase(conn, database.NewReadTx(), database.NewWriteTx())
		if err != nil {
			return err
		}

		return store.Migrate(
			cmd.Context(),
			&adapters.GothAccount{},
			&adapters.GothUser{},
			&adapters.GothSession{},
			&adapters.GothVerificationToken{},
			&models.User{},
			&models.Account{},
			&models.Operator{},
			&models.System{},
			&models.Tag{},
			&models.NKey{},
			&models.Cluster{},
			&models.Token{},
			&models.SigningKeyGroup{},
			&models.UserLimits{},
		)
	},
}
