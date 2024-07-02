package cmd

import (
	"github.com/zeiss/fiber-goth/adapters"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/typhoon/internal/accounts/adapters/database"
	"github.com/zeiss/typhoon/internal/api/models"

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

		store, err := seed.NewDatabase(conn, database.NewReadTx(), database.NewWriteTx())
		if err != nil {
			return err
		}

		return store.Migrate(
			cmd.Context(),
			&adapters.GothAccount{},
			&adapters.GothUser{},
			&adapters.GothSession{},
			&adapters.GothVerificationToken{},
			&adapters.GothTeam{},
			&adapters.GothRole{},
			&models.User{},
			&models.Account{},
			&models.Operator{},
			&models.System{},
			&models.Tag{},
			&models.Cluster{},
			&models.Token{},
			&models.SigningKeyGroup{},
		)
	},
}
