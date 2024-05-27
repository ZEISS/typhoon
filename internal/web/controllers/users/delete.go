package users

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteUserControllerImpl{})

// DeleteUserControllerImpl ...
type DeleteUserControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" param:"id" validate:"required,uuid"`

	ports.Users
	htmx.DefaultController
}

// NewDeleteUserController ...
func NewDeleteUserController(db ports.Users) *DeleteUserControllerImpl {
	return &DeleteUserControllerImpl{
		Users:             db,
		DefaultController: htmx.DefaultController{},
	}
}

// Delete ...
func (l *DeleteUserControllerImpl) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	user := models.User{ID: l.ID}
	err = l.DeleteUser(l.Context(), &user)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/users")

	return nil
}
