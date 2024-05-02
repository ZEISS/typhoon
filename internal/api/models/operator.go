package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Operator is the operator that is used to manage the systems.
type Operator struct {
	// ID is the unique identifier for the operator.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the operator.
	Name string `json:"name" validate:"required,min=3,max=128"`
	// Description is the description of the operator.
	Description string `json:"description" validate:"max=1024"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`

	// Systems is the list of systems that the operator has.
	Systems []System `json:"systems" gorm:"many2many:operator_systems;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SystemID"`

	// Accounts is the list of accounts that the operator has.
	SigningKeys []NKey `json:"signing_keys" gorm:"many2many:operator_signing_keys;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyID"`

	// Owner is the owner of the operator.
	Owner Ownership `json:"owner" gorm:"polymorphic:Ownable;polymorphicValue:operator;"`

	// CreatedAt is the time the operator was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the operator was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the operator was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
