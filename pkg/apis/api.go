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

// GetOperatorToken ...
func (u *Unimplemented) GetOperatorToken(ctx context.Context, request GetOperatorTokenRequestObject) (GetOperatorTokenResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// GetOperatorSystemAccount ...
func (u *Unimplemented) GetOperatorSystemAccount(ctx context.Context, request GetOperatorSystemAccountRequestObject) (GetOperatorSystemAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateOperatorSystemAccount ...
func (u *Unimplemented) UpdateOperatorSystemAccount(ctx context.Context, request UpdateOperatorSystemAccountRequestObject) (UpdateOperatorSystemAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateAccount ...
func (u *Unimplemented) CreateAccount(ctx context.Context, request CreateAccountRequestObject) (CreateAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteAccount ...
func (u *Unimplemented) DeleteAccount(ctx context.Context, request DeleteAccountRequestObject) (DeleteAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// UpdateAccount ...
func (u *Unimplemented) UpdateAccount(ctx context.Context, request UpdateAccountRequestObject) (UpdateAccountResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListAccounts ...
func (u *Unimplemented) ListAccounts(ctx context.Context, request ListAccountsRequestObject) (ListAccountsResponseObject, error) {
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

// Version ...
func (u *Unimplemented) Version(ctx context.Context, request VersionRequestObject) (VersionResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// CreateAccountSigningKeyGroup ...
func (u *Unimplemented) CreateAccountSigningKeyGroup(ctx context.Context, request CreateAccountSigningKeyGroupRequestObject) (CreateAccountSigningKeyGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// ListAccountSigningKeyGroups ...
func (u *Unimplemented) ListAccountSigningKeyGroups(ctx context.Context, request ListAccountSigningKeyGroupsRequestObject) (ListAccountSigningKeyGroupsResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}

// DeleteAccountSigningKeyGroup ...
func (u *Unimplemented) DeleteAccountSigningKeyGroup(ctx context.Context, request DeleteAccountSigningKeyGroupRequestObject) (DeleteAccountSigningKeyGroupResponseObject, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "not implemented")
}
