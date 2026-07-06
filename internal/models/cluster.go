package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Cluster ...
type Cluster struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name" gorm:"unique" validate:"required,min=3,max=128"`
	Description string         `json:"description" validate:"max=1024"`
	ServerURL   string         `json:"url" validate:"required"`
	ID          uuid.UUID      `json:"id" gorm:"primaryKey,type:uuid;default:gen_random_uuid()"`
	SystemID    uuid.UUID      `json:"system_id"`
}
