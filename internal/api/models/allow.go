package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AllowableType is a polymorphic type for allow.
type AllowableType string

// OwnableType are the different types of ownable resources.
const (
	// TeamAllowable is a team.
	TeamAllowable AllowableType = "team"
	// UserAllowable is a user.
	UserAllowable AllowableType = "user"
)

// Allow ...
type Allow struct {
	// ID is the unique identifier for the ownership.
	ID int `json:"id" gorm:"primary_key"`
	// AllowableID is the unique identifier for the resource allowed to.
	AllowableID uuid.UUID `json:"owner_id"`
	// AllowableType is the type of the resource that is allowed to.
	AllowableType string `json:"owner_type"`
	// TeamID is the .
	TeamID uuid.UUID `json:"team_id"`
	// Team is the team that is allowed to.
	Team Team `json:"team" gorm:"foreignKey:TeamID"`
	// CreatedAt is the time the ownership was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the ownership was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the ownership was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
