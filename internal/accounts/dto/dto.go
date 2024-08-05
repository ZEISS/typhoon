package dto

import (
	"bytes"

	"github.com/zeiss/typhoon/internal/models"
	openapi "github.com/zeiss/typhoon/pkg/apis/accounts"
)

// GetAccountQuery ...
type GetAccountQuery struct {
	ID string `json:"id" validate:"required"`
}

// FromGetAccountTokenRequest ...
func FromGetAccountTokenRequest(req openapi.GetAccountTokenRequestObject) GetAccountQuery {
	return GetAccountQuery{
		ID: req.PubKey,
	}
}

// ToGetAccountTokenResponse ...
func ToGetAccountTokenResponse(token models.Token) openapi.GetAccountTokenResponseObject {
	res := openapi.GetAccountToken200ApplicationjwtResponse{}

	body := bytes.NewReader(token.Bytes())
	res.Body = body
	res.ContentLength = int64(len(token.Bytes()))
	res.Headers.ETag = token.ID

	return res
}
