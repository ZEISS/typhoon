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
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// Name is the name of the system.
	Name string `json:"name" gorm:"unique" validate:"required,min=3,max=128"`
	// Description is the description of the system.
	Description string `json:"description" validate:"max=1024"`

	// Clusters is the clusters that are associated with the system.
	Clusters []Cluster `json:"clusters" gorm:"foreignKey:SystemID"`

	// Operator is the operator this is associated with this system to operate.
	Operator   *Operator `json:"operator" gorm:"foreignKey:OperatorID"`
	OperatorID uuid.UUID `json:"operator_id"`

	// SystemAccount is the account that is used to control the system.
	// The system account needs to be signed by the operator.
	SystemAccount   Account   `json:"system_account" gorm:"foreignKey:SystemAccountID"`
	SystemAccountID uuid.UUID `json:"system_account_id"`

	// Tags is the tags that are associated with the system.
	Tags []*Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:system;"`

	// OwnedBy is the owner of the account. This is usually a team.
	OwnedBy Ownership `json:"owned_by" gorm:"polymorphic:Ownable;polymorphicValue:system;"`

	// AllowedBy is the allowed by of the account. This is usually a team.
	AllowedBy []Allow `json:"allowed_by" gorm:"polymorphic:Allowable;polymorphicValue:system;"`

	// CreatedAt is the time the system was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the system was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the system was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
