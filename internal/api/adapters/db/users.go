package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
)

// GetUser ...
func (db *DB) GetUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Preload("Key").Preload("Token").Preload("Account").Preload("Account.SigningKeyGroups").First(user).Error
}

// CreateUser ...
func (db *DB) CreateUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Create(user).Error
}

// UpdateUser ...
func (db *DB) UpdateUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Save(user).Error
}

// ListAccounts ...
func (db *DB) ListUsers(ctx context.Context, pagination models.Pagination[models.User]) (models.Pagination[models.User], error) {
	if err := db.conn.WithContext(ctx).Find(&pagination.Rows).Error; err != nil {
		return pagination, err
	}

	return pagination, nil
}

// DeleteUser ...
func (db *DB) DeleteUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Delete(user).Error
}
