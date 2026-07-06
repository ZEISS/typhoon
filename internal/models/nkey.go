package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nkeys"
	"gorm.io/gorm"
)

const (
	// OperatorNKey is the operator nkey.
	OperatorNKey OwnerType = "operator"
	// AccountNKey is the account nkey.
	AccountNKey OwnerType = "account"
	// UserNKey is the user nkey.
	UserNKey OwnerType = "user"
)

// NKey holds a private key and its metadata.
type NKey struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ID        string         `json:"id" gorm:"primaryKey"`
	OwnerType OwnerType      `json:"owner_type"`
	Seed      []byte         `json:"seed"`
	OwnerID   uuid.UUID      `json:"owner_id"`
}

// KeyPair is a pair of NKeys.
func (n *NKey) KeyPair() (nkeys.KeyPair, error) {
	kp, err := nkeys.FromSeed(n.Seed)
	if err != nil {
		return nil, err
	}

	return kp, nil
}

// PublicKey returns the public key portion of the NKey.
func (n *NKey) PublicKey() (string, error) {
	kp, err := n.KeyPair()
	if err != nil {
		return "", err
	}

	return kp.PublicKey()
}

// PrivateKey returns the private key portion of the NKey.
func (n *NKey) PrivateKey() ([]byte, error) {
	kp, err := n.KeyPair()
	if err != nil {
		return nil, err
	}

	return kp.PrivateKey()
}
