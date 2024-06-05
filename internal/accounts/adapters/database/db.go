package database

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/ports"
	"github.com/zeiss/typhoon/internal/api/models"

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

// GetToken ...
func (db *DB) GetToken(ctx context.Context, account *models.Token) error {
	return db.conn.WithContext(ctx).First(account).Error
}
