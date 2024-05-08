package models

import (
	"time"

	"github.com/nats-io/jwt"
	"gorm.io/gorm"
)

// Token is a model for storing the the JWT token used to authenticate the user.
type Token struct {
	// ID is the unique identifier for the token.
	// This is the public key portion of the NKey.
	ID string `json:"token_id" gorm:"primaryKey"`
	// Token is the JWT token used to authenticate the account.
	Token string `json:"token"`
	// CreatedAt is the time the token was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the token was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the token was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// DeepCopy returns a deep copy of the token.
func (t *Token) DeepCopy() Token {
	return Token{
		ID:        t.ID,
		Token:     t.Token,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		DeletedAt: t.DeletedAt,
	}
}

// PublicKey returns the public key of the token.
func (t *Token) PublicKey() (string, error) {
	claim, err := t.Claim()
	if err != nil {
		return "", err
	}

	return claim.Subject, nil
}

// Claim is returning the claim of the token.
func (t *Token) Claim() (*jwt.GenericClaims, error) {
	claim, err := jwt.DecodeGeneric(t.Token)
	if err != nil {
		return nil, err
	}

	return claim, nil
}
