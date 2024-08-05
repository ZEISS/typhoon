package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
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
	Key NKey `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:user"`
	// KeyID is the foreign key for the key.
	KeyID string `json:"key_id" gorm:"foreignKey:ID"`
	// Token is the JWT token used to authenticate the account.
	Token Token `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:user"`
	// UserLimits is the user limits.
	UserLimits UserLimits `json:"limits"`
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

// UserLimits ...
type UserLimits struct {
	// UserID is the user identifier.
	UserID uuid.UUID `json:"user_id"`
	// MaxSubscriptions is the maximum number of subscriptions the user can have.
	MaxSubscriptions int `json:"max_subscriptions"`
	// MaxJWTLiftime is the maximum payload size the user can have.
	MaxJWTLiftime time.Duration `json:"jwt_lifetime"`
	// MaxPayloadSize is the maximum payload size the user can have.
	MaxPayloadSize int `json:"max_payload_size"`
	// MaxDataRate is the maximum data rate the user can have.
	MaxDataRate int `json:"max_data_rate"`
}

// Credentials returns the user's credentials.
func (u *User) Credentials() ([]byte, error) {
	creds, err := jwt.FormatUserConfig(u.Token.Token, u.Key.Seed)
	if err != nil {
		return nil, err
	}

	return creds, nil
}
