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

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
