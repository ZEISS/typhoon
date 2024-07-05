package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeiss/fiber-goth/adapters"
	"gorm.io/gorm"
)

// Subscription is the subscription that is used to manage the systems.
type Subscription struct {
	// ID is the unique identifier for the account.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the account.
	Name string `json:"name"`
	// Description is the description of the account.
	Description string `json:"description"`
	// System is the system that is associated with the subscription.
	System System `json:"system" gorm:"foreignKey:SystemID"`
	// SystemID is the unique identifier for the system.
	SystemID uuid.UUID `json:"system_id"`
	// Team is the team that is associated with the subscription.
	Team adapters.GothTeam `json:"team" gorm:"foreignKey:TeamID"`
	// TeamID is the unique identifier for the team.
	TeamID uuid.UUID `json:"team_id"`
	// CreatedAt is the time the system was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the system was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the system was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
