package db

import (
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"

	"gorm.io/gorm"
)

var _ ports.Repositories = (*DB)(nil)

// DB ...
type DB struct {
	conn *gorm.DB
}

// NewDB ...
func NewDB(conn *gorm.DB) *DB {
	return &DB{conn}
}

// RunMigrations ...
func (db *DB) RunMigrations() error {
	return db.conn.AutoMigrate(
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
}
