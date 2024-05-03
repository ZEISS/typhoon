package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account ...
type Account struct {
	// ID is the unique identifier for the account.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the account.
	Name string `json:"name"`

	// Operator is the operator that created the account.
	Operator   Operator  `json:"operator"`
	OperatorID uuid.UUID `json:"operator_id" gorm:"foreignKey:ID"`

	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`

	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`

	// SigningKeys is the list of signing keys the account has.
	SigningKeys []NKey `json:"signing_keys" gorm:"many2many:account_signing_keys;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyID"`

	// Users is the list of users that the account has.
	Users []User `json:"users" gorm:"foreignKey:AccountID"`

	// OwnedBy is the owner of the account. This is usually a team.
	OwnedBy Ownership `json:"owner" gorm:"polymorphic:Ownable;polymorphicValue:account;"`

	// CreatedAt is the time the account was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the account was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the account was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
