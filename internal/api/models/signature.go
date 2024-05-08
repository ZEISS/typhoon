package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SignerType is a polymorphic type for ownership.
type SignableType string

// SignableType are the different types that are signed.
const (
	// OperatorSignable is an operator.
	OperatorSignable SignableType = "operator"
	// AccountSignable is a system.
	AccountSignable SignableType = "account"
	// UserSignable is a user.
	UserSignable SignableType = "user"
)

// Signature ...
type Signature struct {
	// ID is the unique identifier for the ownership.
	ID int `json:"id" gorm:"primary_key"`
	// SignableID is the unique identifier for the signable resource.
	SignableID uuid.UUID `json:"signable_id"`
	// SignableType is the type of the signable resource.
	SignableType string `json:"signable_type"`
	// SignerID is the identifier of the signer.
	SignerID uuid.UUID `json:"signer_id"`
	// Signer is the signer.
	Signer SigningKeyGroup `json:"signer" gorm:"foreignKey:SignerID"`
	// CreatedAt is the time the ownership was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the ownership was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the ownership was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
