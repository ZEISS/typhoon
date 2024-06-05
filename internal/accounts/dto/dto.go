package dto

import (
	"bytes"

	"github.com/zeiss/typhoon/internal/accounts/models"
	openapi "github.com/zeiss/typhoon/pkg/apis/accounts"
)

// GetAccountQuery ...
type GetAccountQuery struct {
	ID models.AccountPublicKey `json:"id" validate:"required"`
}

// FromGetAccountTokenRequest ...
func FromGetAccountTokenRequest(req openapi.GetAccountTokenRequestObject) GetAccountQuery {
	return GetAccountQuery{
		ID: models.AccountPublicKey(req.PubKey),
	}
}

// ToGetAccountTokenResponse ...
func ToGetAccountTokenResponse(token models.AccountToken) openapi.GetAccountTokenResponseObject {
	res := openapi.GetAccountToken200ApplicationjwtResponse{}

	body := bytes.NewReader(token)
	res.Body = body
	res.ContentLength = int64(len(token))

	return res
}
