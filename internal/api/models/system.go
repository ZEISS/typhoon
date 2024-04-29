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

	// Operator is the operator that the system belongs to.
	Operator   Operator   `json:"operator" gorm:"foreignKey:OperatorID"`
	OperatorID *uuid.UUID `json:"operator_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
