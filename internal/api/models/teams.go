package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Team ...
type Team struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
