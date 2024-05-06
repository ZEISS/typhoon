package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// GetSystem ...
func (db *DB) GetSystem(ctx context.Context, id uuid.UUID) (*models.System, error) {
	system := &models.System{}
	err := db.conn.Where("id = ?", id).
		Preload("Clusters").
		Preload("Operator").
		First(system).Error
	if err != nil {
		return nil, err
	}

	return system, nil
}

// CreateSystem ...
func (db *DB) CreateSystem(ctx context.Context, system *models.System) error {
	return db.conn.WithContext(ctx).Create(system).Error
}

// ListSystems ...
func (db *DB) ListSystems(ctx context.Context, pagination models.Pagination[models.System]) (models.Pagination[models.System], error) {
	systems := []models.System{}

	err := db.conn.WithContext(ctx).
		Scopes(models.Paginate(&systems, &pagination, db.conn)).
		Preload("Clusters").
		Preload("Operator").
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Find(&systems).Error
	if err != nil {
		return pagination, err
	}
	pagination.Rows = systems

	return pagination, nil
}

// DeleteSystem ...
func (db *DB) DeleteSystem(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Delete(&models.System{}, id).Error
}

// UpdateSystem ...
func (db *DB) UpdateSystem(ctx context.Context, system *models.System) (*models.System, error) {
	if err := db.conn.WithContext(ctx).Save(system).Error; err != nil {
		return nil, err
	}

	return system, nil
}
