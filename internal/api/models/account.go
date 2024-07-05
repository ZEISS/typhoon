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
	// Description is the description of the account.
	Description *string `json:"description"`
	// Key is the issuer key identifier.
	Key NKey `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	// Token is the JWT token used to authenticate the account.
	Token Token `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	// SigningKeyGroups is the list of signing key groups the account has.
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:account_signing_key_groups;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyGroupID"`
	// Users is the list of users the account has.
	Users []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// OwnerID is the owner of the token.
	OwnerID uuid.UUID `json:"owner_id"`
	// OwnerType is the type of the owner.
	OwnerType OwnerType `json:"owner_type"`
	// LimitJetStreamMaxDiskStorage is the limit for JetStream maximum disk storage.
	LimitJetStreamMaxDiskStorage int64 `json:"limit_jetstream_max_disk_storage"`
	// LimitJetStreamMaxStreams is the limit for JetStream maximum streams.
	LimitJetStreamMaxStreams int64 `json:"limit_jetstream_max_streams"`
	// LimitJetStreamMaxAckPending is the limit for JetStream maximum ack pending.
	LimitJetStreamMaxAckPending int64 `json:"limit_jetstream_max_ack_pending"`
	/// LimitJetStreamMaxStreamBytes is the limit for JetStream maximum stream bytes.
	LimitJetStreamMaxStreamBytes int64 `json:"limit_jetstream_max_stream_bytes"`
	// LimitJetStreamMaxBytesRequired indicates if JetStream maximum bytes required is limited.
	LimitJetStreamMaxBytesRequired bool `json:"limit_jetstream_max_bytes_required"`
	// LimitJetStreamMaxConsumers indicates if JetStream maximum consumer is limited.
	LimitJetStreamMaxConsumers int64 `json:"limit_jetstream_max_consumers"`
	// CreatedAt is the time the account was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the account was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the account was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// AfterDelete ...
func (a *Account) AfterDelete(tx *gorm.DB) error {
	return tx.Where("account_id = ?", a.ID).Delete(&User{}).Error
}

// FindSigningKeyGroupByID ...
func (a *Account) FindSigningKeyGroupByID(id uuid.UUID) *SigningKeyGroup {
	for _, skg := range a.SigningKeyGroups {
		if skg.ID == id {
			return &skg
		}
	}

	return nil
}
