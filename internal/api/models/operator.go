package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Operator ...
type Operator struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `json:"name"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`

	// Accounts is the list of accounts that the operator has.
	SigningKeys []NKey `json:"signing_keys" gorm:"many2many:operator_signing_keys;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyID"`

	// Sysem is the system that the operator belongs to.
	System   System    `json:"system" gorm:"foreignKey:SystemID"`
	SystemID uuid.UUID `json:"system_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
