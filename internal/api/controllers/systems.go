package controllers

import "github.com/zeiss/typhoon/internal/api/ports"

// SystemsController ...
type SystemsController struct {
	db ports.Systems
}

// NewSystemsController ...
func NewSystemsController(db ports.Teams) *SystemsController {
	return &SystemsController{db}
}
