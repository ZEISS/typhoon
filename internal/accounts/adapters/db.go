package adapters

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/models"
	"github.com/zeiss/typhoon/internal/accounts/ports"

	"gorm.io/gorm"
)

var _ ports.Accounts = (*DB)(nil)

// DB ...
type DB struct {
	conn *gorm.DB
}

// NewDB ...
func NewDB(conn *gorm.DB) *DB {
	return &DB{conn}
}

// GetToken ...
func (db *DB) GetToken(ctx context.Context, account models.AccountPublicKey) (models.AccountToken, error) {
	return models.AccountToken(""), nil
}
