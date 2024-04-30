package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm"
)

// GetAccount ...
func (db *DB) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	account := &models.Account{}
	err := db.conn.
		Where("id = ?", id).
		Preload("Token").
		Preload("Operator").
		Preload("Operator.Token").
		Preload("Operator.SigningKeys").
		Preload("SigningKeys").
		Preload("Users").
		First(account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
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
func (db *DB) ListAccounts(ctx context.Context, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error) {
	accounts := []*models.Account{}
	if err := db.conn.WithContext(ctx).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return &models.Pagination[*models.Account]{Rows: accounts}, nil
}
