package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt"
	"gorm.io/gorm"
)

// User ...
type User struct {
	// ID is the unique identifier for the user.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the user.
	Name string `json:"name" validate:"required,min=3,max=128"`
	// Description is the description of the user.
	Description string `json:"description" validate:"max=1024"`
	// Key is the issuer key identifier.
	Key NKey `json:"key"`
	// KeyID is the foreign key for the key.
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`
	// Token is the JWT token used to authenticate the account.
	Token Token `json:"token" gorm:"foreignKey:TokenID"`
	// TokenID is the foreign key for the token.
	TokenID string `json:"token_id"`
	// Account is the account that the user belongs to.
	Account Account `json:"account" gorm:"foreignKey:AccountID"`
	// AccountID is the foreign key for the account.
	AccountID uuid.UUID `json:"account_id"`
	// CreatedAt is the time the user was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the user was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the user was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Credentials returns the user's credentials.
func (u *User) Credentials() ([]byte, error) {
	creds, err := jwt.FormatUserConfig(u.Token.Token, u.Key.Seed)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
