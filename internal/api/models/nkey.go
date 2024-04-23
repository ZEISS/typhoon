package models

import (
	"time"

	"gorm.io/gorm"
)

// NKey holds a private key and its metadata.
type NKey struct {
	// ID is the public key portion of the NKey.
	ID string `json:"id" gorm:"primaryKey"`
	// Seed is the private key portion of the NKey.
	Seed []byte `json:"seed"`
	// CreatedAt is the timestamp the key was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the timestamp the key was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the timestamp the key was deleted
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
