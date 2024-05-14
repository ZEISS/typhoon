package adapters

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/ports"
	api "github.com/zeiss/typhoon/internal/api/models"

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
func (db *DB) GetToken(ctx context.Context, token *api.Token) error {
	return db.conn.First(token).Error
}
