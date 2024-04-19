package controllers

import "github.com/zeiss/typhoon/internal/api/ports"

// TeamsController ...
type TeamsController struct {
	db ports.Teams
}

// NewTeamsController ...
func NewTeamsController(db ports.Teams) *TeamsController {
	return &TeamsController{db}
}
