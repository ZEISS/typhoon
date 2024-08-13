package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FromAPI ...
type FromAPI[T any] interface {
	FromAPI(api *T)
}

// ToAPI ...
type ToAPI[T any] interface {
	ToAPI() *T
}

// System ...
type System struct {
	// ID is the unique identifier for the system.
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()" xml:"id" params:"id"`
	// Name is the name of the system.
	Name string `json:"name" gorm:"unique" xml:"name" form:"name" validate:"required,min=3,max=128"`
	// Description is the description of the system.
	Description string `json:"description" form:"description" validate:"max=1024"`
	// Clusters is the clusters that are associated with the system.
	Clusters []Cluster `json:"clusters" gorm:"foreignKey:SystemID"`
	// Operator is the operator this is associated with this system to operate.
	Operator Operator `json:"operator" gorm:"foreignKey:OperatorID" validate:"-"`
	// OperatorID is the operator ID that is associated with the system.
	OperatorID uuid.UUID `json:"operator_id" form:"operator_id" validate:"required,uuid"`
	// Tags are the tags associated with the environment
	Tags []Tag `json:"tags" gorm:"many2many:system_tags;"`
	// CreatedAt is the time the system was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the system was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the system was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
