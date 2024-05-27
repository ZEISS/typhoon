package db

import (
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

var _ ports.Repository = (*database)(nil)

type database struct {
	conn *gorm.DB
}

// NewDB returns a new instance of db.
func NewDB(conn *gorm.DB) *database {
	return &database{conn}
}

// RunMigrations runs the database migrations.
func (d *database) RunMigrations() error {
	return d.conn.AutoMigrate(
		&adapters.GothUser{},
		&adapters.GothAccount{},
		&adapters.GothSession{},
		&adapters.GothVerificationToken{},
		&models.User{},
		&models.Account{},
		&models.Operator{},
		&models.System{},
		&models.Tag{},
		&models.Cluster{},
		&models.Token{},
		&models.SigningKeyGroup{},
	)
}
