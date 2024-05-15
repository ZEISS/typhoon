package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm"
)

// GetAccount ...
func (db *DB) GetAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Preload("Operator").Preload("Token").Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").First(account).Error
}

// CreateAccount ...
func (db *DB) CreateAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Create(account).Error
}

// UpdateAccount ...
func (db *DB) UpdateAccount(ctx context.Context, account *models.Account) error {
	return db.conn.Session(&gorm.Session{FullSaveAssociations: true}).WithContext(ctx).Updates(account).Error
}

// ListAccounts ...
func (db *DB) ListAccounts(ctx context.Context, pagination *models.Pagination[models.Account]) error {
	return db.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, db.conn)).Preload("Key").Find(&pagination.Rows).Error
}

// ListSigningKeys ...
func (db *DB) ListSigningKeys(ctx context.Context, accountID uuid.UUID, pagination models.Pagination[models.NKey]) (models.Pagination[models.NKey], error) {
	keys := []models.NKey{}

	account := &models.Account{}
	err := db.conn.
		Where("id = ?", accountID).
		Preload("SigningKeys").
		First(account).Error
	if err != nil {
		return pagination, err
	}

	pagination.Rows = keys

	return pagination, nil
}
