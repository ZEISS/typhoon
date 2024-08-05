package dto

import (
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromListOperatorsRequest ...
func FromListOperatorsRequest(req openapi.ListOperatorsRequestObject) controllers.ListOperatorsQuery {
	return controllers.ListOperatorsQuery{
		Offset: cast.Value(req.Params.Offset),
		Limit:  cast.Value(req.Params.Limit),
	}
}

// ToListOperatorsResponse ...
func ToListOperatorsResponse(ops models.Pagination[models.Operator]) openapi.ListOperators200JSONResponse {
	res := openapi.ListOperators200JSONResponse{}

	res.Limit = cast.Ptr(ops.Limit)
	res.Offset = cast.Ptr(ops.Offset)
	res.Total = cast.Ptr(ops.TotalRows)

	operators := make([]openapi.Operator, 0, len(ops.Rows))
	for _, op := range ops.Rows {
		operator := openapi.Operator{
			Id:          cast.Ptr(op.ID),
			Name:        op.Name,
			Description: cast.Ptr(op.Description),
			CreatedAt:   cast.Ptr(op.CreatedAt),
			UpdatedAt:   cast.Ptr(op.UpdatedAt),
			DeletedAt:   cast.Ptr(op.DeletedAt.Time),
			Key:         &openapi.KeyPair{PublicKey: op.Key.ID},
		}

		skg := []openapi.KeyPair{}
		for _, sk := range op.SigningKeyGroups {
			skg = append(skg, openapi.KeyPair{PublicKey: sk.Key.ID})
		}
		operator.SigningKeys = &skg

		operators = append(operators, operator)
	}
	res.Results = &operators

	return res
}

// FromGetOperatorRequest ...
func FromGetOperatorRequest(req openapi.GetOperatorRequestObject) controllers.GetOperatorQuery {
	return controllers.GetOperatorQuery{
		ID: req.OperatorId,
	}
}

// ToGetOperatorResponse ...
func ToGetOperatorResponse(op models.Operator) openapi.GetOperator200JSONResponse {
	res := openapi.GetOperator200JSONResponse{}

	res.Id = cast.Ptr(op.ID)
	res.Name = op.Name
	res.Description = cast.Ptr(op.Description)
	res.CreatedAt = cast.Ptr(op.CreatedAt)
	res.UpdatedAt = cast.Ptr(op.UpdatedAt)
	res.DeletedAt = cast.Ptr(op.DeletedAt.Time)

	return res
}

// FromGetOperatorTokenRequest ...
func FromGetOperatorTokenRequest(req openapi.GetOperatorTokenRequestObject) controllers.GetOperatorTokenQuery {
	return controllers.GetOperatorTokenQuery{
		ID: req.OperatorId,
	}
}

// ToGetOperatorTokenResponse ...
func ToGetOperatorTokenResponse(token models.Token) openapi.GetOperatorToken200JSONResponse {
	res := openapi.GetOperatorToken200JSONResponse{}

	res.Token = cast.Ptr(token.Token)

	return res
}

// FromCreateOperatorRequest ...
func FromCreateOperatorRequest(req openapi.CreateOperatorRequestObject) controllers.CreateOperatorCommand {
	return controllers.CreateOperatorCommand{
		Name:        req.Body.Name,
		Description: cast.Value(req.Body.Description),
	}
}

// ToCreateOperatorResponse ...
func ToCreateOperatorResponse(op models.Operator) openapi.CreateOperator201JSONResponse {
	res := openapi.CreateOperator201JSONResponse{}

	res.Id = cast.Ptr(op.ID)
	res.Name = op.Name
	res.Description = cast.Ptr(op.Description)
	res.CreatedAt = cast.Ptr(op.CreatedAt)
	res.UpdatedAt = cast.Ptr(op.UpdatedAt)
	res.DeletedAt = cast.Ptr(op.DeletedAt.Time)

	return res
}

// FromDeleteOperatorRequest ...
func FromDeleteOperatorRequest(req openapi.DeleteOperatorRequestObject) controllers.DeleteOperatorCommand {
	return controllers.DeleteOperatorCommand{
		ID: req.OperatorId,
	}
}

// ToDeleteOperatorResponse ...
func ToDeleteOperatorResponse() openapi.DeleteOperator204Response {
	res := openapi.DeleteOperator204Response{}

	return res
}

// FromCreateSigningKeyGroupRequest ...
func FromCreateSigningKeyGroupRequest(req openapi.CreateOperatorSigningKeyGroupRequestObject) controllers.CreateOperatorSigningKeyGroupCommand {
	return controllers.CreateOperatorSigningKeyGroupCommand{
		OperatorID:  req.OperatorId,
		Name:        req.Body.Name,
		Description: cast.Value(req.Body.Description),
	}
}

// ToCreateOperatorSigningKeyGroupResponse ...
func ToCreateOperatorSigningKeyGroupResponse(skg models.SigningKeyGroup) openapi.CreateOperatorSigningKeyGroup201JSONResponse {
	res := openapi.CreateOperatorSigningKeyGroup201JSONResponse{}

	res.Id = cast.Ptr(skg.ID)
	res.Name = skg.Name
	res.Description = cast.Ptr(skg.Description)
	res.CreatedAt = cast.Ptr(skg.CreatedAt)
	res.UpdatedAt = cast.Ptr(skg.UpdatedAt)
	res.DeletedAt = cast.Ptr(skg.DeletedAt.Time)

	return res
}

// FromGetOperatorSystemAccountRequest ...
func FromGetOperatorSystemAccountRequest(req openapi.GetOperatorSystemAccountRequestObject) controllers.GetOperatorSystemAccountQuery {
	return controllers.GetOperatorSystemAccountQuery{
		OperatorID: req.OperatorId,
	}
}

// ToGetOperatorSystemAccountResponse ...
func ToGetOperatorSystemAccountResponse(account models.Account) openapi.GetOperatorSystemAccount200JSONResponse {
	res := openapi.GetOperatorSystemAccount200JSONResponse{}

	res.Id = cast.Ptr(account.ID)
	res.Name = account.Name
	res.CreatedAt = cast.Ptr(account.CreatedAt)
	res.UpdatedAt = cast.Ptr(account.UpdatedAt)
	res.DeletedAt = cast.Ptr(account.DeletedAt.Time)

	return res
}

// FromUpdateOperatorSystemAccountRequest ...
func FromUpdateOperatorSystemAccountRequest(req openapi.UpdateOperatorSystemAccountRequestObject) controllers.UpdateOperatorSystemAccountCommand {
	return controllers.UpdateOperatorSystemAccountCommand{
		OperatorID: req.OperatorId,
		AccountID:  req.Body.AccountId,
	}
}

// ToUpdateOperatorSystemAccountResponse ...
func ToUpdateOperatorSystemAccountResponse(account models.Account) openapi.UpdateOperatorSystemAccount201JSONResponse {
	res := openapi.UpdateOperatorSystemAccount201JSONResponse{}

	res.Id = cast.Ptr(account.ID)
	res.Name = account.Name
	res.CreatedAt = cast.Ptr(account.CreatedAt)
	res.UpdatedAt = cast.Ptr(account.UpdatedAt)
	res.DeletedAt = cast.Ptr(account.DeletedAt.Time)

	return res
}
