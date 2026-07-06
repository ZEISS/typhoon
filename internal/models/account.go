package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	// OperatorAccount is the operator account.
	OperatorAccount OwnerType = "operator"
	// TeamAccount is the team account.
	TeamAccount OwnerType = "team"
)

// Account ...
type Account struct {
	UpdatedAt                      time.Time         `json:"updated_at"`
	CreatedAt                      time.Time         `json:"created_at"`
	Description                    *string           `json:"description"`
	Key                            NKey              `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	Signer                         NKey              `json:"signer" gorm:"foreignKey:SignerID" validate:"-"`
	Token                          Token             `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:account"`
	DeletedAt                      gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
	SignerID                       string            `json:"signer_id" form:"signer_id" validate:"-"`
	Name                           string            `json:"name"`
	OwnerType                      OwnerType         `json:"owner_type"`
	SigningKeyGroups               []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:account_signing_key_groups;foreignKey:ID;joinForeignKey:AccountID;joinReferences:SigningKeyGroupID"`
	Users                          []User            `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LimitJetStreamMaxDiskStorage   int64             `json:"limit_jetstream_max_disk_storage"`
	LimitJetStreamMaxStreams       int64             `json:"limit_jetstream_max_streams"`
	LimitJetStreamMaxAckPending    int64             `json:"limit_jetstream_max_ack_pending"`
	LimitJetStreamMaxStreamBytes   int64             `json:"limit_jetstream_max_stream_bytes"`
	LimitJetStreamMaxConsumers     int64             `json:"limit_jetstream_max_consumers"`
	OwnerID                        uuid.UUID         `json:"owner_id"`
	ID                             uuid.UUID         `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	LimitJetStreamEnabled          bool              `json:"limit_jetstream_enabled"`
	LimitJetStreamMaxBytesRequired bool              `json:"limit_jetstream_max_bytes_required"`
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
