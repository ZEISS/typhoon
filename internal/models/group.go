package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SigningKeyGroup is a model for storing the signing key group.
type SigningKeyGroup struct {
	// ID is the unique identifier for the group.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the group.
	Name string `json:"name" validate:"required,min=3,max=128"`
	// Description is the description of the group.
	Description string `json:"description" validate:"max=1024"`
	// Key is the signing key of this group.
	Key NKey `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:signing_key_group"`
	// KeyID is the foreign key for the key.
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`
	// CreatedAt is the time the group was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the group was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the group was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
