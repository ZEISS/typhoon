package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"

	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	accountUpdateFormat = "$SYS.ACCOUNT.%s.CLAIMS.UPDATE"
)

type database struct {
	conn *gorm.DB
	nc   *nats.Conn
}

// NewDatastore returns a new instance of db.
func NewDB(conn *gorm.DB, nc *nats.Conn) (ports.Datastore, error) {
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

// RunMigrations runs the database migrations.
func (d *database) Migrate(ctx context.Context) error {
	return d.conn.WithContext(ctx).AutoMigrate(
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
		&models.UserLimits{},
	)
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
	return t.tx.Preload("Accounts").
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Key").
		Preload("Token").
		First(operator).Error
}

// CreateAccount is creating a new account.
func (t *datastoreTx) CreateAccount(ctx context.Context, account *models.Account) error {
	if err := t.tx.Preload("Accounts").
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Key").
		Preload("Token").
		Create(account).Error; err != nil {
		return err
	}

	return t.nc.Publish(fmt.Sprintf(accountUpdateFormat, account.Token.ID), account.Token.Bytes())
}

// ListAccounts ...
func (t *datastoreTx) ListAccounts(ctx context.Context, pagination *models.Pagination[models.Account]) error {
	return t.tx.Scopes(models.Paginate(&pagination.Rows, pagination, t.tx)).
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Key").
		Find(&pagination.Rows).Error
}

// GetAccount ...
func (t *datastoreTx) GetAccount(ctx context.Context, account *models.Account) error {
	return t.tx.Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Key").
		Preload("Token").
		Preload("Operator").
		Preload("Operator.Key").
		First(account).Error
}

// UpdateAccount ...
func (t *datastoreTx) UpdateAccount(ctx context.Context, account *models.Account) error {
	return t.tx.Save(account).Error
}

// DeleteAccount ...
func (t *datastoreTx) DeleteAccount(ctx context.Context, account *models.Account) error {
	return t.tx.Select(clause.Associations).Delete(account).Error
}

// ListOperators ...
func (t *datastoreTx) ListOperators(ctx context.Context, pagination *tables.Results[models.Operator]) error {
	return t.tx.Scopes(tables.PaginatedResults(&pagination.Rows, pagination, t.tx)).
		Preload("SigningKeyGroups").
		Preload("SigningKeyGroups.Key").
		Preload("Key").
		Find(&pagination.Rows).Error
}

// CreateOperator ...
func (t *datastoreTx) CreateOperator(ctx context.Context, operator *models.Operator) error {
	return t.tx.Create(operator).Error
}

// UpdateOperator ...
func (t *datastoreTx) UpdateOperator(ctx context.Context, operator *models.Operator) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(operator).Error
}

// DeleteOperator ...
func (t *datastoreTx) DeleteOperator(ctx context.Context, operator *models.Operator) error {
	return t.tx.Debug().Select(clause.Associations).Delete(operator).Debug().Error
}

// GetUser ...
func (t *datastoreTx) GetUser(ctx context.Context, user *models.User) error {
	return t.tx.Preload("Key").
		Preload("Token").
		Preload("Account").
		Preload("Account.SigningKeyGroups").
		First(user).Error
}

// ListUsers ...
func (t *datastoreTx) ListUsers(ctx context.Context, pagination *models.Pagination[models.User]) error {
	return t.tx.Scopes(models.Paginate(&pagination.Rows, pagination, t.tx)).Find(&pagination.Rows).Error
}

// CreateUser ...
func (t *datastoreTx) CreateUser(ctx context.Context, user *models.User) error {
	return t.tx.Create(user).Error
}

// DeleteUser ...
func (t *datastoreTx) DeleteUser(ctx context.Context, user *models.User) error {
	return t.tx.Select(clause.Associations).Delete(user).Error
}

// UpdateUser ...
func (t *datastoreTx) UpdateUser(ctx context.Context, user *models.User) error {
	return t.tx.Updates(user).Error
}

// GetProfile is a method that returns the profile of the current user
func (t *datastoreTx) GetProfile(ctx context.Context, user *adapters.GothUser) error {
	return t.tx.First(user).Error
}

// CreateSystem is a method that creates a new system
func (t *datastoreTx) CreateSystem(ctx context.Context, system *models.System) error {
	return t.tx.Create(system).Error
}

// GetSystem is a method that returns a system by ID
func (t *datastoreTx) GetSystem(ctx context.Context, system *models.System) error {
	return t.tx.Preload("Tags").First(system).Error
}

// ListSystems is a method that returns a list of systems
func (t *datastoreTx) ListSystems(ctx context.Context, pagination *models.Pagination[models.System]) error {
	return t.tx.Scopes(models.Paginate(&pagination.Rows, pagination, t.tx)).Preload("Tags").Find(&pagination.Rows).Error
}

// DeleteSystem is a method that deletes a system
func (t *datastoreTx) DeleteSystem(ctx context.Context, system *models.System) error {
	return t.tx.Select(clause.Associations).Delete(system).Error
}

// UpdateSystem is a method that updates a system
func (t *datastoreTx) UpdateSystem(ctx context.Context, system *models.System) error {
	return t.tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(system).Error
}
