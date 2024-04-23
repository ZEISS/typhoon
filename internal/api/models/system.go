package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// System ...
type System struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Cluster ...
type Cluster struct {
	ID  uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	URL string    `json:"url"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
