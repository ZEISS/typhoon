package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
)

// CreateTeam creates a new team.
func (db *DB) CreateTeam(ctx context.Context, team *models.Team) error {
	return db.conn.WithContext(ctx).Create(team).Error
}

// GetTeam retrieves a team by its ID.
func (db *DB) GetTeam(ctx context.Context, team *models.Team) error {
	return db.conn.WithContext(ctx).First(team).Error
}

// DeleteTeam deletes a team by its ID.
func (db *DB) DeleteTeam(ctx context.Context, team *models.Team) error {
	return db.conn.WithContext(ctx).Delete(team).Error
}

// ListTeams retrieves all teams.
func (db *DB) ListTeams(ctx context.Context, pagination *models.Pagination[models.Team]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Find(&pagination.Rows).Error
}
