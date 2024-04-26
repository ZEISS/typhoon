package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// GetOperator ...
func (db *DB) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	operator := &models.Operator{}
	if err := db.conn.Where("id = ?", id).Preload("SigningKeys").Preload("Key").First(operator).Error; err != nil {
		return nil, err
	}

	return operator, nil
}

// DeleteOperator ...
func (db *DB) DeleteOperator(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.Operator{}).Error
}

// CreateOperator ...
func (db *DB) CreateOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Create(operator).Error
}

// UpdateOperator ...
func (db *DB) UpdateOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Save(operator).Error
}

// ListOperators ...
func (db *DB) ListOperators(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error) {
	operators := []*models.Operator{}
	if err := db.conn.WithContext(ctx).Find(&operators).Error; err != nil {
		return nil, err
	}

	return &models.Pagination[*models.Operator]{Rows: operators}, nil
}
