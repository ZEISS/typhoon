package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"gorm.io/gorm"
)

// NewOperator is creating a new operator.
func NewOperator(name, description string) (Operator, error) {
	op := Operator{
		Name:        name,
		Description: description,
	}

	pk, err := nkeys.CreateOperator()
	if err != nil {
		return op, err
	}

	id, err := pk.PublicKey()
	if err != nil {
		return op, err
	}

	seed, err := pk.Seed()
	if err != nil {
		return op, err
	}

	// Create a token for the operator
	oc := jwt.NewOperatorClaims(id)
	oc.Name = name

	token, err := oc.Encode(pk)
	if err != nil {
		return op, err
	}

	op.Key = NKey{ID: id, Seed: seed}
	op.Token = Token{ID: id, Token: token}

	return op, nil
}

// OperatorPagination is the pagination for operators.
type OperatorPagination Pagination[Operator]

// Operator is the operator that is used to manage the systems.
type Operator struct {
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	SystemAccount    Account           `json:"system_account" gorm:"polymorphic:Owner;polymorphicValue:operator;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Key              NKey              `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:operator"`
	Token            Token             `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:operator"`
	DeletedAt        gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
	AccountServerURL string            `json:"account_server_url" form:"account_server_url" xml:"account_server_url" validate:"url"`
	Description      string            `json:"description" validate:"max=1024"`
	Name             string            `json:"name" validate:"required,min=3,max=128"`
	Systems          []System          `json:"systems" gorm:"foreignKey:OperatorID"`
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:operator_signing_key_groups;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyGroupID"`
	Tags             []Tag             `json:"tags" gorm:"many2many:operator_tags;"`
	ID               uuid.UUID         `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}

// Compare ...
func (o Operator) Compare() int {
	return 0
}

// Claims returns the operator claims.
func (o Operator) Claims() (*jwt.OperatorClaims, error) {
	claims, err := jwt.DecodeOperatorClaims(o.Token.Token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
