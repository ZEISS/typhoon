package services

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zeiss/typhoon/internal/api/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
	"gorm.io/gorm"
)

// ListOperatorAccountSigningKeys ...
func (a *ApiHandlers) ListOperatorAccountSigningKeys(ctx context.Context, req openapi.ListOperatorAccountSigningKeysRequestObject) (openapi.ListOperatorAccountSigningKeysResponseObject, error) {
	pagination := models.NewPagination[models.NKey]()
	pagination.Limit = 0
	pagination.Offset = 0

	keys, err := a.accounts.ListSigningKeys(ctx, req.AccountId, pagination)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return openapi.ListOperatorAccountSigningKeysdefaultJSONResponse{StatusCode: 404, Body: openapi.ErrorNotFound(err.Error())}, nil
	}

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	kp := make([]openapi.KeyPair, len(keys.Rows))
	for i := range keys.Rows {
		pk, err := keys.Rows[i].PublicKey()
		if err != nil {
			return openapi.ListOperatorAccountSigningKeysdefaultJSONResponse{}, err
		}

		kp[i] = openapi.KeyPair{
			PublicKey: pk,
		}
	}

	return openapi.ListOperatorAccountSigningKeys200JSONResponse(openapi.ListOperatorAccountSigningKeys200JSONResponse{Results: &kp}), nil
}
