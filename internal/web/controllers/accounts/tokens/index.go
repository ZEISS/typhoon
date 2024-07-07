package tokens

import (
	"context"
	"fmt"
	"strings"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// GetAccountTokenControllerImpl ...
type GetAccountTokenControllerImpl struct {
	account models.Account
	store   ports.Datastore
	htmx.DefaultController
}

// NewGetAccountTokenController ...
func NewGetAccountTokenController(store ports.Datastore) *GetAccountTokenControllerImpl {
	return &GetAccountTokenControllerImpl{store: store}
}

// Prepare ...
func (l *GetAccountTokenControllerImpl) Prepare() error {
	err := l.BindParams(&l.account)
	if err != nil {
		return err
	}

	return l.store.ReadTx(l.Context(), func(ctx context.Context, rt ports.ReadTx) error {
		return rt.GetAccount(ctx, &l.account)
	})
}

// Get ...
func (l *GetAccountTokenControllerImpl) Get() error {
	r := strings.NewReader(l.account.Token.Token)

	l.Ctx().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jwt", l.account.Name))

	return l.Ctx().SendStream(r)
}
