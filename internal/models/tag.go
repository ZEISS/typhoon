package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tag ...
type Tag struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name" gorm:"uniqueIndex:idx_name_value"`
	Value     string         `json:"value" gorm:"uniqueIndex:idx_name_value"`
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}
