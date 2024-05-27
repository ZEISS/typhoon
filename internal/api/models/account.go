package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AccountPagination is the pagination for operators.
type AccountPagination Pagination[Operator]

// Account ...
type Account struct {
	// ID is the unique identifier for the account.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the account.
	Name string `json:"name"`
	// Description is the description of the account.
	Description *string `json:"description"`
	// Key is the issuer key identifier.
	Key NKey `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	// Token is the JWT token used to authenticate the account.
	Token Token `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	// Operator is the operator this account is associated with.
	Operator Operator `json:"operator" gorm:"foreignKey:OperatorID"`
	// OperatorID is the operator ID.
	OperatorID uuid.UUID `json:"operator_id" gorm:"foreignKey:ID"`
	// SigningKeyGroups is the list of signing key groups the account has.
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:account_signing_key_groups;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyGroupID"`
	// Users is the list of users the account has.
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// CreatedAt is the time the account was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the account was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the account was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// AfterDelete ...
func (a *Account) AfterDelete(tx *gorm.DB) error {
	return tx.Where("account_id = ?", a.ID).Delete(&User{}).Error
}

// FindSigningKeyGroupByID ...
func (a *Account) FindSigningKeyGroupByID(id uuid.UUID) *SigningKeyGroup {
	for _, skg := range a.SigningKeyGroups {
		if skg.ID == id {
			return &skg
		}
	}

	return nil
}
