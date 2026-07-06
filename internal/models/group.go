package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SigningKeyGroup is a model for storing the signing key group.
type SigningKeyGroup struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Key         NKey           `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:signing_key_group"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name" validate:"required,min=3,max=128"`
	Description string         `json:"description" validate:"max=1024"`
	KeyID       string         `json:"key_id" gorm:"foreignKey:ID"`
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}
