package database

import (
	"context"

	"github.com/zeiss/typhoon/internal/accounts/ports"
	"github.com/zeiss/typhoon/internal/api/models"

	"gorm.io/gorm"
)

var _ ports.ReadWriteTx = (*datastoreImpl)(nil)

type datastoreImpl struct {
	conn *gorm.DB
}

// NewDatastore ...
func NewDatastore(conn *gorm.DB) *datastoreImpl {
	return &datastoreImpl{conn}
}

// GetToken us
func (d *datastoreImpl) GetToken(ctx context.Context, account *models.Token) error {
	return d.conn.WithContext(ctx).Find(account).Error
}
