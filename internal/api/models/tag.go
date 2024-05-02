package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TaggableType ...
type TaggableType string

// TaggableType ...
const (
	SystemTaggable TaggableType = "system"
)

// Tag is the model for adding tags to resources.
type Tag struct {
	// ID is the unique identifier for the tag.
	ID int `json:"id" gorm:"primary_key"`
	// Name is the name of the tag.
	Name string `json:"name"`
	// TaggableID is the unique identifier for the taggable.
	TaggableID uuid.UUID `json:"taggable_id"`
	// TaggableType is the type of the taggable.
	TaggableType TaggableType `json:"taggable_type"`
	// CreatedAt is the time the tag was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the tag was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the tag was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
