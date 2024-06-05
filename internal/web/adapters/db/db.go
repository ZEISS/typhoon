package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

const (
	accountUpdateFormat = "$SYS.ACCOUNT.%s.CLAIMS.UPDATE"
)

var _ ports.Repository = (*database)(nil)

type database struct {
	conn *gorm.DB
	nc   *nats.Conn
}

// NewDB returns a new instance of db.
func NewDB(conn *gorm.DB) *database {
	return &database{
		conn: conn,
	}
}

// NewDatastore returns a new instance of db.
func NewDatastore(conn *gorm.DB, nc *nats.Conn) (ports.Datastore, error) {
	return &database{
		conn: conn,
		nc:   nc,
	}, nil
}

// Close closes the database connection.
func (d *database) Close() error {
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// ReadWriteTx starts a read only transaction.
func (d *database) ReadWriteTx(ctx context.Context, fn func(context.Context, ports.ReadWriteTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx, d.nc}); err != nil {
		tx.Rollback()
	}

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ReadTx starts a read only transaction.
func (d *database) ReadTx(ctx context.Context, fn func(context.Context, ports.ReadTx) error) error {
	tx := d.conn.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, &datastoreTx{tx, d.nc}); err != nil {
		tx.Rollback()
	}

	if err := tx.Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

var _ ports.ReadTx = (*datastoreTx)(nil)
var _ ports.ReadWriteTx = (*datastoreTx)(nil)

type datastoreTx struct {
	tx *gorm.DB
	nc *nats.Conn
}

// GetOperator is a method that returns an operator by ID.
func (t *datastoreTx) GetOperator(ctx context.Context, operator *models.Operator) error {
	return t.tx.Preload("Accounts").Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Preload("Token").First(operator).Error
}

// CreateAccount is creating a new account.
func (t *datastoreTx) CreateAccount(ctx context.Context, account *models.Account) error {
	if err := t.tx.Preload("Accounts").Preload("SigningKeyGroups").Preload("SigningKeyGroups.Key").Preload("Key").Preload("Token").Create(account).Error; err != nil {
		return err
	}

	return t.nc.Publish(fmt.Sprintf(accountUpdateFormat, account.Token.ID), account.Token.Bytes())
}

// Conn returns the database connection.
func (d *database) Conn() *gorm.DB {
	return d.conn
}

// RunMigrations runs the database migrations.
func (d *database) RunMigrations() error {
	return d.conn.AutoMigrate(
		&adapters.GothUser{},
		&adapters.GothAccount{},
		&adapters.GothSession{},
		&adapters.GothVerificationToken{},
		&models.User{},
		&models.Account{},
		&models.Operator{},
		&models.System{},
		&models.Tag{},
		&models.Cluster{},
		&models.Token{},
		&models.SigningKeyGroup{},
	)
}
