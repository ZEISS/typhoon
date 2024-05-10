package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromListOperatorsRequest ...
func FromListOperatorsRequest(req openapi.ListOperatorsRequestObject) controllers.ListOperatorsQuery {
	return controllers.ListOperatorsQuery{
		Offset: utils.IntPtr(req.Params.Offset),
		Limit:  utils.IntPtr(req.Params.Limit),
	}
}

// ToListOperatorsResponse ...
func ToListOperatorsResponse(ops models.Pagination[models.Operator]) openapi.ListOperators200JSONResponse {
	res := openapi.ListOperators200JSONResponse{}
	res.Limit = utils.PtrInt(ops.Limit)
	res.Offset = utils.PtrInt(ops.Offset)
	res.Total = utils.PtrInt(ops.TotalRows)

	operators := make([]openapi.Operator, 0, len(ops.Rows))
	for _, op := range ops.Rows {
		operators = append(operators, openapi.Operator{
			Id:        utils.PtrUUID(op.ID),
			Name:      op.Name,
			CreatedAt: utils.PtrTime(op.CreatedAt),
			UpdatedAt: utils.PtrTime(op.UpdatedAt),
			DeletedAt: utils.PtrTime(op.DeletedAt.Time),
		})
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

	res.Id = utils.PtrUUID(op.ID)
	res.Name = op.Name
	res.Description = utils.StrPtr(op.Description)
	res.CreatedAt = utils.PtrTime(op.CreatedAt)
	res.UpdatedAt = utils.PtrTime(op.UpdatedAt)
	res.DeletedAt = utils.PtrTime(op.DeletedAt.Time)

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

	res.Token = utils.StrPtr(token.Token)

	return res
}

// FromCreateOperatorRequest ...
func FromCreateOperatorRequest(req openapi.CreateOperatorRequestObject) controllers.CreateOperatorCommand {
	return controllers.CreateOperatorCommand{
		Name:        req.Body.Name,
		Description: utils.PtrStr(req.Body.Description),
	}
}

// ToCreateOperatorResponse ...
func ToCreateOperatorResponse(op models.Operator) openapi.CreateOperator201JSONResponse {
	res := openapi.CreateOperator201JSONResponse{}

	res.Id = utils.PtrUUID(op.ID)
	res.Name = op.Name
	res.Description = utils.StrPtr(op.Description)
	res.CreatedAt = utils.PtrTime(op.CreatedAt)
	res.UpdatedAt = utils.PtrTime(op.UpdatedAt)
	res.DeletedAt = utils.PtrTime(op.DeletedAt.Time)

	return res
}
