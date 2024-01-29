package controllers

import (
	"context"

	openapi "github.com/zeiss/typhoon/api"
	"github.com/zeiss/typhoon/internal/ports"
)

var _ ports.Teams = (*Teams)(nil)

// Teams ...
type Teams struct {
	port ports.Teams
}

// NewTeamsController ...
func NewTeamsController(port ports.Teams) *Teams {
	return &Teams{port}
}

// CreateTeam ...
func (t *Teams) CreateTeam(ctx context.Context, team *openapi.Team) error {
	return t.port.CreateTeam(ctx, team)
}

// ListTeams ...
func (t *Teams) ListTeams(ctx context.Context) ([]*openapi.Team, error) {
	return t.port.ListTeams(ctx)
}
