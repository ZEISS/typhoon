package apis

import (
	"context"
	"errors"
)

var _ StrictServerInterface = (*Unimplemented)(nil)

// CreateOperator ...
func (u *Unimplemented) CreateOperator(ctx context.Context, request CreateOperatorRequestObject) (CreateOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListAccounts ...
func (u *Unimplemented) ListOperatorAccounts(ctx context.Context, request ListOperatorAccountsRequestObject) (ListOperatorAccountsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// CreateOperatorAccount ...
func (u *Unimplemented) CreateOperatorAccount(ctx context.Context, request CreateOperatorAccountRequestObject) (CreateOperatorAccountResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteOperatorAccount ...
func (u *Unimplemented) DeleteOperatorAccount(ctx context.Context, request DeleteOperatorAccountRequestObject) (DeleteOperatorAccountResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorAccount ...
func (u *Unimplemented) GetOperatorAccount(ctx context.Context, request GetOperatorAccountRequestObject) (GetOperatorAccountResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListOperatorSignKeys ...
func (u *Unimplemented) ListOperatorSignKeys(ctx context.Context, request ListOperatorSigningKeysRequestObject) (ListOperatorSigningKeysResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListOperatorSigningKeys ...
func (u *Unimplemented) ListOperatorSigningKeys(ctx context.Context, request ListOperatorSigningKeysRequestObject) (ListOperatorSigningKeysResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperator ...
func (u *Unimplemented) GetOperator(ctx context.Context, request GetOperatorRequestObject) (GetOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteOperator ...
func (u *Unimplemented) DeleteOperator(ctx context.Context, request DeleteOperatorRequestObject) (DeleteOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateOperator ...
func (u *Unimplemented) UpdateOperator(ctx context.Context, request UpdateOperatorRequestObject) (UpdateOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListOperators ...
func (u *Unimplemented) ListOperator(ctx context.Context, request ListOperatorRequestObject) (ListOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteOperatorAccountToken ...
func (u *Unimplemented) DeleteOperatorAccountToken(ctx context.Context, request DeleteOperatorAccountTokenRequestObject) (DeleteOperatorAccountTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorAccountToken ...
func (u *Unimplemented) GetOperatorAccountToken(ctx context.Context, request GetOperatorAccountTokenRequestObject) (GetOperatorAccountTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateOperatorAccount ...
func (u *Unimplemented) UpdateOperatorAccount(ctx context.Context, request UpdateOperatorAccountRequestObject) (UpdateOperatorAccountResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListOperatorAccountSigningKeys ...
func (u *Unimplemented) ListOperatorAccountSigningKeys(ctx context.Context, request ListOperatorAccountSigningKeysRequestObject) (ListOperatorAccountSigningKeysResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteSigningKeyGroup ...
func (u *Unimplemented) DeleteSigningKeyGroup(ctx context.Context, request DeleteSigningKeyGroupRequestObject) (DeleteSigningKeyGroupResponseObject, error) {
	return nil, errors.New("not implemented")
}

// ListOperatorAccountUsers ...
func (u *Unimplemented) ListOperatorAccountUsers(ctx context.Context, request ListOperatorAccountUsersRequestObject) (ListOperatorAccountUsersResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorAccountUser ...
func (u *Unimplemented) GetOperatorAccountUser(ctx context.Context, request GetOperatorAccountUserRequestObject) (GetOperatorAccountUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorAccountUserCredentials ...
func (u *Unimplemented) GetOperatorAccountUserCredentials(ctx context.Context, request GetOperatorAccountUserCredentialsRequestObject) (GetOperatorAccountUserCredentialsResponseObject, error) {
	return nil, errors.New("not implemented")
}

// CreateOperatorAccountUser ...
func (u *Unimplemented) CreateOperatorAccountUser(ctx context.Context, request CreateOperatorAccountUserRequestObject) (CreateOperatorAccountUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteTeamAccountUser ...
func (u *Unimplemented) DeleteTeamAccountUser(ctx context.Context, request DeleteTeamAccountUserRequestObject) (DeleteTeamAccountUserResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorAccountUserToken ...
func (u *Unimplemented) GetOperatorAccountUserToken(ctx context.Context, request GetOperatorAccountUserTokenRequestObject) (GetOperatorAccountUserTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetOperatorToken ...
func (u *Unimplemented) GetOperatorToken(ctx context.Context, request GetOperatorTokenRequestObject) (GetOperatorTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateOperatorToken ...
func (u *Unimplemented) UpdateOperatorToken(ctx context.Context, request UpdateOperatorTokenRequestObject) (UpdateOperatorTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteOperatorToken ...
func (u *Unimplemented) DeleteOperatorToken(ctx context.Context, request DeleteOperatorTokenRequestObject) (DeleteOperatorTokenResponseObject, error) {
	return nil, errors.New("not implemented")
}

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
	return GetSystem501JSONResponse(GetOperator501JSONResponse{UnimplementedJSONResponse(NotImplemented("Not implemented"))}), nil
}

// UpdateSystem ...
func (u *Unimplemented) UpdateSystem(ctx context.Context, request UpdateSystemRequestObject) (UpdateSystemResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteSystem ...
func (u *Unimplemented) DeleteSystem(ctx context.Context, request DeleteSystemRequestObject) (DeleteSystemResponseObject, error) {
	return nil, errors.New("not implemented")
}

// DeleteSystemOperator ...
func (u *Unimplemented) DeleteSystemOperator(ctx context.Context, request DeleteSystemOperatorRequestObject) (DeleteSystemOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// GetSystemOperator ...
func (u *Unimplemented) GetSystemOperator(ctx context.Context, request GetSystemOperatorRequestObject) (GetSystemOperatorResponseObject, error) {
	return nil, errors.New("not implemented")
}

// UpdateSystemOperator ...
func (u *Unimplemented) UpdateSystemOperator(ctx context.Context, request UpdateSystemOperatorRequestObject) (UpdateSystemOperatorResponseObject, error) {
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
func (u *Unimplemented) ListTeamAccounts(ctx context.Context, request ListTeamAccountsRequestObject) (ListTeamAccountsResponseObject, error) {
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
