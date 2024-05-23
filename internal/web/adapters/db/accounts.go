package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
)

// ListAccounts ...
func (db *database) ListAccounts(ctx context.Context, pagination *models.Pagination[models.Account]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Find(&pagination.Rows).Error
}

// CreateAccount ...
func (db *database) CreateAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Create(account).Error
}

// GetAccount ...
func (db *database) GetAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").First(account).Error
}

// UpdateAccount ...
func (db *database) UpdateAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Save(account).Error
}

// DeleteAccount ...
func (db *database) DeleteAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Delete(account).Error
}
