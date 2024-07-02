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
	// ID is the unique identifier for the operator.
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	// Name is the name of the operator.
	Name string `json:"name" validate:"required,min=3,max=128"`
	// Description is the description of the operator.
	Description string `json:"description" validate:"max=1024"`
	// Key is the issuer key identifier.
	Key NKey `json:"key" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:operator"`
	// Token is the JWT token used to authenticate the account.
	Token Token `json:"token" gorm:"foreignKey:ID;polymorphic:Owner;polymorphicValue:operator"`
	// SystemAccount is the account that is used to manage the systems.
	SystemAccount Account `json:"system_account" gorm:"Constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// SystemAccountID is the account that is used to manage the systems.
	SystemAccountID uuid.UUID `json:"system_account_id" gorm:"type:uuid"`
	// Accounts is the accounts that are associated with the operator.
	Accounts []Account `json:"accounts" gorm:"many2many:operator_accounts;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:AccountID"`
	// Systems is the systems that are associated with the operator.
	Systems []System `json:"systems" gorm:"foreignKey:OperatorID"`
	// SigningKeyGroups is the list of signing key groups the account has.
	SigningKeyGroups []SigningKeyGroup `json:"signing_key_groups" gorm:"many2many:operator_signing_key_groups;foreignKey:ID;joinForeignKey:OperatorID;joinReferences:SigningKeyGroupID"`
	// CreatedAt is the time the operator was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the operator was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the time the operator was deleted.
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
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
