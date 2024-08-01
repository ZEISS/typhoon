package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/models"
)

// GetSystem is returning a system by its ID.
func (db *DB) GetSystem(ctx context.Context, system *models.System) error {
	return db.conn.WithContext(ctx).Preload("Clusters").Preload("Operator").First(&system).Error
}

// CreateSystem ...
func (db *DB) CreateSystem(ctx context.Context, system *models.System) error {
	return db.conn.WithContext(ctx).Create(system).Error
}

// ListSystems ...
func (db *DB) ListSystems(ctx context.Context, pagination *models.Pagination[models.System]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Find(&pagination.Rows).Error
}

// DeleteSystem ...
func (db *DB) DeleteSystem(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Delete(&models.System{}, id).Error
}

// UpdateSystem ...
func (db *DB) UpdateSystem(ctx context.Context, system *models.System) error {
	return db.conn.WithContext(ctx).Save(system).Error
}
