package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// GetSystem is returning a system by its ID.
func (db *DB) GetSystem(ctx context.Context, system *models.System) error {
	err := db.conn.
		Preload("Clusters").
		Preload("Operator").
		First(&system).Error
	if err != nil {
		return err
	}

	return nil
}

// CreateSystem ...
func (db *DB) CreateSystem(ctx context.Context, system *models.System) error {
	return db.conn.WithContext(ctx).Create(system).Error
}

// ListSystems ...
func (db *DB) ListSystems(ctx context.Context, pagination *models.Pagination[models.System]) error {
	err := db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Find(&pagination.Rows).Error
	if err != nil {
		return err
	}

	return nil
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
