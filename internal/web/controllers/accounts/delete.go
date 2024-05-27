package accounts

import (
	"github.com/google/uuid"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/web/ports"
)

var _ = htmx.Controller(&DeleteAccountControllerImpl{})

// DeleteAccountControllerImpl ...
type DeleteAccountControllerImpl struct {
	ID uuid.UUID `json:"name" form:"name" param:"id" validate:"required,uuid"`

	ports.Accounts
	htmx.DefaultController
}

// NewDeleteAccountController ...
func NewDeleteAccountController(db ports.Accounts) *DeleteAccountControllerImpl {
	return &DeleteAccountControllerImpl{
		Accounts:          db,
		DefaultController: htmx.DefaultController{},
	}
}

// Delete ...
func (l *DeleteAccountControllerImpl) Delete() error {
	err := l.BindParams(l)
	if err != nil {
		return err
	}

	account := models.Account{ID: l.ID}
	err = l.DeleteAccount(l.Context(), &account)
	if err != nil {
		return err
	}

	htmx.Redirect(l.Ctx(), "/accounts")

	return nil
}
