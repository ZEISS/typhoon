package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account ...
type Account struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `json:"name"`

	// Operator is the operator that created the account.
	Operator   Operator  `json:"operator"`
	OperatorID uuid.UUID `json:"operator_id" gorm:"foreignKey:ID"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`

	// Accounts is the list of accounts that the operator has.
	SigningKeys []NKey `json:"signing_keys" gorm:"many2many:account_signing_keys;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
