package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OwnableType is a polymorphic type for ownership.
type OwnableType string

// OwnableType are the different types of ownable resources.
const (
	// SystemOwnable is a system.
	SystemOwnable OwnableType = "system"
	// AccountOwnable is an account.
	AccountOwnable OwnableType = "account"
	// OperatorOwnable is an operator.
	OperatorOwnable OwnableType = "operator"
	// UserOwnable is a user.
	UserOwnable OwnableType = "user"
)

// Ownership ...
type Ownership struct {
	// ID is the unique identifier for the ownership.
	ID int `json:"id" gorm:"primary_key"`
	// OwnableID is the unique identifier for .
	OwnableID uuid.UUID `json:"owner_id"`
	// OwnableType is the type of the owner.
	OwnableType string `json:"owner_type"`
	// TeamID is the identifier of the team.
	TeamID uuid.UUID `json:"team_id"`
	// Team is the team that this is owned by.
	Team Team `json:"team" gorm:"foreignKey:TeamID"`
	// CreatedAt is the time the ownership was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the ownership was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the ownership was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
