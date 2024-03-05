package apis

import (
	"context"
)

var _ StrictServerInterface = (*Unimplemented)(nil)

// Unimplemented ...
type Unimplemented struct{}

// ListSystems ...
func (u *Unimplemented) ListSystems(ctx context.Context, request ListSystemsRequestObject) (ListSystemsResponseObject, error) {
	return nil, nil
}

// CreateTeam ...
func (u *Unimplemented) CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	return nil, nil
}

// GetTeamTeamId ...
func (u *Unimplemented) GetTeamTeamId(ctx context.Context, request GetTeamTeamIdRequestObject) (GetTeamTeamIdResponseObject, error) {
	return nil, nil
}

// ListTeam ...
func (u *Unimplemented) ListTeam(ctx context.Context, request ListTeamRequestObject) (ListTeamResponseObject, error) {
	return nil, nil
}

// ShowSystem ...
func (u *Unimplemented) ShowSystem(ctx context.Context, request ShowSystemRequestObject) (ShowSystemResponseObject, error) {
	return nil, nil
}

// Version ...
func (u *Unimplemented) Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error) {
	return nil, nil
}
