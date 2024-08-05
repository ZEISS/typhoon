package database

import (
	"context"

	"github.com/zeiss/pkg/dbx"
	"github.com/zeiss/typhoon/internal/accounts/ports"
	"github.com/zeiss/typhoon/internal/models"

	"gorm.io/gorm"
)

var _ ports.ReadTx = (*readTxImpl)(nil)

type readTxImpl struct {
	conn *gorm.DB
}

// GetToken is returning a token.
func (r *readTxImpl) GetToken(ctx context.Context, token *models.Token) error {
	return r.conn.WithContext(ctx).Where(token).First(token).Error
}

// NewReadTx ...
func NewReadTx() dbx.ReadTxFactory[ports.ReadTx] {
	return func(db *gorm.DB) (ports.ReadTx, error) {
		return &readTxImpl{conn: db}, nil
	}
}

type writeTxImpl struct {
	conn *gorm.DB
	readTxImpl
}

// NewWriteTx ...
func NewWriteTx() dbx.ReadWriteTxFactory[ports.ReadWriteTx] {
	return func(db *gorm.DB) (ports.ReadWriteTx, error) {
		return &writeTxImpl{conn: db}, nil
	}
}
