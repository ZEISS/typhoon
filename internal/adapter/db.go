package adapter

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
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
func (db *DB) CreateTeam(ctx context.Context, team *openapi.Team) error {
	return db.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(team).Error; err != nil {
			return err
		}

		return nil
	})
}

// ListTeams ...
func (db *DB) ListTeams(ctx context.Context) ([]*openapi.Team, error) {
	var users []*openapi.Team
	result := db.conn.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
