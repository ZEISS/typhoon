package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team is a model for storing the team.
type Team struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name" validate:"required,alphanum,min=3,max=128"`
	Description string         `json:"description" validate:"max=1024"`
	ID          uuid.UUID      `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
}

//
