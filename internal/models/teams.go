package models

import (
	"time"

	"github.com/google/uuid"
	openapi "github.com/zeiss/typhoon/api"
	"gorm.io/gorm"
)

// Team ...
type Team struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string
	Description *string
	Systems     *[]System `gorm:"many2many:team_systems;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	gorm.Model
}

// User ...
type User struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Role ...
type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	gorm.Model
}

// Permission ...
type Permission struct {
	ID          uint `gorm:"primaryKey"`
	Slug        string
	Description *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	gorm.Model
}

// RolePermission ...
type RolePermission struct {
	ID uint `gorm:"primaryKey"`

	RoleID uuid.UUID
	Role   Role

	PermissionID uuid.UUID
	Permission   Permission
}

// UserRole ...
type UserRole struct {
	ID uint `gorm:"primaryKey"`

	UserID uuid.UUID
	User   User

	RoleID uuid.UUID
	Role   Role

	TeamID uint
	Team   Team

	gorm.Model
}

// UserTeam ...
type UserTeam struct {
	ID uint `gorm:"primaryKey"`

	UserID uuid.UUID
	User   User

	TeamID uint
	Team   Team
}

// System ...
type System struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string
	Description *string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	gorm.Model
}

// PaginatedListTeams ...
type PaginatedListTeams struct {
	Limit   *float32        `json:"limit,omitempty"`
	Offset  *float32        `json:"offset,omitempty"`
	Results *[]openapi.Team `json:"results,omitempty"`
	Total   *float32        `json:"total,omitempty"`
}
