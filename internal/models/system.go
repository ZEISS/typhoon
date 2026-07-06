package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FromAPI ...
type FromAPI[T any] interface {
	FromAPI(api *T)
}

// ToAPI ...
type ToAPI[T any] interface {
	ToAPI() *T
}

// System ...
type System struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Operator    Operator       `json:"operator" gorm:"foreignKey:OperatorID" validate:"-"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name        string         `json:"name" gorm:"unique" xml:"name" form:"name" validate:"required,min=3,max=128"`
	Description string         `json:"description" form:"description" validate:"max=1024"`
	Clusters    []Cluster      `json:"clusters" gorm:"foreignKey:SystemID"`
	Tags        []Tag          `json:"tags" gorm:"many2many:system_tags;"`
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid()" xml:"id" params:"id"`
	OperatorID  uuid.UUID      `json:"operator_id" form:"operator_id" validate:"required,uuid"`
}
