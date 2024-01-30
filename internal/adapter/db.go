package adapter

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/ports"

	openapi_types "github.com/oapi-codegen/runtime/types"
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
func (db *DB) CreateTeam(ctx context.Context, team *openapi.Team) (openapi.Team, error) {
	t := *team
	err := db.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return t, err
	}

	return t, nil
}

// ListTeams ...
func (db *DB) ListTeams(ctx context.Context) ([]openapi.Team, error) {
	var users []openapi.Team
	result := db.conn.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// GetTeam
func (db *DB) GetTeamByID(ctx context.Context, id openapi_types.UUID) (openapi.Team, error) {
	var user openapi.Team
	result := db.conn.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
