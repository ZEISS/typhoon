package controllers

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
)

// TeamsController ...
type TeamsController interface {
	// CreateTeam ...
	CreateTeam(ctx context.Context, name string, description *string) (models.Team, error)
	// DeleteTeam ...
	DeleteTeam(ctx context.Context, id uuid.UUID) error
	// GetTeam ...
	GetTeam(ctx context.Context, id uuid.UUID) (models.Team, error)
}

type teamsController struct {
	db ports.Teams
}

// NewTeamsController ...
func NewTeamsController(db ports.Teams) *teamsController {
	return &teamsController{db}
}

// CreateTeam ...
func (c *teamsController) CreateTeam(ctx context.Context, name string, description *string) (models.Team, error) {
	return c.db.CreateTeam(ctx, models.Team{Team: &authz.Team{Name: name, Description: description}})
}

// DeleteTeam ...
func (c *teamsController) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	return c.db.DeleteTeam(ctx, id)
}

// GetTeam ...
func (c *teamsController) GetTeam(ctx context.Context, id uuid.UUID) (models.Team, error) {
	return c.db.GetTeam(ctx, id)
}
