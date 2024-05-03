package db

import (
	authz "github.com/zeiss/fiber-authz"
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
		&models.Team{},
		&authz.User{},
		&authz.Role{},
		&authz.Permission{},
		&authz.UserRole{},
		&adapters.Account{},
		&adapters.Session{},
		&models.User{},
		&models.Account{},
		&models.Operator{},
		&models.System{},
		&models.Tag{},
		&models.Ownership{},
	)
}
