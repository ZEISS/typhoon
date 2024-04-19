package apis

import (
	"context"
	"errors"
)

// Unimplemented ...
type Unimplemented struct{}

var _ StrictServerInterface = (*Unimplemented)(nil)

// ListSystems ...
func (u *Unimplemented) ListSystems(ctx context.Context, request ListSystemsRequestObject) (ListSystemsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ShowSystem ...
func (u *Unimplemented) ShowSystem(ctx context.Context, request ShowSystemRequestObject) (ShowSystemResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListTeams ...
func (u *Unimplemented) ListTeams(ctx context.Context, request ListTeamsRequestObject) (ListTeamsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// CreateTeam ...
func (u *Unimplemented) CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetTeam ...
func (u *Unimplemented) GetTeam(ctx context.Context, request GetTeamRequestObject) (GetTeamResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListAccounts ...
func (u *Unimplemented) ListAccounts(ctx context.Context, request ListAccountsRequestObject) (ListAccountsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetAccount ...
func (u *Unimplemented) GetAccount(ctx context.Context, request GetAccountRequestObject) (GetAccountResponseObject, error) {
	return nil, errors.New("not implemented")
}

// Version ...
func (u *Unimplemented) Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error) {
	return nil, errors.New("not implemented")
}
