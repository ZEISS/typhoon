package config

import (
	"github.com/zeiss/typhoon/internal/models"
	"gorm.io/gorm"
)

// RunMigrations ...
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Role{},
		&models.Team{},
		&models.User{},
		&models.UserRole{},
	)
}
