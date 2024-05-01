package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// System ...
type System struct {
	// ID is the unique identifier for the system.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Name is the name of the system.
	Name string `json:"name" gorm:"unique" validate:"required,min=3,max=128"`
	// Description is the description of the system.
	Description string `json:"description" validate:"max=1024"`

	// Operator is the operator this is associated with this system to operate.
	Operator   Operator   `json:"operator" gorm:"foreignKey:OperatorID"`
	OperatorID *uuid.UUID `json:"operator_id"`

	// SystemAccount is the account used to control the system.
	// The system account needs to be signed by the operator.
	SystemAccount   Account    `json:"system_account" gorm:"foreignKey:SystemAccountID"`
	SystemAccountID *uuid.UUID `json:"system_account_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
