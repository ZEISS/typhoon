package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm"
)

// GetOperator is a method to get an operator from the database.
func (db *DB) GetOperator(ctx context.Context, op *models.Operator) error {
	err := db.conn.
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Token").
		Preload("Key").
		First(&op).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteOperator ...
func (db *DB) DeleteOperator(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.Operator{}).Error
}

// CreateOperator is a method to create an operator in the database.
func (db *DB) CreateOperator(ctx context.Context, op *models.Operator) error {
	return db.conn.WithContext(ctx).Create(op).Error
}

// UpdateOperator ...
func (db *DB) UpdateOperator(ctx context.Context, op *models.Operator) error {
	return db.conn.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Save(op).Error
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
