package credentials

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

// IndexUserCredentialsControllerImpl ...
type IndexUserCredentialsControllerImpl struct {
	ID uuid.UUID `json:"id" form:"id" param:"id" validate:"required:uuid"`

	ports.Users
	htmx.DefaultController
}

// NewIndexUserCredentialsController ...
func NewIndexUserCredentialsController(db ports.Users) *IndexUserCredentialsControllerImpl {
	return &IndexUserCredentialsControllerImpl{
		Users:             db,
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

	err = l.GetUser(l.Context(), &user)
	if err != nil {
		return err
	}

	bb, err := user.Credentials()
	if err != nil {
		return err
	}

	r := bytes.NewReader(bb)

	l.Ctx().Set("Content-Disposition", `attachment; filename="credentials.creds"`)
	return l.Ctx().SendStream(r)
}
