package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// System ...
type System struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`

	Cluster   Cluster   `json:"cluster"`
	ClusterID uuid.UUID `json:"cluster_id" gorm:"foreignKey:Cluster"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Cluster ...
type Cluster struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	URL             string    `json:"url"`
	Operator        Operator  `json:"operator"`
	OperatorID      uuid.UUID `json:"operator_id" gorm:"foreignKey:Operator"`
	SystemAccount   Account   `json:"system_account"`
	SystemAccountID uuid.UUID `json:"system_account_id" gorm:"foreignKey:Account"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Account ...
type Account struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name  string    `json:"name"`
	Key   Key       `json:"key"`
	KeyID uuid.UUID `json:"key_id" gorm:"foreignKey:Key"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Key ...
type Key struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Public string    `json:"public"`
	Seed   []byte    `json:"seed"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Group ...
type Group struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name"`

	Key   Key       `json:"key"`
	KeyID uuid.UUID `json:"key_id" gorm:"foreignKey:Key"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Operator ...
type Operator struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Keys []Key     `json:"keys" gorm:"foreignKey:ID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// User ...
type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	Key   Key       `json:"key"`
	KeyID uuid.UUID `json:"key_id" gorm:"foreignKey:Key"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
