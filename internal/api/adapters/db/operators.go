package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm"
)

// GetOperator ...
func (db *DB) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	operator := &models.Operator{}

	err := db.conn.
		Where("id = ?", id).
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Token").
		Preload("Key").
		First(operator).Error
	if err != nil {
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
	return db.conn.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Save(operator).Error
}

// ListOperators ...
func (db *DB) ListOperators(ctx context.Context, pagination models.OperatorPagination) (models.OperatorPagination, error) {
	operators := []models.Operator{}

	err := db.conn.WithContext(ctx).Find(&operators).Error
	if err != nil {
		return pagination, err
	}
	pagination.Rows = operators

	return pagination, nil
}
