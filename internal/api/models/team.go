package models

import (
	authz "github.com/zeiss/fiber-authz"
)

// Team ...
type Team struct {
	*authz.Team
	// The systems that the teams have access to.
	Systems []*System `gorm:"many2many:team_systems;"`
}
