package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"gorm.io/gorm"
)

// User ...
type User struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Account     Account        `json:"account" gorm:"foreignKey:AccountID"`
	Key         NKey           `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:user"`
	Token       Token          `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:user"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name" validate:"required,min=3,max=128"`
	Description string         `json:"description" validate:"max=1024"`
	KeyID       string         `json:"key_id" gorm:"foreignKey:ID"`
	UserLimits  UserLimits     `json:"limits"`
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AccountID   uuid.UUID      `json:"account_id"`
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
