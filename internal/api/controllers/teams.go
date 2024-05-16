package controllers

import (
	"context"

	"github.com/google/uuid"
	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/api/ports"
	"github.com/zeiss/typhoon/internal/utils"
)

// CreateTeamCommand ...
type CreateTeamCommand struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=1024"`
}

// DeleteTeamCommand ...
type DeleteTeamCommand struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// GetTeamQuery ...
type GetTeamQuery struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// ListTeamsQuery ...
type ListTeamsQuery struct {
	Offset int    `json:"offset" validate:"required"`
	Limit  int    `json:"limit" validate:"required"`
	Sort   string `json:"sort" validate:"required"`
	Search string `json:"search" validate:"required"`
}

var _ TeamsController = (*teamsController)(nil)

// TeamsController ...
type TeamsController interface {
	// CreateTeam ...
	CreateTeam(ctx context.Context, cmd CreateTeamCommand) (authz.Team, error)
	// DeleteTeam ...
	DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error
	// GetTeam ...
	GetTeam(ctx context.Context, query GetTeamQuery) (authz.Team, error)
	// ListTeams ...
	ListTeams(ctx context.Context, query ListTeamsQuery) (models.Pagination[authz.Team], error)
}

type teamsController struct {
	db ports.Teams
}

// NewTeamsController ...
func NewTeamsController(db ports.Teams) *teamsController {
	return &teamsController{db}
}

// CreateTeam ...
func (c *teamsController) CreateTeam(ctx context.Context, cmd CreateTeamCommand) (authz.Team, error) {
	team := authz.Team{
		Name:        cmd.Name,
		Description: utils.StrPtr(cmd.Description),
	}

	err := c.db.CreateTeam(ctx, &team)
	if err != nil {
		return team, err
	}

	return team, err
}

// DeleteTeam ...
func (c *teamsController) DeleteTeam(ctx context.Context, cmd DeleteTeamCommand) error {
	team := authz.Team{
		ID: cmd.ID,
	}

	return c.db.DeleteTeam(ctx, &team)
}

// GetTeam ...
func (c *teamsController) GetTeam(ctx context.Context, query GetTeamQuery) (authz.Team, error) {
	team := authz.Team{
		ID: query.ID,
	}

	err := c.db.GetTeam(ctx, &team)
	if err != nil {
		return team, err
	}

	return team, nil
}

// ListTeams ...
func (c *teamsController) ListTeams(ctx context.Context, query ListTeamsQuery) (models.Pagination[authz.Team], error) {
	pagination := models.Pagination[authz.Team]{
		Offset: query.Offset,
		Limit:  query.Limit,
		Search: query.Search,
	}

	err := c.db.ListTeams(ctx, &pagination)
	if err != nil {
		return pagination, err
	}

	return pagination, nil
}
