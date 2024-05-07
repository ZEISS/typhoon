package models

import (
	"time"

	"gorm.io/gorm"
)

// Token is a model for storing the the JWT token used to authenticate the user.
type Token struct {
	// ID is the unique identifier for the token.
	// This is the public key portion of the NKey.
	ID string `json:"token_id" gorm:"primaryKey"`
	// Token is the JWT token used to authenticate the account.
	Token string `json:"token"`
	// CreatedAt is the time the token was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the token was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the token was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
