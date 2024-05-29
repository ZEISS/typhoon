package db

import (
	"context"

	"github.com/zeiss/typhoon/internal/api/models"
	"gorm.io/gorm/clause"
)

// GetUser ...
func (d *database) GetUser(ctx context.Context, user *models.User) error {
	return d.conn.WithContext(ctx).
		Preload("Key").
		Preload("Token").
		Preload("Account").
		Preload("Account.SigningKeyGroups").
		First(user).Error
}

// ListUsers ...
func (d *database) ListUsers(ctx context.Context, pagination *models.Pagination[models.User]) error {
	return d.conn.WithContext(ctx).Scopes(models.Paginate(&pagination.Rows, pagination, d.conn)).Find(&pagination.Rows).Error
}

// CreateUser ...
func (d *database) CreateUser(ctx context.Context, user *models.User) error {
	return d.conn.WithContext(ctx).Create(user).Error
}

// DeleteUser ...
func (d *database) DeleteUser(ctx context.Context, user *models.User) error {
	return d.conn.Select(clause.Associations).WithContext(ctx).Delete(user).Error
}

// UpdateUser ...
func (d *database) UpdateUser(ctx context.Context, user *models.User) error {
	return d.conn.WithContext(ctx).Updates(user).Error
}
