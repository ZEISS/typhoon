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

// CreateSystem ...
func (u *Unimplemented) CreateSystem(ctx context.Context, request CreateSystemRequestObject) (CreateSystemResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetSystem ...
func (u *Unimplemented) GetSystem(ctx context.Context, request GetSystemRequestObject) (GetSystemResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateSystem ...
func (u *Unimplemented) UpdateSystem(ctx context.Context, request UpdateSystemRequestObject) (UpdateSystemResponseObject, error) {
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

// CreateGroup ...
func (u *Unimplemented) CreateGroup(ctx context.Context, request CreateGroupRequestObject) (CreateGroupResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetGroup ...
func (u *Unimplemented) GetGroup(ctx context.Context, request GetGroupRequestObject) (GetGroupResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateGroup ...
func (u *Unimplemented) UpdateGroup(ctx context.Context, request UpdateGroupRequestObject) (UpdateGroupResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListGroups ...
func (u *Unimplemented) ListGroups(ctx context.Context, request ListGroupsRequestObject) (ListGroupsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListUsers ...
func (u *Unimplemented) ListUsers(ctx context.Context, request ListUsersRequestObject) (ListUsersResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetUser ...
func (u *Unimplemented) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// CreateUser ...
func (u *Unimplemented) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateUser ...
func (u *Unimplemented) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// Version ...
func (u *Unimplemented) Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error) {
	return nil, errors.New("not implemented")
}
