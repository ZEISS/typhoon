package credentials

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// IndexUserCredentialsControllerImpl ...
type IndexUserCredentialsControllerImpl struct {
	ID uuid.UUID `json:"id" form:"id" params:"id" validate:"required:uuid"`

	store ports.Datastore
	htmx.DefaultController
}

// NewIndexUserCredentialsController ...
func NewIndexUserCredentialsController(store ports.Datastore) *IndexUserCredentialsControllerImpl {
	return &IndexUserCredentialsControllerImpl{
		store:             store,
		DefaultController: htmx.DefaultController{},
	}
}

// Error ...
func (l *IndexUserCredentialsControllerImpl) Error(err error) error {
	fmt.Println(err)
	return nil
}

// Get ...
func (l *IndexUserCredentialsControllerImpl) Get() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	user := models.User{ID: l.ID}
	err = l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		return tx.GetUser(ctx, &user)
	})
	if err != nil {
		return err
	}

	bb, err := user.Credentials()
	if err != nil {
		return err
	}

	fmt.Println(bb)

	r := bytes.NewReader(bb)

	l.Ctx().Set("Content-Disposition", `attachment; filename="credentials.creds"`)

	return l.Ctx().SendStream(r)
}
