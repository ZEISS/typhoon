package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team is a model for storing the team.
type Team struct {
	// ID is the unique identifier of the team.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// Name is the name of the team.
	Name string `json:"name" validate:"required,alphanum,min=3,max=128"`
	// Description is the description of the team.
	Description string `json:"description" validate:"max=255"`
	// CreatedAt is the creation time of the team.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the team.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the team.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
