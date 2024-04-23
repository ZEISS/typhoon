package models

import (
	"time"

	"gorm.io/gorm"
)

// Operator ...
type Operator struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
