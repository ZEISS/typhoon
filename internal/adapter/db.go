package adapter

import (
	"context"

	"github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/ports"

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

// CreateTeam ...
func (db *DB) CreateTeam(ctx context.Context, team *api.Team) error {
	return db.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(team).Error; err != nil {
			return err
		}

		return nil
	})
}
