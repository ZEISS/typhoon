package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"gorm.io/gorm"
)

// OperatorPagination is the pagination for operators.
type OperatorPagination Pagination[Operator]

// Operator is the operator that is used to manage the systems.
type Operator struct {
	// ID is the unique identifier for the operator.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the operator.
	Name string `json:"name" validate:"required,min=3,max=128"`
	// Description is the description of the operator.
	Description string `json:"description" validate:"max=1024"`
	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`
	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`
	// SystemAdminAccount is the account that is used to manage the systems.
	SystemAdminAccount   *Account   `json:"system_admin_account" gorm:"many2many:operator_system_admin_accounts;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SystemAdminAccountID"`
	SystemAdminAccountID *uuid.UUID `json:"system_admin_account_id"`
	// Systems is the systems that are associated with the operator.
	Systems []System `json:"systems" gorm:"foreignKey:OperatorID"`
	// SigningKeyGroups is the list of signing key groups the account has.
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:operator_signing_key_groups;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyGroupID"`
	// CreatedAt is the time the operator was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the operator was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the operator was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Claims returns the operator claims.
func (o Operator) Claims() (*jwt.OperatorClaims, error) {
	claims, err := jwt.DecodeOperatorClaims(o.Token.Token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
