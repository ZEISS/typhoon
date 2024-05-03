package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// CreateTeam creates a new team.
func (db *DB) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if err := db.conn.WithContext(ctx).Create(&team).Error; err != nil {
		return models.Team{}, err
	}

	return team, nil
}

// GetTeam retrieves a team by its ID.
func (db *DB) GetTeam(ctx context.Context, id uuid.UUID) (models.Team, error) {
	team := models.Team{}
	if err := db.conn.WithContext(ctx).Where("id = ?", id).First(&team).Error; err != nil {
		return models.Team{}, err
	}

	return team, nil
}

// DeleteTeam deletes a team by its ID.
func (db *DB) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.Team{}).Error
}
