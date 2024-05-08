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
	// Key is the issuer key identifier.
	Key   NKey   `json:"key"`
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`
	// Token is the JWT token used to authenticate the account.
	Token   Token  `json:"token" gorm:"foreignKey:TokenID"`
	TokenID string `json:"token_id"`
	// SigningKeyGroups is the list of signing key groups the account has.
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:account_signing_key_groups;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyGroupID"`
	// SignedBy is the entity that signs this one.
	SignedBy Signature `json:"signed_by" gorm:"polymorphic:Signable;polymorphicValue:account;"`
	// CreatedAt is the time the account was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the account was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the account was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
