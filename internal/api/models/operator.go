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

	// Tokens is the list of tokens that the operator has.
	Token   *Token `json:"tokens" gorm:"foreignKey:ID"`
	TokenID string `json:"token_id" gorm:"foreignKey:ID"`

	// Accounts is the list of accounts that the operator has.
	SigningKeys []NKey `json:"signing_keys" gorm:"many2many:operator_signing_keys;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
