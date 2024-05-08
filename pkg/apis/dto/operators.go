package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromGetOperatorRequest ...
func FromGetOperatorRequest(req openapi.GetOperatorRequestObject) controllers.GetOperatorQuery {
	return controllers.GetOperatorQuery{
		ID: req.OperatorId,
	}
}

// ToGetOperatorResponse ...
func ToGetOperatorResponse(op models.Operator) openapi.GetOperator200JSONResponse {
	response := openapi.GetOperator200JSONResponse{}

	response.Id = utils.PtrUUID(op.ID)
	response.Name = op.Name
	response.Description = utils.StrPtr(op.Description)
	response.CreatedAt = utils.PtrTime(op.CreatedAt)
	response.UpdatedAt = utils.PtrTime(op.UpdatedAt)
	response.DeletedAt = utils.PtrTime(op.DeletedAt.Time)

	return response
}

// FromGetOperatorTokenRequest ...
func FromGetOperatorTokenRequest(req openapi.GetOperatorTokenRequestObject) controllers.GetOperatorTokenQuery {
	return controllers.GetOperatorTokenQuery{
		ID: req.OperatorId,
	}
}

// ToGetOperatorTokenResponse ...
func ToGetOperatorTokenResponse(token models.Token) openapi.GetOperatorToken200JSONResponse {
	response := openapi.GetOperatorToken200JSONResponse{}

	response.Token = utils.StrPtr(token.Token)

	return response
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
	response := openapi.CreateOperator201JSONResponse{}

	response.Id = utils.PtrUUID(op.ID)
	response.Name = op.Name
	response.Description = utils.StrPtr(op.Description)
	response.CreatedAt = utils.PtrTime(op.CreatedAt)
	response.UpdatedAt = utils.PtrTime(op.UpdatedAt)
	response.DeletedAt = utils.PtrTime(op.DeletedAt.Time)

	return response
}
