package apis

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

var _ StrictServerInterface = (*Unimplemented)(nil)

// CreateOperator ...
func (u *Unimplemented) CreateOperator(ctx context.Context, request CreateOperatorRequestObject) (CreateOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateOperatorSigningKeyGroup ...
func (u *Unimplemented) CreateOperatorSigningKeyGroup(ctx context.Context, request CreateOperatorSigningKeyGroupRequestObject) (CreateOperatorSigningKeyGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListOperatorSigningKeyGroups ...
func (u *Unimplemented) ListOperatorSigningKeyGroups(ctx context.Context, request ListOperatorSigningKeyGroupsRequestObject) (ListOperatorSigningKeyGroupsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListAccounts ...
func (u *Unimplemented) ListOperatorAccounts(ctx context.Context, request ListOperatorAccountsRequestObject) (ListOperatorAccountsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateOperatorAccount ...
func (u *Unimplemented) CreateOperatorAccount(ctx context.Context, request CreateOperatorAccountRequestObject) (CreateOperatorAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteOperatorAccount ...
func (u *Unimplemented) DeleteOperatorAccount(ctx context.Context, request DeleteOperatorAccountRequestObject) (DeleteOperatorAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorAccount ...
func (u *Unimplemented) GetOperatorAccount(ctx context.Context, request GetOperatorAccountRequestObject) (GetOperatorAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperator ...
func (u *Unimplemented) GetOperator(ctx context.Context, request GetOperatorRequestObject) (GetOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteOperator ...
func (u *Unimplemented) DeleteOperator(ctx context.Context, request DeleteOperatorRequestObject) (DeleteOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateOperator ...
func (u *Unimplemented) UpdateOperator(ctx context.Context, request UpdateOperatorRequestObject) (UpdateOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListOperators ...
func (u *Unimplemented) ListOperators(ctx context.Context, request ListOperatorsRequestObject) (ListOperatorsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteOperatorAccountToken ...
func (u *Unimplemented) DeleteOperatorAccountToken(ctx context.Context, request DeleteOperatorAccountTokenRequestObject) (DeleteOperatorAccountTokenResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorAccountToken ...
func (u *Unimplemented) GetOperatorAccountToken(ctx context.Context, request GetOperatorAccountTokenRequestObject) (GetOperatorAccountTokenResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateOperatorAccount ...
func (u *Unimplemented) UpdateOperatorAccount(ctx context.Context, request UpdateOperatorAccountRequestObject) (UpdateOperatorAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteSigningKeyGroup ...
func (u *Unimplemented) DeleteSigningKeyGroup(ctx context.Context, request DeleteSigningKeyGroupRequestObject) (DeleteSigningKeyGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListOperatorAccountUsers ...
func (u *Unimplemented) ListOperatorAccountUsers(ctx context.Context, request ListOperatorAccountUsersRequestObject) (ListOperatorAccountUsersResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorAccountUser ...
func (u *Unimplemented) GetOperatorAccountUser(ctx context.Context, request GetOperatorAccountUserRequestObject) (GetOperatorAccountUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorAccountUserCredentials ...
func (u *Unimplemented) GetOperatorAccountUserCredentials(ctx context.Context, request GetOperatorAccountUserCredentialsRequestObject) (GetOperatorAccountUserCredentialsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateOperatorAccountUser ...
func (u *Unimplemented) CreateOperatorAccountUser(ctx context.Context, request CreateOperatorAccountUserRequestObject) (CreateOperatorAccountUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteTeamAccountUser ...
func (u *Unimplemented) DeleteTeamAccountUser(ctx context.Context, request DeleteTeamAccountUserRequestObject) (DeleteTeamAccountUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorAccountUserToken ...
func (u *Unimplemented) GetOperatorAccountUserToken(ctx context.Context, request GetOperatorAccountUserTokenRequestObject) (GetOperatorAccountUserTokenResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorToken ...
func (u *Unimplemented) GetOperatorToken(ctx context.Context, request GetOperatorTokenRequestObject) (GetOperatorTokenResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListSystems ...
func (u *Unimplemented) ListSystems(ctx context.Context, request ListSystemsRequestObject) (ListSystemsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateSystem ...
func (u *Unimplemented) CreateSystem(ctx context.Context, request CreateSystemRequestObject) (CreateSystemResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetSystem ...
func (u *Unimplemented) GetSystem(ctx context.Context, request GetSystemRequestObject) (GetSystemResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateSystem ...
func (u *Unimplemented) UpdateSystem(ctx context.Context, request UpdateSystemRequestObject) (UpdateSystemResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteSystem ...
func (u *Unimplemented) DeleteSystem(ctx context.Context, request DeleteSystemRequestObject) (DeleteSystemResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteSystemOperator ...
func (u *Unimplemented) DeleteSystemOperator(ctx context.Context, request DeleteSystemOperatorRequestObject) (DeleteSystemOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetSystemOperator ...
func (u *Unimplemented) GetSystemOperator(ctx context.Context, request GetSystemOperatorRequestObject) (GetSystemOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateSystemOperator ...
func (u *Unimplemented) UpdateSystemOperator(ctx context.Context, request UpdateSystemOperatorRequestObject) (UpdateSystemOperatorResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListTeams ...
func (u *Unimplemented) ListTeams(ctx context.Context, request ListTeamsRequestObject) (ListTeamsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListTeamSystems ...
func (u *Unimplemented) ListTeamSystems(ctx context.Context, request ListTeamSystemsRequestObject) (ListTeamSystemsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateTeam ...
func (u *Unimplemented) CreateTeam(ctx context.Context, request CreateTeamRequestObject) (CreateTeamResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetTeam ...
func (u *Unimplemented) GetTeam(ctx context.Context, request GetTeamRequestObject) (GetTeamResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteTeam ...
func (u *Unimplemented) DeleteTeam(ctx context.Context, request DeleteTeamRequestObject) (DeleteTeamResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListAccounts ...
func (u *Unimplemented) ListTeamAccounts(ctx context.Context, request ListTeamAccountsRequestObject) (ListTeamAccountsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetAccount ...
func (u *Unimplemented) GetAccount(ctx context.Context, request GetAccountRequestObject) (GetAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateGroup ...
func (u *Unimplemented) CreateGroup(ctx context.Context, request CreateGroupRequestObject) (CreateGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetGroup ...
func (u *Unimplemented) GetGroup(ctx context.Context, request GetGroupRequestObject) (GetGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateGroup ...
func (u *Unimplemented) UpdateGroup(ctx context.Context, request UpdateGroupRequestObject) (UpdateGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListGroups ...
func (u *Unimplemented) ListGroups(ctx context.Context, request ListGroupsRequestObject) (ListGroupsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListUsers ...
func (u *Unimplemented) ListUsers(ctx context.Context, request ListUsersRequestObject) (ListUsersResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetUser ...
func (u *Unimplemented) GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateUser ...
func (u *Unimplemented) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateUser ...
func (u *Unimplemented) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// Version ...
func (u *Unimplemented) Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}
