package adapter

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
	openapi "github.com/zeiss/typhoon/pkg/apis"

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
	err := db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
func (db *DB) ListTeams(ctx context.Context, params openapi.ListTeamParams) (models.PaginatedListTeams, error) {
	var users []openapi.Team
	var totalRows int64
	db.conn.Find(&users).WithContext(ctx).Count(&totalRows)

	result := db.conn.WithContext(ctx).Limit(*params.Limit).Offset(*params.Offset).Find(&users)
	if result.Error != nil {
		return models.PaginatedListTeams{}, result.Error
	}

	return models.PaginatedListTeams{Results: &users}, nil
}

// GetTeam
func (db *DB) GetTeamByID(ctx context.Context, id openapi_types.UUID) (openapi.Team, error) {
	var user openapi.Team
	result := db.conn.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
