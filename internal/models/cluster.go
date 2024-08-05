package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Cluster ...
type Cluster struct {
	// ID is the unique identifier for the cluster.
	ID uuid.UUID `json:"id" gorm:"primaryKey,type:uuid;default:gen_random_uuid()"`
	// Name is the name of the cluster.
	Name string `json:"name" gorm:"unique" validate:"required,min=3,max=128"`
	// Description is the description of the cluster.
	Description string `json:"description" validate:"max=1024"`
	// ServerURL is the URL of the server.
	ServerURL string `json:"url" validate:"required"`
	// SystemID is the ID of the system the cluster belongs to.
	SystemID uuid.UUID `json:"system_id"`
	// CreatedAt is the time the cluster was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the cluster was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the cluster was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
