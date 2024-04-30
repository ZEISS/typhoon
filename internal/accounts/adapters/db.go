package adapters

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/models"
	"github.com/zeiss/typhoon/internal/accounts/ports"
	api "github.com/zeiss/typhoon/internal/api/models"

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
func (db *DB) GetToken(ctx context.Context, pubKey models.AccountPublicKey) (models.AccountToken, error) {
	var token api.Token

	if err := db.conn.Where("id = ?", pubKey).First(&token).Error; err != nil {
		return models.AccountToken(""), err
	}

	return models.AccountToken(token.Token), nil
}
