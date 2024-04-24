package adapters

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"

	"gorm.io/gorm"
)

var (
	_ ports.Teams     = (*DB)(nil)
	_ ports.Systems   = (*DB)(nil)
	_ ports.Operators = (*DB)(nil)
)

// DB ...
type DB struct {
	conn *gorm.DB
}

// NewDB ...
func NewDB(conn *gorm.DB) *DB {
	return &DB{conn}
}

// RunMigrations ...
func (db *DB) RunMigrations() error {
	return db.conn.AutoMigrate(
		&models.NKey{},
		&models.Operator{},
		&models.Account{},
		&models.System{},
		&models.Token{},
	)
}

// GetOperator ...
func (db *DB) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	operator := &models.Operator{}
	if err := db.conn.Where("id = ?", id).First(operator).Error; err != nil {
		return nil, err
	}

	return operator, nil
}

// DeleteOperator ...
func (db *DB) DeleteOperator(ctx context.Context, id uuid.UUID) error {
	return db.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.Operator{}).Error
}

// CreateAccount ...
func (db *DB) CreateAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Create(account).Error
}

// CreateOperatorSigningKey ...
func (db *DB) CreateOperatorSigningKey(ctx context.Context, operatorID uuid.UUID, key *models.NKey) error {
	return db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var operator models.Operator
		if err := tx.Where("id = ?", operatorID).First(&operator).Error; err != nil {
			return err
		}

		err := tx.Model(&operator).Association("SigningKeys").Append(key)
		if err != nil {
			return err
		}

		return nil
	})
}

// ListOperator ...
func (db *DB) ListOperator(ctx context.Context, pagination models.Pagination[*models.Operator]) (*models.Pagination[*models.Operator], error) {
	operators := []*models.Operator{}

	err := db.conn.WithContext(ctx).Scopes(models.Paginate(&operators, &pagination, db.conn)).Limit(pagination.Limit).Offset(pagination.Offset).Find(&operators).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = operators

	return &pagination, nil
}

// CreateOperator ...
func (db *DB) CreateOperator(ctx context.Context, operator *models.Operator) error {
	return db.conn.WithContext(ctx).Create(operator).Error
}
