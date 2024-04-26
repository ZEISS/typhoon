package models

import (
	"time"

	"gorm.io/gorm"
)

// Token ...
type Token struct {
	ID    string `json:"token_id" gorm:"primaryKey"`
	Token string `json:"token"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
