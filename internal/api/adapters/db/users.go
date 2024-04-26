package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
)

// GetUser ...
func (db *DB) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := db.conn.Where("id = ?", id).Preload("Key").Preload("Account").Preload("Account.SigningKeys").First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser ...
func (db *DB) CreateUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Create(user).Error
}

// UpdateUser ...
func (db *DB) UpdateUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Save(user).Error
}

// ListUsers ...
func (db *DB) ListUsers(ctx context.Context, pagination models.Pagination[*models.User]) (*models.Pagination[*models.User], error) {
	users := []*models.User{}
	if err := db.conn.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return &models.Pagination[*models.User]{Rows: users}, nil
}

// DeleteUser ...
func (db *DB) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
}
