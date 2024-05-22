package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// ListOperators ...
func (db *database) ListOperators(ctx context.Context, pagination *models.Pagination[models.Operator]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Find(&pagination.Rows).Error
}
