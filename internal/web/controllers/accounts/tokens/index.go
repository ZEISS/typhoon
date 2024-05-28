package tokens

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// GetAccountTokenControllerParams ...
type GetAccountTokenControllerParams struct {
	ID uuid.UUID `json:"id" form:"id" param:"id" validate:"required,uuid"`
}

// GetAccountTokenControllerImpl ...
type GetAccountTokenControllerImpl struct {
	Params GetAccountTokenControllerParams
	ports.Accounts
	htmx.DefaultController
}

// NewGetAccountTokenController ...
func NewGetAccountTokenController(db ports.Accounts) *GetAccountTokenControllerImpl {
	return &GetAccountTokenControllerImpl{
		Params:            GetAccountTokenControllerParams{},
		Accounts:          db,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *GetAccountTokenControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	account := models.Account{ID: l.Params.ID}

	err = l.GetAccount(l.Context(), &account)
	if err != nil {
		return err
	}

	r := strings.NewReader(account.Token.Token)

	l.Ctx().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jwt", account.Name))

	return l.Ctx().SendStream(r)
}
