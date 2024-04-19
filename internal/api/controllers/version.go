package controllers

import "github.com/zeiss/typhoon/internal/api/ports"

// VersionController ...
type VersionController struct {
	v ports.Build
}

// NewVersionController ...
func NewVersionController(v ports.Build) *VersionController {
	return &VersionController{v}
}

// Version ...
func (c *VersionController) Version() (string, error) {
	return c.v.Version()
}

// Build ...
func (c *VersionController) Build() (string, error) {
	return c.v.Build()
}

// Date ...
func (c *VersionController) Date() (string, error) {
	return c.v.Date()
}
