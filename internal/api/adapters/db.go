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
		&models.User{},
		&models.System{},
		&models.Token{},
	)
}

// GetOperator ...
func (db *DB) GetOperator(ctx context.Context, id uuid.UUID) (*models.Operator, error) {
	operator := &models.Operator{}
	if err := db.conn.Where("id = ?", id).Preload("SigningKeys").Preload("Key").First(operator).Error; err != nil {
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

// CreateOperatorAccountSigningKey ...
func (db *DB) CreateOperatorAccountSigningKey(ctx context.Context, accountID uuid.UUID, key *models.NKey) error {
	return db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.Where("id = ?", accountID).First(&account).Error; err != nil {
			return err
		}

		err := tx.Model(&account).Association("SigningKeys").Append(key)
		if err != nil {
			return err
		}

		return nil
	})
}

// ListOperatorAccountsSigningKey ...
func (db *DB) ListOperatorAccountsSigningKey(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error) {
	accounts := []*models.Account{}

	err := db.conn.WithContext(ctx).Scopes(models.Paginate(&accounts, &pagination, db.conn)).Limit(pagination.Limit).Offset(pagination.Offset).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = accounts

	return &pagination, nil
}

// CreateOperatorAccount ...
func (db *DB) CreateOperatorAccount(ctx context.Context, account *models.Account) error {
	return db.conn.WithContext(ctx).Create(account).Error
}

// CreateOperatorAccountToken ...
func (db *DB) CreateOperatorAccountToken(ctx context.Context, accountID uuid.UUID, token *models.Token) error {
	return db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.Where("id = ?", accountID).First(&account).Error; err != nil {
			return err
		}

		err := tx.Model(&account).Association("Token").Replace(token)
		if err != nil {
			return err
		}

		return nil
	})
}

// GetOperatorAccount ...
func (db *DB) GetOperatorAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	account := &models.Account{}
	if err := db.conn.Where("id = ?", id).Preload("SigningKeys").First(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

// CreateOperatorToken ...
func (db *DB) CreateOperatorToken(ctx context.Context, operatorID uuid.UUID, token *models.Token) error {
	return db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var operator models.Operator
		if err := tx.Where("id = ?", operatorID).First(&operator).Error; err != nil {
			return err
		}

		err := tx.Model(&operator).Association("Token").Replace(token)
		if err != nil {
			return err
		}

		return nil
	})
}

// CreateOperatorAccountUser ...
func (db *DB) CreateOperatorAccountUser(ctx context.Context, user *models.User) error {
	return db.conn.WithContext(ctx).Create(user).Error
}

// GetOperatorAccountUser ...
func (db *DB) GetOperatorAccountUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := db.conn.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ListOperatorAccounts ...
func (db *DB) ListOperatorAccounts(ctx context.Context, operatorID uuid.UUID, pagination models.Pagination[*models.Account]) (*models.Pagination[*models.Account], error) {
	accounts := []*models.Account{}

	err := db.conn.WithContext(ctx).Scopes(models.Paginate(&accounts, &pagination, db.conn)).Limit(pagination.Limit).Offset(pagination.Offset).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = accounts

	return &pagination, nil
}

// CreateOperatorAccountUserToken ...
func (db *DB) CreateOperatorAccountUserToken(ctx context.Context, userID uuid.UUID, token *models.Token) error {
	return db.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
			return err
		}

		err := tx.Model(&user).Association("Token").Replace(token)
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
