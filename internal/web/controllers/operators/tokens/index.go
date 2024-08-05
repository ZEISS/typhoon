package tokens

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// IndexOperatorTokenControllerImpl ...
type IndexOperatorTokenControllerImpl struct {
	ID uuid.UUID `json:"id" form:"id" params:"id" validate:"required:uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewIndexOperatorTokenController ...
func NewIndexOperatorTokenController(store ports.Datastore) *IndexOperatorTokenControllerImpl {
	return &IndexOperatorTokenControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Get ...
func (l *IndexOperatorTokenControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	operator := models.Operator{ID: l.ID}

	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetOperator(ctx, &operator)
	})

	if err != nil {
		return err
	}

	r := strings.NewReader(operator.Token.Token)

	l.Ctx().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jwt", operator.Name))

	return l.Ctx().SendStream(r)
}
