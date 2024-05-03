package controllers

import "github.com/zeiss/typhoon/internal/api/ports"

// TeamsController ...
type TeamsController interface{}

type teamsController struct {
	db ports.Teams
}

// NewTeamsController ...
func NewTeamsController(db ports.Teams) *teamsController {
	return &teamsController{db}
}
