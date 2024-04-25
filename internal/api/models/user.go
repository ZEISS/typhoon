package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User ...
type User struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `json:"name"`

	// Account is the account that created the user.
	Account   Account   `json:"account"`
	AccountID uuid.UUID `json:"account_id" gorm:"foreignKey:ID"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	// Token is the JWT token used to authenticate the account.
	Token   string `json:"token"`
	TokenID string `json:"token_id" gorm:"foreignKey:ID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
