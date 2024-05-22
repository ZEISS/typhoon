package db

import (
	"context"

	"github.com/zeiss/fiber-goth/adapters"
)

// GetUser ...
func (d *database) GetUser(ctx context.Context, user *adapters.GothUser) error {
	return d.conn.WithContext(ctx).First(user).Error
}
