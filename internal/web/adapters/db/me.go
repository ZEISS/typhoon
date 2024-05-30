package db

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
)

// GetProfile is a method that returns the profile of the current user
func (db *database) GetProfile(ctx context.Context, user *adapters.GothUser) error {
	return db.conn.WithContext(ctx).First(user).Error
}
