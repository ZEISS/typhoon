package tokens

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// GetAccountTokenControllerParams ...
type GetAccountTokenControllerParams struct {
	ID uuid.UUID `json:"id" form:"id" params:"id" validate:"required,uuid"`
}

// GetAccountTokenControllerImpl ...
type GetAccountTokenControllerImpl struct {
	Params GetAccountTokenControllerParams

	store ports.Datastore
	htmx.DefaultController
}

// NewGetAccountTokenController ...
func NewGetAccountTokenController(store ports.Datastore) *GetAccountTokenControllerImpl {
	return &GetAccountTokenControllerImpl{
		Params:            GetAccountTokenControllerParams{},
		store:             store,
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

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, rt ports.ReadTx) error {
		return rt.GetAccount(ctx, &account)
	})
	if err != nil {
		return err
	}

	r := strings.NewReader(account.Token.Token)

	l.Ctx().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jwt", account.Name))

	return l.Ctx().SendStream(r)
}
