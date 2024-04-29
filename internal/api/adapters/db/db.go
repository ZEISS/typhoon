package db

import (
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"

	"gorm.io/gorm"
)

var (
	_ ports.Users     = (*DB)(nil)
	_ ports.Accounts  = (*DB)(nil)
	_ ports.Operators = (*DB)(nil)
	_ ports.Systems   = (*DB)(nil)
)

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
		&models.User{},
		&models.Account{},
		&models.Operator{},
		&models.System{},
	)
}
