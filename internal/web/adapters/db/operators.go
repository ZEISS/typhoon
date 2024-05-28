package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm/clause"
)

// ListOperators ...
func (db *database) ListOperators(ctx context.Context, pagination *models.Pagination[models.Operator]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Find(&pagination.Rows).Error
}

// CreateOperator ...
func (db *database) CreateOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Create(operator).Error
}

// GetOperator ...
func (db *database) GetOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Preload("Accounts").Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Preload("Token").First(operator).Error
}

// UpdateOperator ...
func (db *database) UpdateOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Save(operator).Error
}

// DeleteOperator ...
func (db *database) DeleteOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.Debug().WithContext(ctx).Select(clause.Associations).Delete(operator).Debug().Error
}
